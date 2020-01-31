package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDomainD(t *testing.T) {
	name := "www.google.com"
	ripKey := DomainDKeyPrefix + "1"

	ele := NewDomainD(name, ripKey)

	assert.Equal(t, DomainDKeyPrefix+name, ele.Key)
	assert.Equal(t, ripKey, ele.RdKey)
}
