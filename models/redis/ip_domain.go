package redis

const IPDomainKeyPrefix = "ip_d:"

// IPDomain is the data structure storing the index between ip and resolved_ip in redis.
type IPDomain struct {
	Key    string // ip_d:[ip]
	RIPKey string // value
}

// NewIPDomain is the way to create IPDomain
func NewIPDomain(ip string, ripKey string) IPDomain {
	return IPDomain{
		Key:    IPDomainKeyPrefix + ip,
		RIPKey: ripKey}
}
