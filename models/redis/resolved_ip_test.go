package redis

import (
	"fmt"

	"testing"

	"time"

	"github.com/stretchr/testify/assert"

	"passive-dns/models"
)

func TestResolvedIPStrings(t *testing.T) {
	loc, _ := time.LoadLocation("Etc/GMT+0")
	ele := models.ResolvedIPDIP{
		ID:        1,
		FirstSeen: time.Date(2019, time.November, 6, 12, 00, 00, 00, loc),
		LastSeen:  time.Date(2019, time.November, 6, 12, 00, 00, 00, loc),
		Dname:     "www.google.com",
		Ip:        "8.8.8.8",
		Type:      "A"}
	expK := "resolved_ip:" + fmt.Sprint(ele.ID)
	expV := []string{
		"id",
		fmt.Sprint(ele.ID),
		"domain",
		ele.Dname,
		"ip",
		ele.Ip,
		"type",
		ele.Type,
		"first_seen",
		ele.FirstSeen.String(),
		"last_seen",
		ele.LastSeen.String()}
	rEle := NewResolvedIPByModel(ele)

	values := rEle.VStrings()

	assert.Equal(t, expK, rEle.Key)
	assert.EqualValues(t, expV, values)
}
