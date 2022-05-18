package cache

import (
	"context"
	"github.com/aiio/core/config"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	initOnce sync.Once
	// RDB redis client
	RDB *redis.Client
)

var ctx = context.Background()

// init cache
func init() {
	initOnce.Do(func() {
		RDB = redis.NewClient(&redis.Options{
			Addr:     config.V.Redis.Host,
			Password: config.V.Redis.Pass,
			DB:       config.V.Redis.DB,
		})
	})
}

func Set(key, value string, exp time.Duration) error {
	return RDB.Set(ctx, key, value, exp).Err()
}

func Get(key string) (string, error) {
	return RDB.Get(ctx, key).Result()
}

func Del(key string) {
	RDB.Del(ctx, key)
}
