package cache

import (
	"context"
	"errors"
	"reflect"

	"github.com/goccy/go-json"
	"gorm.io/gorm"

	"github.com/redis/go-redis/v9"
)

func (r *Repository) GetOrSetCache(ctx context.Context, key string, out any, fallbackFn func(ctx context.Context) (any, error), opts ...CacheOpt) error {
	valOut := reflect.ValueOf(out)

	if valOut.Kind() != reflect.Ptr {
		return errors.ErrUnsupported
	}

	err := r.GetCache(ctx, key, out, opts...)
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}

	// cache found
	if err == nil {
		return nil
	}

	// call fallback to get value from other source
	value, err := fallbackFn(ctx)
	// ignore record not found error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	// set new cache
	err = r.SetCache(ctx, key, value, opts...)
	if err != nil {
		return nil
	}

	// assign to out param
	newCacheData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = json.Unmarshal(newCacheData, &out)
	if err != nil {
		return err
	}

	return nil
}
