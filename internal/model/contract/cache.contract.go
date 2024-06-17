package contract

import (
	"context"

	"github.com/api-monolith-template/internal/repository/cache"
)

type CacheRepository interface {
	SetCache(ctx context.Context, key string, value any, opts ...cache.CacheOpt) error
	GetCache(ctx context.Context, key string, out any, opts ...cache.CacheOpt) error
	GetOrSetCache(ctx context.Context, key string, out any, fallbackFn func(ctx context.Context) (any, error), opts ...cache.CacheOpt) error
	DeleteCache(ctx context.Context, cacheKeyPatterns ...string) error
}
