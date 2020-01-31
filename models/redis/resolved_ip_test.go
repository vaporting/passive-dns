package redis

import (
	"encoding/binary"
	"fmt"

	"testing"

	"time"

	"github.com/stretchr/testify/assert"

	"passive-dns/models"
)

func TestNewResolvedIPByModel(t *testing.T) {
	loc, _ := time.LoadLocation("Etc/GMT+0")
	ele := models.ResolvedIPDIP{
		ID:        1,
		FirstSeen: time.Date(2019, time.November, 6, 12, 00, 00, 00, loc),
		LastSeen:  time.Date(2019, time.November, 6, 12, 00, 00, 00, loc),
		Dname:     "www.google.com",
		Ip:        "8.8.8.8",
		Type:      "A"}
	expK := RIPKeyPrefix + fmt.Sprint(ele.ID)
	idBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(idBytes, uint32(ele.ID))
	fsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(fsBytes, uint64(ele.FirstSeen.Unix()))
	lsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lsBytes, uint64(ele.LastSeen.Unix()))
	expV := []interface{}{
		RIPID,
		idBytes,
		RIPDOMAIN,
		ele.Dname,
		RIPIP,
		ele.Ip,
		RIPType,
		ele.Type,
		RIPFirstSeen,
		fsBytes,
		RIPLastSeen,
		lsBytes}

	rEle := NewResolvedIPByModel(ele)

	values := rEle.Values()
	assert.Equal(t, expK, rEle.Key)
	assert.EqualValues(t, expV, values)
}

func TestNewResolvedIPByKeyValues(t *testing.T) {
	loc, _ := time.LoadLocation("Etc/GMT+0")
	id := uint(1)
	key := RIPKeyPrefix + fmt.Sprint(id)
	domain := "www.google.com"
	ip := "8.8.8.8"
	ipType := "A"
	firstSeen := time.Date(2019, time.November, 6, 12, 00, 00, 00, loc)
	lastSeen := time.Date(2019, time.November, 6, 12, 00, 00, 00, loc)
	values := make([]string, 7)
	idBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(idBytes, uint32(id))
	values[0] = string(idBytes)
	values[1] = domain
	values[2] = ip
	values[3] = ipType
	seenBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(seenBytes, uint64(firstSeen.Unix()))
	values[4] = string(seenBytes)
	seenBytes = make([]byte, 8)
	binary.LittleEndian.PutUint64(seenBytes, uint64(lastSeen.Unix()))
	values[5] = string(seenBytes)

	rEle := NewResolvedIPByKeyValues(key, values)

	assert.Equal(t, key, rEle.Key)
	assert.Equal(t, id, rEle.ID)
	assert.Equal(t, domain, rEle.Domain)
	assert.Equal(t, ip, rEle.IP)
	assert.Equal(t, ipType, rEle.Type)
	assert.Equal(t, true, firstSeen.Equal(rEle.FirstSeen))
	assert.Equal(t, true, lastSeen.Equal(rEle.LastSeen))
}
