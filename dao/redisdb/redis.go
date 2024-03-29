package redisdb

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"raddit/config"
	"time"
)

var (
	rdb *redis.Client
	ctx context.Context
)

func Init(cfg *config.RedisConfig) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx = context.Background()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(timeoutCtx).Result()
	return err
}
