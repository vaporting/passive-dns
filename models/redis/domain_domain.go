package redis

// DomainD is the data structure storing the index between domain and resolved_domain in redis
type DomainD struct {
	Key   string // key: d_d:[domain]
	RdKey string // value
}

// NewDomainD is the way to create DomainD
func NewDomainD(domain string, rdKey string) DomainD {
	return DomainD{
		Key:   "d_d:" + domain,
		RdKey: rdKey}
}
