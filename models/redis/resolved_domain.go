package redis

import (
	"encoding/binary"

	"fmt"

	"time"

	"passive-dns/models"
)

// RDVar describes the ResolvedIP value
type RDVar int

// MarshalBinary uses to decode as binary array
func (rdv RDVar) MarshalBinary() ([]byte, error) {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(rdv))

	return []byte{bytes[0]}, nil
}

const (
	RDKey RDVar = iota // RDKEY == 0
	RDID
	RDDOMAIN
	RDCNAME
	RDFirstSeen
	RDLastSeen
)

// ResolvedDomain is the data structure of resolved_ip in redis
type ResolvedDomain struct {
	Key       string
	ID        uint
	Domain    string
	Cname     string
	FirstSeen time.Time
	LastSeen  time.Time
}

// Values builds value:array of interface
func (m *ResolvedDomain) Values() []interface{} {
	idBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(idBytes, uint32(m.ID))
	fsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(fsBytes, uint64(m.FirstSeen.Unix()))
	lsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lsBytes, uint64(m.LastSeen.Unix()))
	return []interface{}{
		RDID,
		idBytes,
		RDDOMAIN,
		m.Domain,
		RDCNAME,
		m.Cname,
		RDFirstSeen,
		fsBytes,
		RDLastSeen,
		lsBytes}
}

// NewResolvedDomainByModel creates ResolvedDomain by models.ResolvedIPDD
func NewResolvedDomainByModel(s models.ResolvedDomainDD) ResolvedDomain {
	return ResolvedDomain{
		Key:       "r_d:" + fmt.Sprint(s.ID),
		ID:        s.ID,
		Domain:    s.Dname,
		Cname:     s.Cname,
		FirstSeen: s.FirstSeen,
		LastSeen:  s.LastSeen}
}
