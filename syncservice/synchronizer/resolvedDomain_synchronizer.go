package synchronizer

import (
	"fmt"

	"sync"

	"time"

	"github.com/go-pg/pg"

	"github.com/go-redis/redis/v7"

	"passive-dns/cache"
	
	mRedis "passive-dns/models/redis"

	"passive-dns/types"

	"passive-dns/db"

	"passive-dns/models"
)

// ResolvedDomainSynchronizer synchronizes resolved_domain between sql db and redis
type ResolvedDomainSynchronizer struct {
	db         *pg.DB
	cacher     *redis.Client
	joinCmd    string
	writerNo   int
	tube       chan []models.ResolvedDomainDD
	wg         sync.WaitGroup
	mux        sync.Mutex
	cacheCount int

	// Interface
	ISynchronizer
}

// Sync resolved_ip between sql db and redis
func (syncer *ResolvedDomainSynchronizer) Sync() {
	defer close(syncer.tube)
	extNum := 1
	syncer.wg.Add(extNum)
	for i := 0; i < extNum; i++ {
		go syncer.extract()
	}
	syncer.wg.Add(syncer.writerNo)
	for i := 0; i < syncer.writerNo; i++ {
		go syncer.insert()
	}
	syncer.wg.Wait()
	fmt.Println("Total number of element in cache: ", syncer.cacheCount)
}

func (syncer *ResolvedDomainSynchronizer) extract() {
	defer syncer.wg.Done()
	var index uint = 0
	num := 100000
	count := 0
	var err error = nil
	for err == nil {
		eles := []models.ResolvedDomainDD{}
		err = syncer.db.Model(&eles).
			ColumnExpr("resolved_domain.id, resolved_domain.first_seen, resolved_domain.last_seen").
			ColumnExpr("d.name AS dname, rd.name AS cname").
			Join(syncer.joinCmd, index).
			Order("resolved_domain.id ASC").
			Limit(num).
			Select()
		if len(eles) == 0 {
			break
		}
		index = eles[len(eles)-1].ID
		count += len(eles)
		next := false
		for {
			select {
			case syncer.tube <- eles:
				next = true
			default:
				time.Sleep(time.Second)
			}
			if next {
				break
			}
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("number of element from DB: ", count)
}

func (syncer *ResolvedDomainSynchronizer) insert() {
	defer syncer.wg.Done()
	baseT := time.Now()
	timeout := false
	for !timeout {
		select {
		case eles := <-syncer.tube:
			for i := 0; i < len(eles); i++ {
				rEle := mRedis.NewResolvedDomainByModel(eles[i])
				ddEle := mRedis.NewDomainD(eles[i].Dname, rEle.Key)
				var err error
				var cmd []redis.Cmder
				for n := 0; n < 1; n++ {
					pipe := syncer.cacher.TxPipeline()
					pipe.HMSet(rEle.Key, rEle.Values()...)
					pipe.SAdd(ddEle.Key, ddEle.RdKey)
					cmd, err = pipe.Exec()
					if err != nil {
						time.Sleep(time.Millisecond * 100)
					} else {
						break
					}
				}
				if err != nil {
					fmt.Println(cmd, err)
				}
			}
			syncer.mux.Lock()
			syncer.cacheCount += len(eles)
			fmt.Println("Current number in cache: ", syncer.cacheCount)
			syncer.mux.Unlock()
			baseT = time.Now()
		default:
			time.Sleep(time.Second)
			curT := time.Now()
			if curT.Sub(baseT) > time.Second*3 {
				timeout = true
			}
		}
	}
}

// NewResolvedDomainSynchronizer creates ResolvedIPSynchronizer
func NewResolvedDomainSynchronizer(config *types.Config) (ResolvedDomainSynchronizer, error) {
	joinCmd := "INNER JOIN domains AS d ON d.id = resolved_domain.domain_id" +
		" INNER JOIN domains AS rd ON rd.id = resolved_domain.resolved_domain_id" +
		" AND resolved_domain.id > ?"
	db, err := db.GetDB()
	if err != nil {
		fmt.Println(err)
		return ResolvedDomainSynchronizer{}, err
	}
	cacher, err := cache.GetCacher()
	if err != nil {
		fmt.Println(err)
		return ResolvedDomainSynchronizer{}, err
	}
	writerNo := int(cacher.PoolStats().IdleConns)
	fmt.Println("number of writer: ", writerNo)
	return ResolvedDomainSynchronizer{
		db:         db,
		cacher:     cacher,
		joinCmd:    joinCmd,
		writerNo:   writerNo,
		tube:       make(chan []models.ResolvedDomainDD, writerNo*2),
		wg:         sync.WaitGroup{},
		cacheCount: 0}, nil
}
