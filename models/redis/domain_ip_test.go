package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDomainIP(t *testing.T) {
	domain := "www.google.com"
	ripKey := RIPKeyPrefix + "1"

	ele := NewDomainIP(domain, ripKey)

	assert.Equal(t, DomainIPKeyPrefix+domain, ele.Key)
	assert.Equal(t, ripKey, ele.RIPKey)
}
