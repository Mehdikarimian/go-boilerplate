package db

import (
	"fmt"
	"go-boilerplate/src/config"

	_redis "github.com/go-redis/redis/v7"
)

var RedisClient *_redis.Client

//InitRedis ...
func InitRedis(selectDB ...int) {

	var redisHost = config.LoadConfig("REDIS_HOST")
	var redisPassword = config.LoadConfig("REDIS_PASSWORD")

	RedisClient = _redis.NewClient(&_redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       selectDB[0],
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

	fmt.Println("Connected to Redis!")
}

//GetRedis ...
func GetRedis() *_redis.Client {
	return RedisClient
}
