package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDomainIP(t *testing.T) {
	domain := "www.google.com"
	ripKey := "resolved_ip:1"

	ele := NewDomainIP(domain, ripKey)

	assert.Equal(t, "d_ip:"+domain, ele.Key)
	assert.Equal(t, ripKey, ele.RIPKey)
}
