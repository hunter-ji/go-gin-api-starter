// @Title redis.go
// @Description
// @Author Hunter 2024/9/3 18:26

package database

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"go-gin-api-starter/config"
)

var (
	RDB                           *redis.Client
	rdbHost, rdbPort, rdbPassword string
	rdbDB                         int
)

func init() {
	rdbHost = config.RedisConfig.Host
	rdbPort = config.RedisConfig.Port
	rdbPassword = config.RedisConfig.Password
	rdbDB = config.RedisConfig.DB

	fmt.Printf("Redis: %s:%s\n\n", rdbHost, rdbPort)

	RedisConnect()
}

// RedisConnect
// Description: Connect to Redis database and return the client
// @return rdb Connection client
func RedisConnect() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rdbHost, rdbPort),
		Password: rdbPassword,
		DB:       rdbDB,
	})
}

func preDetectRedis() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}
