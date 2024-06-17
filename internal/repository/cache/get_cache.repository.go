package cache

import (
	"context"
	"errors"
	"reflect"

	"github.com/api-monolith-template/internal/config"
	"github.com/goccy/go-json"
)

func (r *Repository) GetCache(ctx context.Context, key string, out any, opts ...CacheOpt) error {
	if config.Env.Redis.IsCacheDisable {
		return nil
	}

	valOut := reflect.ValueOf(out)

	if valOut.Kind() != reflect.Ptr {
		return errors.ErrUnsupported
	}

	val, err := r.rdb.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	err = json.Unmarshal(val, out)
	if err != nil {
		return err
	}

	return nil
}
