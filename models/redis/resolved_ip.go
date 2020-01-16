package redis

import (
	"encoding/binary"

	"time"

	"fmt"

	"passive-dns/models"
)

// RIPVar describes the ResolvedIP value
type RIPVar int

// MarshalBinary uses to decode as binary array
func (ripv RIPVar) MarshalBinary() ([]byte, error) {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(ripv))

	return []byte{bytes[0]}, nil
}

const (
	RIPKey RIPVar = iota // RIPKEY == 0
	RIPID
	RIPDOMAIN
	RIPIP
	RIPType
	RIPFirstSeen
	RIPLastSeen
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

// Values builds value:array of interface
func (m *ResolvedIP) Values() []interface{} {
	//return []interface{}{"first_seen", []byte{0x5d, 0x7d, 0x7f, 0x00}}
	idBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(idBytes, uint32(m.ID))
	fsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(fsBytes, uint64(m.FirstSeen.Unix()))
	lsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lsBytes, uint64(m.LastSeen.Unix()))
	return []interface{}{
		RIPID,
		idBytes,
		RIPDOMAIN,
		m.Domain,
		RIPIP,
		m.IP,
		RIPType,
		m.Type,
		RIPFirstSeen,
		fsBytes,
		RIPLastSeen,
		lsBytes}
}

// NewResolvedIPByModel creates ResolvedIp by models.ResolvedIPDIP
func NewResolvedIPByModel(s models.ResolvedIPDIP) ResolvedIP {
	return ResolvedIP{
		Key:       "r_ip:" + fmt.Sprint(s.ID),
		ID:        s.ID,
		Domain:    s.Dname,
		IP:        s.Ip,
		Type:      s.Type,
		FirstSeen: s.FirstSeen,
		LastSeen:  s.LastSeen}
}
