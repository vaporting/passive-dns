package types

import "time"

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

// ResolvedRow is the data before store to resolved table
type ResolvedRow struct {
	OriginalID uint
	PassiveID  uint
	SourceID   uint
	FirstSeen  time.Time
	LastSeen   time.Time
}
