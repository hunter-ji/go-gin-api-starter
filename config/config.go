// @Title config.go
// @Description
// @Author Hunter 2024/9/3 17:52

package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// CommonSplicePrefix Common prefixes for splicing, such as token, redis key, etc., to distinguish different projects
const CommonSplicePrefix = "go-gin-api-starter"

var DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var RedisConfig struct {
	Server   string
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

func Load() error {
	env := os.Getenv("NODE_ENV")
	if env == "" {
		env = "development" // default
	}

	// 构建 .env 文件名
	envFile := fmt.Sprintf(".env.%s", env)

	if err := godotenv.Load(envFile); err != nil {
		return err
	}

	/*
		redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
		if err != nil {
			return err
		}

		// DB
		DBConfig.Host = os.Getenv("DB_HOST")
		DBConfig.Port = os.Getenv("DB_PORT")
		DBConfig.User = os.Getenv("DB_USER")
		DBConfig.Password = os.Getenv("DB_PASSWORD")
		DBConfig.DBName = os.Getenv("DB_NAME")

		// Redis
		RedisConfig.Server = os.Getenv("REDIS_SERVER")
		RedisConfig.Port = os.Getenv("REDIS_PORT")
		RedisConfig.Password = os.Getenv("REDIS_PASSWORD")
		RedisConfig.DB = redisDB

		// MessageQueue
		MessageQueueConfig.Uri = os.Getenv("MESSAGE_QUEUE_ADDRESS")
		MessageQueueConfig.JobExchangeName = os.Getenv("MESSAGE_QUEUE_JOB_EXCHANGE_NAME")
		MessageQueueConfig.JobExchangeType = os.Getenv("MESSAGE_QUEUE_JOB_EXCHANGE_TYPE")
	*/

	return nil
}
