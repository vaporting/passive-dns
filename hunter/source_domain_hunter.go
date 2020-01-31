package hunter

import (
	"fmt"

	"passive-dns/cache"

	"passive-dns/types"

	"passive-dns/models/redis"

	"encoding/json"
)

// SourceDomainHunter hunt targets with ips
type SourceDomainHunter struct {
	*hunter
}

// Hunt hunts targets from sources
func (hunter *SourceDomainHunter) Hunt(sources []string) ([]byte, error) {
	results := make(map[string]map[string][]types.TargetName)
	for _, source := range sources {
		results[source] = make(map[string][]types.TargetName)
		result, _ := hunter.huntIPs(source)
		results[source]["ips"] = result
		result, _ = hunter.huntCnames(source)
		results[source]["cnames"] = result
	}
	fmt.Println(results)
	jsonBytes, err := json.Marshal(results)
	return jsonBytes, err
}

// huntTarget hunts targets from source
func (hunter *SourceDomainHunter) huntCnames(source string) ([]types.TargetName, error) {
	rDomains, err := hunter.cacher.SMembers(redis.DomainDKeyPrefix + source).Result()
	if err != nil {
		fmt.Println(err)
		return []types.TargetName{}, err
	}
	results := []types.TargetName{}
	for _, rDomain := range rDomains {
		vals, err := hunter.cacher.HVals(rDomain).Result()
		if err != nil {
			fmt.Println(err)
			continue
		}
		realRDomain := redis.NewResolvedDomainByKeyValues(rDomain, vals)
		temp := types.TargetName{}
		temp[realRDomain.Cname] = &types.SeenGroup{
			FirstSeen: realRDomain.FirstSeen.Format("2006-01-02"),
			LastSeen:  realRDomain.LastSeen.Format("2006-01-02"),
		}
		results = append(results, temp)
	}
	return results, nil
}

func (hunter *SourceDomainHunter) huntIPs(source string) ([]types.TargetName, error) {
	rIps, err := hunter.cacher.SMembers(redis.DomainIPKeyPrefix + source).Result()
	if err != nil {
		fmt.Println(err)
		return []types.TargetName{}, err
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
		temp[realRIp.IP] = &types.SeenGroup{
			FirstSeen: realRIp.FirstSeen.Format("2006-01-02"),
			LastSeen:  realRIp.LastSeen.Format("2006-01-02"),
		}
		results = append(results, temp)
	}
	return results, nil
}

// NewSourceDomainHunter creates SourceDomainHunter
func NewSourceDomainHunter() (*SourceDomainHunter, error) {
	cacher, err := cache.GetCacher()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	hunter := SourceDomainHunter{hunter: &hunter{cacher: cacher}}
	hunter.SourceTypes = []string{types.SourceDomainsType}
	return &hunter, err
}
