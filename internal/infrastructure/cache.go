package infrastructure

import (
	"github.com/api-monolith-template/internal/config"
	"github.com/api-monolith-template/internal/util"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	opts, err := redis.ParseURL(config.Env.Redis.CacheDSN)
	util.ContinueOrFatal(err)

	opts.MaxIdleConns = config.Env.Redis.MaxIdleConns
	opts.MaxActiveConns = config.Env.Redis.MaxActiveConns
	opts.MaxRetries = config.Env.Redis.MaxRetry
	opts.ConnMaxLifetime = config.Env.Redis.MaxConnLifetime

	return redis.NewClient(opts)
}
