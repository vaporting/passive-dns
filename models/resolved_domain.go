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
