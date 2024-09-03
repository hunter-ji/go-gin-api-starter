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
	rdbServer, rdbPort, rdbPassword string
	rdbDB                           int
	RDB                             *redis.Client
)

func init() {
	rdbServer = config.RedisConfig.Server
	rdbPort = config.RedisConfig.Port
	rdbPassword = config.RedisConfig.Password
	rdbDB = config.RedisConfig.DB

	RDB = RedisConnect()

	fmt.Printf("Redis: %s:%s\n\n", rdbServer, rdbPort)
}

// RedisConnect
// Description: Connect to Redis database and return the client
// @return rdb Connection client
func RedisConnect() (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rdbServer, rdbPort),
		Password: rdbPassword,
		DB:       rdbDB,
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
		return nil
	}

	return
}
