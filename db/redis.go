package db

import (
	"context"
	"fmt"
	"os"

	_redis "github.com/go-redis/redis/v8"
)

// RedisClient
var (
	rdb *_redis.Client
	ctx = context.Background()
)

// InitRedis
func InitRedis() {

	var redisHost = os.Getenv("REDIS_HOST")
	var redisPassword = os.Getenv("REDIS_PASSWORD")

	rdb = _redis.NewClient(&_redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
		// DialTimeout:        10 * time.Second,
		// ReadTimeout:        30 * time.Second,
		// WriteTimeout:       30 * time.Second,
		// PoolSize:           10,
		// PoolTimeout:        30 * time.Second,
		// IdleTimeout:        500 * time.Millisecond,
		// IdleCheckFrequency: 500 * time.Millisecond,
		// TLSConfig: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		fmt.Println("[redis] ", err)
		return
	}

	fmt.Println("[redis] redis connected!")
}

// GetRedis
func GetRedis() *_redis.Client {
	return rdb
}
