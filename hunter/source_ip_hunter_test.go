package hunter

import (
	"os"

	"path"

	"testing"

	"time"

	"passive-dns/cache"

	"passive-dns/types"

	"passive-dns/util"

	mRedis "passive-dns/models/redis"

	"runtime"

	"github.com/go-redis/redis/v7"

	"github.com/stretchr/testify/suite"
)

type SourceIpHunterTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
	config                        *types.Config
	cacher                        *redis.Client
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *SourceIpHunterTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")

	os.Chdir(dir)
	var err error
	suite.config, err = util.ReadConfig()
	suite.Empty(err)
	suite.cacher, err = cache.CreateCacher(suite.config)
	suite.Empty(err)
}

func (suite *SourceIpHunterTestSuite) TearDownTest() {
	suite.cacher.FlushAll()
}

func (suite *SourceIpHunterTestSuite) TestSourceIpHunterHuntTragets() {
	loc, _ := time.LoadLocation("Etc/GMT+0")
	rIP := mRedis.ResolvedIP{
		Key:       "r_ip:1",
		ID:        uint(1),
		Domain:    "www.google.com",
		IP:        "8.8.8.8",
		Type:      "A",
		FirstSeen: time.Date(2019, time.November, 6, 12, 00, 00, 00, loc),
		LastSeen:  time.Date(2019, time.November, 6, 12, 00, 00, 00, loc)}
	ipdEle := mRedis.NewIPDomain(rIP.IP, rIP.Key)
	pipe := suite.cacher.TxPipeline()
	pipe.HMSet(rIP.Key, rIP.Values()...)
	pipe.SAdd(ipdEle.Key, ipdEle.RIPKey)
	pipe.Exec()
	hunter, err := NewSourceIPHunter()
	suite.Empty(err)
	targets, err := hunter.huntTargets(rIP.IP)
	suite.Empty(err)

	v, ok := targets[0][rIP.Domain]

	suite.Equal(true, ok)
	suite.Equal(rIP.FirstSeen.Format("2006-01-02"), v.FirstSeen)
	suite.Equal(rIP.FirstSeen.Format("2006-01-02"), v.LastSeen)
}

func TestSourceIpHunterTestSuite(t *testing.T) {
	suite.Run(t, new(SourceIpHunterTestSuite))
}
