package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIPDomain(t *testing.T) {
	ip := "10.10.10.10"
	ripKey := "resolved_ip:1"

	ele := NewIPDomain(ip, ripKey)

	assert.Equal(t, "ip_d:"+ip, ele.Key)
	assert.Equal(t, ripKey, ele.RIPKey)
}
