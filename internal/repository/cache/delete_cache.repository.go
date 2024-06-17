package cache

import (
	"context"
	"fmt"
	"sync"

	"github.com/api-monolith-template/internal/config"
	"github.com/sirupsen/logrus"
)

func (r *Repository) DeleteCache(ctx context.Context, cacheKeyPatterns ...string) error {
	if config.Env.Redis.IsCacheDisable {
		return nil
	}

	var wg sync.WaitGroup
	wg.Add(len(cacheKeyPatterns))

	for _, cacheKeyPattern := range cacheKeyPatterns {
		go func(pattern string) {
			defer wg.Done()
			err := r.deleteKeysMatchingPattern(ctx, pattern)
			if err != nil {
				logrus.Errorf("error deleting cache keys for pattern %s: %v", pattern, err)
			}
		}(cacheKeyPattern)
	}

	wg.Wait()
	return nil
}

func (r *Repository) deleteKeysMatchingPattern(ctx context.Context, cacheKeyPattern string) error {
	var cursor uint64 = 0
	var keys []string
	var err error

	for {
		keys, cursor, err = r.rdb.Scan(ctx, cursor, cacheKeyPattern, 100).Result()
		if err != nil {
			return err
		}

		var wg sync.WaitGroup
		keyChan := make(chan string, len(keys))
		for _, key := range keys {
			wg.Add(1)
			keyChan <- key
			go r.deleteKey(ctx, &wg, keyChan)
		}

		close(keyChan)
		wg.Wait()

		if cursor == 0 {
			break
		}
	}

	return nil
}

func (r *Repository) deleteKey(ctx context.Context, wg *sync.WaitGroup, keyChan <-chan string) {
	defer wg.Done()
	pipe := r.rdb.Pipeline()
	for key := range keyChan {
		pipe.Del(ctx, key)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		logrus.Error(fmt.Errorf("error deleting keys: %v", err.Error()))
	}
}
