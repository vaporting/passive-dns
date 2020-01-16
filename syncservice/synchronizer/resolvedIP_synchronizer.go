package synchronizer

import (
	"fmt"

	"time"

	"sync"

	"github.com/go-redis/redis/v7"

	"github.com/go-pg/pg"

	"passive-dns/db"

	"passive-dns/models"

	"passive-dns/cache"

	"passive-dns/types"

	mRedis "passive-dns/models/redis"
)

// ResolvedIPSynchronizer synchronizes resolved_ip between sql db and redis
type ResolvedIPSynchronizer struct {
	db         *pg.DB
	cacher     *redis.Client
	joinCmd    string
	writerNo   int
	tube       chan []models.ResolvedIPDIP
	wg         sync.WaitGroup
	mux        sync.Mutex
	cacheCount int

	// Interface
	ISynchronizer
}

// Sync resolved_ip between sql db and redis
func (syncer *ResolvedIPSynchronizer) Sync() {
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

func (syncer *ResolvedIPSynchronizer) extract() {
	defer syncer.wg.Done()
	var index uint = 0
	num := 100000
	count := 0
	var err error = nil
	for err == nil {
		eles := []models.ResolvedIPDIP{}
		err = syncer.db.Model(&eles).
			ColumnExpr("resolved_ip.id, resolved_ip.first_seen, resolved_ip.last_seen").
			ColumnExpr("domains.name AS dname, encode(ips.ip, 'escape') AS ip, ips.type AS type").
			Join(syncer.joinCmd, index).
			Order("resolved_ip.id ASC").
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

func (syncer *ResolvedIPSynchronizer) insert() {
	defer syncer.wg.Done()
	baseT := time.Now()
	timeout := false
	for !timeout {
		select {
		case eles := <-syncer.tube:
			for i := 0; i < len(eles); i++ {
				rEle := mRedis.NewResolvedIPByModel(eles[i])
				ipdEle := mRedis.NewIPDomain(eles[i].Ip, rEle.Key)
				dipEle := mRedis.NewDomainIP(eles[i].Dname, rEle.Key)
				var err error
				var cmd []redis.Cmder
				for n := 0; n < 1; n++ {
					pipe := syncer.cacher.TxPipeline()
					pipe.HMSet(rEle.Key, rEle.Values()...)
					pipe.SAdd(ipdEle.Key, ipdEle.RIPKey)
					pipe.SAdd(dipEle.Key, dipEle.RIPKey)
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

// NewResolvedIPSynchronizer creates ResolvedIPSynchronizer
func NewResolvedIPSynchronizer(config *types.Config) (ResolvedIPSynchronizer, error) {
	joinCmd := "INNER JOIN domains ON domains.id = resolved_ip.domain_id" +
		" INNER JOIN ips ON ips.id = resolved_ip.resolved_ip_id" +
		" AND resolved_ip.id > ?"
	db, err := db.GetDB()
	if err != nil {
		fmt.Println(err)
		return ResolvedIPSynchronizer{}, err
	}
	cacher, err := cache.GetCacher()
	if err != nil {
		fmt.Println(err)
		return ResolvedIPSynchronizer{}, err
	}
	writerNo := int(cacher.PoolStats().IdleConns)
	fmt.Println("number of writer: ", writerNo)
	return ResolvedIPSynchronizer{
		db:         db,
		cacher:     cacher,
		joinCmd:    joinCmd,
		writerNo:   writerNo,
		tube:       make(chan []models.ResolvedIPDIP, writerNo*2),
		wg:         sync.WaitGroup{},
		cacheCount: 0}, nil
}
