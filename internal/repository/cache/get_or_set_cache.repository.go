package cache

import (
	"context"
	"errors"
	"reflect"

	"github.com/api-monolith-template/internal/config"
	"github.com/goccy/go-json"
	"gorm.io/gorm"

	"github.com/redis/go-redis/v9"
)

func (r *Repository) GetOrSetCache(ctx context.Context, key string, out any, fallbackFn func(ctx context.Context) (any, error), opts ...CacheOpt) error {
	if config.Env.Redis.IsCacheDisable {
		return r.callFallbackAndSetCache(ctx, key, fallbackFn, out, opts...)
	}

	valOut := reflect.ValueOf(out)

	if valOut.Kind() != reflect.Ptr {
		return errors.ErrUnsupported
	}

	err := r.GetCache(ctx, key, out, opts...)
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}

	if err == nil {
		return nil
	}

	return r.callFallbackAndSetCache(ctx, key, fallbackFn, out, opts...)
}

func (r *Repository) callFallbackAndSetCache(ctx context.Context, key string, fallbackFn func(ctx context.Context) (any, error), out any, opts ...CacheOpt) error {
	value, err := fallbackFn(ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	err = r.SetCache(ctx, key, value, opts...)
	if err != nil {
		return nil
	}

	newCacheData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return json.Unmarshal(newCacheData, out)
}
