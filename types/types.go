package types

// DNS record types
const (
	// DNSIpv4Type record type
	DNSIpv4Type string = "A"
	// DNSIpv6Type record type
	DNSIpv6Type string = "AAAA"
	// DNSDomainType record type
	DNSDomainType string = "CNAME"
)

// ResolvedEntry is used to store the resolved entry from request
type ResolvedEntry struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	FirstSeen string `json:"first_seen"`
	LastSeen  string `json:"last_seen"`
}
