package redis

import (
	"encoding/binary"
	"fmt"

	"testing"

	"time"

	"github.com/stretchr/testify/assert"

	"passive-dns/models"
)

func TestNewResolvedDomainByModels(t *testing.T) {
	loc, _ := time.LoadLocation("Etc/GMT+0")
	ele := models.ResolvedDomainDD{
		ID:        1,
		FirstSeen: time.Date(2019, time.November, 6, 12, 00, 00, 00, loc),
		LastSeen:  time.Date(2019, time.November, 6, 12, 00, 00, 00, loc),
		Dname:     "www.google.com",
		Cname:     "tw.google.com"}
	expK := RDomainKeyPrefix + fmt.Sprint(ele.ID)
	idBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(idBytes, uint32(ele.ID))
	fsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(fsBytes, uint64(ele.FirstSeen.Unix()))
	lsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lsBytes, uint64(ele.LastSeen.Unix()))
	expV := []interface{}{
		RDID,
		idBytes,
		RDDOMAIN,
		ele.Dname,
		RDCNAME,
		ele.Cname,
		RDFirstSeen,
		fsBytes,
		RDLastSeen,
		lsBytes}

	rEle := NewResolvedDomainByModel(ele)

	values := rEle.Values()
	assert.Equal(t, expK, rEle.Key)
	assert.EqualValues(t, expV, values)
}

func TestNewResolvedDomainByKeyValues(t *testing.T) {
	loc, _ := time.LoadLocation("Etc/GMT+0")
	id := uint(1)
	key := RIPKeyPrefix + fmt.Sprint(id)
	domain := "www.google.com"
	cname := "tw.google.com"
	firstSeen := time.Date(2019, time.November, 6, 12, 00, 00, 00, loc)
	lastSeen := time.Date(2019, time.November, 6, 12, 00, 00, 00, loc)
	values := make([]string, 7)
	idBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(idBytes, uint32(id))
	values[0] = string(idBytes)
	values[1] = domain
	values[2] = cname
	seenBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(seenBytes, uint64(firstSeen.Unix()))
	values[3] = string(seenBytes)
	seenBytes = make([]byte, 8)
	binary.LittleEndian.PutUint64(seenBytes, uint64(lastSeen.Unix()))
	values[4] = string(seenBytes)

	rEle := NewResolvedDomainByKeyValues(key, values)

	assert.Equal(t, key, rEle.Key)
	assert.Equal(t, id, rEle.ID)
	assert.Equal(t, domain, rEle.Domain)
	assert.Equal(t, cname, rEle.Cname)
	assert.Equal(t, true, firstSeen.Equal(rEle.FirstSeen))
	assert.Equal(t, true, lastSeen.Equal(rEle.LastSeen))
}
