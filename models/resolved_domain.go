package models

import "time"

// ResolvedDomain is the table:Resolved_domains template
type ResolvedDomain struct {
	BaseModel
	DomainID         uint
	ResolvedDomainID uint
	SourceID         uint
	FirstSeen        time.Time
	LastSeen         time.Time
}

// ResolvedDomainDD stores table:Resolved_domains with domains, cname(resolved_domain)
type ResolvedDomainDD struct {
	tableName struct{} `sql:"resolved_domains,alias:resolved_domain"`
	ID        uint
	FirstSeen time.Time
	LastSeen  time.Time

	// relations
	// Dname is relative to table:Domains:name
	Dname string
	// Cname is realtive to table:Domains:name
	Cname string
}
