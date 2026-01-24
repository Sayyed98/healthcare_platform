package utils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis(cfg *Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       0,
	})

	return rdb
}

func CreateSession(rdb *redis.Client, sessionID string, userID int64) error {
	return rdb.Set(ctx, sessionID, userID, 24*time.Hour).Err()
}

func GetSession(rdb *redis.Client, sessionID string) (string, error) {
	return rdb.Get(ctx, sessionID).Result()
}

func DeleteSession(rdb *redis.Client, sessionID string) error {
	return rdb.Del(ctx, sessionID).Err()
}
