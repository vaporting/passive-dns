package redis

import (
	"encoding/binary"
	"fmt"

	"testing"

	"time"

	"github.com/stretchr/testify/assert"

	"passive-dns/models"
)

func TestResolvedIPKeyValues(t *testing.T) {
	loc, _ := time.LoadLocation("Etc/GMT+0")
	ele := models.ResolvedIPDIP{
		ID:        1,
		FirstSeen: time.Date(2019, time.November, 6, 12, 00, 00, 00, loc),
		LastSeen:  time.Date(2019, time.November, 6, 12, 00, 00, 00, loc),
		Dname:     "www.google.com",
		Ip:        "8.8.8.8",
		Type:      "A"}
	expK := "r_ip:" + fmt.Sprint(ele.ID)
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
