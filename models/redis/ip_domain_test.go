package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIPDomain(t *testing.T) {
	ip := "10.10.10.10"
	ripKey := RIPKeyPrefix + "1"

	ele := NewIPDomain(ip, ripKey)

	assert.Equal(t, IPDomainKeyPrefix+ip, ele.Key)
	assert.Equal(t, ripKey, ele.RIPKey)
}
