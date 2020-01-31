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

type SourceDomainHunterTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
	config                        *types.Config
	cacher                        *redis.Client
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *SourceDomainHunterTestSuite) SetupTest() {
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

func (suite *SourceDomainHunterTestSuite) TearDownTest() {
	suite.cacher.FlushAll()
}

func (suite *SourceDomainHunterTestSuite) TestSourceDomainHunterHuntIPs() {
	loc, _ := time.LoadLocation("Etc/GMT+0")
	rIP := mRedis.ResolvedIP{
		Key:       mRedis.RIPKeyPrefix + "1",
		ID:        uint(1),
		Domain:    "www.google.com",
		IP:        "8.8.8.8",
		Type:      "A",
		FirstSeen: time.Date(2019, time.November, 6, 12, 00, 00, 00, loc),
		LastSeen:  time.Date(2019, time.November, 6, 12, 00, 00, 00, loc)}
	dipEle := mRedis.NewDomainIP(rIP.Domain, rIP.Key)
	pipe := suite.cacher.TxPipeline()
	pipe.HMSet(rIP.Key, rIP.Values()...)
	pipe.SAdd(dipEle.Key, dipEle.RIPKey)
	pipe.Exec()
	hunter, err := NewSourceDomainHunter()
	suite.Empty(err)
	targets, err := hunter.huntIPs(rIP.Domain)
	suite.Empty(err)

	v, ok := targets[0][rIP.IP]

	suite.Equal(true, ok)
	suite.Equal(rIP.FirstSeen.Format("2006-01-02"), v.FirstSeen)
	suite.Equal(rIP.FirstSeen.Format("2006-01-02"), v.LastSeen)
}

func (suite *SourceDomainHunterTestSuite) TestSourceDomainHunterHuntCnames() {
	loc, _ := time.LoadLocation("Etc/GMT+0")
	rDomain := mRedis.ResolvedDomain{
		Key:       mRedis.RDomainKeyPrefix + "1",
		ID:        uint(1),
		Domain:    "www.google.com",
		Cname:     "tw.google.com",
		FirstSeen: time.Date(2019, time.November, 6, 12, 00, 00, 00, loc),
		LastSeen:  time.Date(2019, time.November, 6, 12, 00, 00, 00, loc)}
	ddEle := mRedis.NewDomainD(rDomain.Domain, rDomain.Key)
	pipe := suite.cacher.TxPipeline()
	pipe.HMSet(rDomain.Key, rDomain.Values()...)
	pipe.SAdd(ddEle.Key, ddEle.RdKey)
	pipe.Exec()
	hunter, err := NewSourceDomainHunter()
	suite.Empty(err)
	targets, err := hunter.huntCnames(rDomain.Domain)
	suite.Empty(err)

	v, ok := targets[0][rDomain.Cname]

	suite.Equal(true, ok)
	suite.Equal(rDomain.FirstSeen.Format("2006-01-02"), v.FirstSeen)
	suite.Equal(rDomain.FirstSeen.Format("2006-01-02"), v.LastSeen)
}

func TestSourceDomainHunterTestSuite(t *testing.T) {
	suite.Run(t, new(SourceDomainHunterTestSuite))
}
