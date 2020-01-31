package redis

import (
	"encoding/binary"

	"fmt"

	"time"

	"passive-dns/models"
)

const RDomainKeyPrefix = "r_d:"

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
		Key:       RDomainKeyPrefix + fmt.Sprint(s.ID),
		ID:        s.ID,
		Domain:    s.Dname,
		Cname:     s.Cname,
		FirstSeen: s.FirstSeen,
		LastSeen:  s.LastSeen}
}

// NewResolvedDomainByKeyValues creates ResolvedIp by key values
func NewResolvedDomainByKeyValues(key string, vals []string) ResolvedDomain {
	bID := []byte(vals[0])
	for i := len(bID); i <= 4; i++ {
		bID = append(bID, 0)
	}
	tID := binary.LittleEndian.Uint32(bID)
	bFirst := []byte(vals[3])
	for i := len(bFirst); i <= 8; i++ {
		bFirst = append(bFirst, 0)
	}
	tFirst := binary.LittleEndian.Uint64(bFirst)
	bLast := []byte(vals[4])
	for i := len(bLast); i <= 8; i++ {
		bLast = append(bLast, 0)
	}
	tLast := binary.LittleEndian.Uint64(bLast)
	return ResolvedDomain{
		Key:       key,
		ID:        uint(tID),
		Domain:    vals[1],
		Cname:     vals[2],
		FirstSeen: time.Unix(int64(tFirst), 0),
		LastSeen:  time.Unix(int64(tLast), 0)}
}
