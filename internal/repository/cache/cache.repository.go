package cache

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheOption struct {
	ExpiredDuration time.Duration
}

type CacheOpt func(*CacheOption)

func WithCustomExpiredDuration(expDuration time.Duration) CacheOpt {
	return func(o *CacheOption) {
		o.ExpiredDuration = expDuration
	}
}

type Repository struct {
	rdb *redis.Client
}

func NewRepository() *Repository {
	return new(Repository)
}

func (r *Repository) WithRedisDB(rdb *redis.Client) *Repository {
	r.rdb = rdb
	return r
}
