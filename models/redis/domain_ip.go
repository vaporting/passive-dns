package redis

const DomainIPKeyPrefix = "d_ip:"

// DomainIP is the data structure storing the index between domain and resolved_ip in redis
type DomainIP struct {
	Key    string // key: d_ip:[domain]
	RIPKey string // value
}

// NewDomainIP is the way to create DomainIP
func NewDomainIP(domain string, ripKey string) DomainIP {
	return DomainIP{
		Key:    DomainIPKeyPrefix + domain,
		RIPKey: ripKey}
}
