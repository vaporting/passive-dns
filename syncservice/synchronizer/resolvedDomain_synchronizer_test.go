package synchronizer

import (
	"encoding/binary"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/stretchr/testify/suite"

	"passive-dns/models"
	mRedis "passive-dns/models/redis"
	"passive-dns/util"

	"passive-dns/db"

	"passive-dns/cache"

	"passive-dns/types"

	"os"

	"path"

	"runtime"
)

type ResolvedDomainSynchronizerTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
	config                        *types.Config
	cacher                        *redis.Client
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ResolvedDomainSynchronizerTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")

	os.Chdir(dir)
	var err error
	suite.config, err = util.ReadConfig()
	suite.Empty(err)
	_, err = db.InitDB(suite.config)
	suite.Empty(err)
	suite.cacher, err = cache.CreateCacher(suite.config)
	suite.Empty(err)
}

func (suite *ResolvedDomainSynchronizerTestSuite) TearDownTest() {
	suite.cacher.FlushAll()
}

func (suite *ResolvedDomainSynchronizerTestSuite) TestResolvedDomainSynchronizerInsert() {
	syncer, _ := NewResolvedDomainSynchronizer(suite.config)

	loc, _ := time.LoadLocation("Etc/GMT+0")
	var id uint = 1
	ele := models.ResolvedDomainDD{
		ID:        id,
		FirstSeen: time.Date(2019, time.November, 6, 12, 00, 00, 00, loc),
		LastSeen:  time.Date(2019, time.November, 6, 12, 00, 00, 00, loc),
		Dname:     "www.google.com",
		Cname:     "tw.google.com"}

	syncer.tube <- []models.ResolvedDomainDD{ele}
	syncer.wg.Add(1)
	go syncer.insert()
	syncer.wg.Wait()

	cEle := suite.cacher.HGetAll("r_d:" + strconv.Itoa(int(id))).Val()

	exp := map[string]string{}
	idBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(idBytes, uint32(ele.ID))
	fsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(fsBytes, uint64(ele.FirstSeen.Unix()))
	lsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lsBytes, uint64(ele.LastSeen.Unix()))
	exp[string(mRedis.RDID)] = string(idBytes)
	exp[string(mRedis.RDDOMAIN)] = ele.Dname
	exp[string(mRedis.RDCNAME)] = ele.Cname
	exp[string(mRedis.RDFirstSeen)] = string(fsBytes)
	exp[string(mRedis.RDLastSeen)] = string(lsBytes)

	suite.EqualValues(exp, cEle)
}

func TestResolvedDomainSynchronizerTestSuite(t *testing.T) {
	suite.Run(t, new(ResolvedDomainSynchronizerTestSuite))
}
