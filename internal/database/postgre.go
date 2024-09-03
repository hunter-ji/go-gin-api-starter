// @Title postgre.go
// @Description
// @Author Hunter 2024/9/3 18:29

package database

import (
	"fmt"
	"sync"
	"time"

	"go-gin-api-starter/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	ormDB   *gorm.DB
	ormOnce sync.Once

	dbHost, dbPort, dbUser, dbPassword, dbName string
)

func init() {
	dbHost = config.DBConfig.Host
	dbPort = config.DBConfig.Port
	dbUser = config.DBConfig.User
	dbPassword = config.DBConfig.Password
	dbName = config.DBConfig.DBName

	fmt.Printf("Database: %s@%s:%s\n", dbUser, dbHost, dbPort)
}

// GetDB returns a singleton instance of the database connection
func GetDB() *gorm.DB {
	ormOnce.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			dbHost, dbUser, dbPassword, dbName, dbPort)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
				NoLowerCase:   true,
			},
		})

		if err != nil {
			panic(fmt.Errorf("failed to connect to PostgreSQL: %v", err))
		}
		ormDB = db

		sqlDB, err := ormDB.DB()
		if err != nil {
			panic(fmt.Errorf("failed to get underlying *sql.DB: %v", err))
		}

		// Configure connection pool
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	})

	return ormDB
}
