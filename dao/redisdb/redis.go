package redisdb

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"raddit/config"
	"time"
)

var rdb *redis.Client

func Init(cfg *config.RedisConfig) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	return err
}
