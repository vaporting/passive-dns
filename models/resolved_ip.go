package models

import "time"

// ResolvedIP is the table:Resolved_ips template
type ResolvedIP struct {
	BaseModel
	DomainID uint
	ResolvedIPID uint
	SourceID uint
	FirstSeen time.Time
	LastSeen time.Time
}