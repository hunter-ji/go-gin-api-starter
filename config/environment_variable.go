// @Title environment_variable.go
// @Description
// @Author Hunter 2024/9/4 10:04

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

// NodeEnv current running environment
var NodeEnv = Development

var TokenConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
}

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

	projectRoot, err := findProjectRoot()
	if err != nil {
		panic(fmt.Errorf("failed to find project root: %w", err))
	}

	envFile := filepath.Join(projectRoot, fmt.Sprintf(".env.%s", NodeEnv))

	fmt.Printf("Loading env file: %s\n", envFile)

	if err := godotenv.Load(envFile); err != nil {
		panic(fmt.Errorf("failed to load env file: %w", err))
	}

	// Token
	TokenConfig.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
	TokenConfig.RefreshTokenSecret = os.Getenv("REFRESH_TOKEN_SECRET")

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

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find go.mod file")
		}
		dir = parent
	}
}
