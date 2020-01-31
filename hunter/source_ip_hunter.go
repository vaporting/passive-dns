package hunter

import (
	"fmt"

	"passive-dns/cache"

	"passive-dns/types"

	"passive-dns/models/redis"

	"encoding/json"
)

// SourceIPHunter hunt targets with ips
type SourceIPHunter struct {
	*hunter
}

// Hunt hunts targets from sources
func (hunter *SourceIPHunter) Hunt(sources []string) ([]byte, error) {
	results := make(map[string][]types.TargetName)
	for _, source := range sources {
		result, _ := hunter.huntTargets(source)
		results[source] = result
	}

	jsonBytes, err := json.Marshal(results)
	return jsonBytes, err
}

// huntTarget hunts targets from source
func (hunter *SourceIPHunter) huntTargets(source string) ([]types.TargetName, error) {
	rIps, err := hunter.cacher.SMembers(redis.IPDomainKeyPrefix + source).Result()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	results := []types.TargetName{}
	for _, rIp := range rIps {
		vals, err := hunter.cacher.HVals(rIp).Result()
		if err != nil {
			fmt.Println(err)
			continue
		}
		realRIp := redis.NewResolvedIPByKeyValues(rIp, vals)
		temp := types.TargetName{}
		temp[realRIp.Domain] = &types.SeenGroup{
			FirstSeen: realRIp.FirstSeen.Format("2006-01-02"),
			LastSeen:  realRIp.LastSeen.Format("2006-01-02"),
		}
		results = append(results, temp)
	}
	return results, nil
}

// NewSourceIPHunter creates SourceIPHunter
func NewSourceIPHunter() (*SourceIPHunter, error) {
	cacher, err := cache.GetCacher()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	hunter := SourceIPHunter{hunter: &hunter{cacher: cacher}}
	hunter.SourceTypes = []string{types.SourceIpsType}
	return &hunter, err
}
