package redis

import (
	"time"

	"fmt"

	"passive-dns/models"
)

// ResolvedIP is the data structure of resolved_ip in redis
type ResolvedIP struct {
	Key       string
	ID        uint
	Domain    string
	IP        string
	Type      string
	FirstSeen time.Time
	LastSeen  time.Time
}

// VStrings build variables to key:string, value:array of string
func (m *ResolvedIP) VStrings() []string {
	return []string{
		"id",
		fmt.Sprint(m.ID),
		"domain",
		m.Domain,
		"ip",
		m.IP,
		"type",
		m.Type,
		"first_seen",
		m.FirstSeen.String(),
		"last_seen",
		m.LastSeen.String()}
}

// NewResolvedIPByModel creates ResolvedIp by models.ResolvedIPDIP
func NewResolvedIPByModel(s models.ResolvedIPDIP) ResolvedIP {
	return ResolvedIP{
		Key:       "resolved_ip:" + fmt.Sprint(s.ID),
		ID:        s.ID,
		Domain:    s.Dname,
		IP:        s.Ip,
		Type:      s.Type,
		FirstSeen: s.FirstSeen,
		LastSeen:  s.LastSeen}
}
