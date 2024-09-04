// @Title environment_variable.go
// @Description
// @Author Hunter 2024/9/4 10:04

package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// NodeEnv current running environment
var NodeEnv = Development

var DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// MessageQueueConfig rabbitMQ
var MessageQueueConfig struct {
	Uri             string
	JobExchangeName string
	JobExchangeType string
}

func init() {
	env := os.Getenv("NODE_ENV")
	if env != "" {
		NodeEnv = env
	}

	envFile := fmt.Sprintf(".env.%s", NodeEnv)

	fmt.Printf("Load env file: %s\n", envFile)

	if err := godotenv.Load(envFile); err != nil {
		panic(err)
	}

	// DB
	DBConfig.Host = os.Getenv("DB_HOST")
	DBConfig.Port = os.Getenv("DB_PORT")
	DBConfig.User = os.Getenv("DB_USER")
	DBConfig.Password = os.Getenv("DB_PASSWORD")
	DBConfig.DBName = os.Getenv("DB_DATABASE_NAME")

	// Redis
	RedisConfig.Host = os.Getenv("REDIS_HOST")
	RedisConfig.Port = os.Getenv("REDIS_PORT")
	RedisConfig.Password = os.Getenv("REDIS_PASSWORD")
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}
	RedisConfig.DB = redisDB

	/*
		// MessageQueue
		MessageQueueConfig.Uri = os.Getenv("MESSAGE_QUEUE_ADDRESS")
		MessageQueueConfig.JobExchangeName = os.Getenv("MESSAGE_QUEUE_JOB_EXCHANGE_NAME")
		MessageQueueConfig.JobExchangeType = os.Getenv("MESSAGE_QUEUE_JOB_EXCHANGE_TYPE")
	*/
}
