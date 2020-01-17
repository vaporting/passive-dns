package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDomainD(t *testing.T) {
	name := "www.google.com"
	ripKey := "d_d:1"

	ele := NewDomainD(name, ripKey)

	assert.Equal(t, "d_d:"+name, ele.Key)
	assert.Equal(t, ripKey, ele.RdKey)
}
