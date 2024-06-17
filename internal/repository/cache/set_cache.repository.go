package cache

import (
	"context"

	"github.com/api-monolith-template/internal/config"
	"github.com/goccy/go-json"
)

func (r *Repository) SetCache(ctx context.Context, key string, value any, opts ...CacheOpt) error {
	expDuration := config.Env.Redis.DefaultCacheDuration

	cacheOpts := new(CacheOption)

	for _, opt := range opts {
		opt(cacheOpts)
	}

	if cacheOpts.ExpiredDuration.Seconds() != 0 {
		expDuration = cacheOpts.ExpiredDuration
	}

	cacheData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.rdb.Set(ctx, key, cacheData, expDuration).Err()
	if err != nil {
		return err
	}

	return nil
}
