package cache

import (
	"errors"

	"strconv"

	"github.com/go-redis/redis/v7"

	"passive-dns/types"
)

var cacher *redis.Client

// CreateCacher creates cacher
func CreateCacher(config *types.Config) (*redis.Client, error) {
	poolSize, _ := strconv.Atoi(config.CACHE.PoolSize)
	cacher = redis.NewClient(&redis.Options{
		Addr:         config.CACHE.Host + ":" + config.CACHE.Port,
		Password:     config.CACHE.PWD,
		DB:           0, // use default DB
		PoolSize:     poolSize,
		MinIdleConns: poolSize,
	})
	_, err := cacher.Ping().Result()
	if err != nil {
		cacher = nil
	}
	return cacher, err
}

// GetCacher gets cacher
func GetCacher() (*redis.Client, error) {
	if cacher == nil {
		err := errors.New("Cacher uninitialized")
		return nil, err
	}
	return cacher, nil
}

// CloseCacher closes cacher
func CloseCacher() error {
	return cacher.Close()
}
