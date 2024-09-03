// @Title main.go
// @Description
// @Author Hunter 2024/9/3 16:51

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go-gin-api-starter/config"
	"go-gin-api-starter/internal/api"
	"go-gin-api-starter/pkg/util/customBindValidator"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// get port from environment variable, or use default port 9000
	var port string
	if value, ok := os.LookupEnv("NODE_PORT"); ok {
		port = value
	} else {
		port = "9000"
	}

	r := gin.Default()

	// load router
	api.LoadRouter(r)

	// register custom validator
	if err := customBindValidator.Register(); err != nil {
		os.Exit(1)
	}

	_ = r.Run("0.0.0.0:" + port)
}
