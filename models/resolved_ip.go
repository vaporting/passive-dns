package models

import "time"

// ResolvedIP is the table:Resolved_ips template
type ResolvedIP struct {
	BaseModel
	DomainID     uint
	ResolvedIpID uint
	SourceID     uint
	FirstSeen    time.Time
	LastSeen     time.Time

	// relations
	// Dname is relative to table:Domains:name
	Dname string
}

// ResolvedIPDIP stores table:Resolved_ips with domains, ips
type ResolvedIPDIP struct {
	tableName struct{} `sql:"resolved_ips,alias:resolved_ip"`
	ID        uint
	FirstSeen time.Time
	LastSeen  time.Time

	// relations
	// Dname is relative to table:Domains:name
	Dname string
	// IP is realtive to table:IPs:ip
	Ip string
	// Type is realtive to table:IPs:ip_type
	Type string
}
