package hunter

import (
	"github.com/go-redis/redis/v7"
)

// IHunter is base interface for polymorphism uses
type IHunter interface {
	Hunt(sources []string) ([]byte, error)
}

type hunter struct {
	SourceTypes []string
	cacher      *redis.Client

	// Interface
	IHunter
}
