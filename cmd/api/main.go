// @Title main.go
// @Description
// @Author Hunter 2024/9/3 16:51

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go-gin-api-starter/internal/api"
	"go-gin-api-starter/internal/database"
	"go-gin-api-starter/internal/middleware"
	"go-gin-api-starter/pkg/util/customBindValidator"
)

const defaultPort = "9000"

func main() {
	if err := run(); err != nil {
		log.Fatalf("Application failed to start: %v", err)
	}
}

func run() error {
	// Pre-detect database
	if err := database.PreDetectDatabase(); err != nil {
		return fmt.Errorf("failed to pre-detect database: %w", err)
	}

	// Initialize router
	r := configureGinEngine()

	// Start server
	port := getPort()
	log.Printf("Starting server on port %s", port)
	return r.Run("0.0.0.0:" + port)
}

func configureGinEngine() *gin.Engine {
	r := gin.Default()

	middleware.SetupMiddleware(r)
	api.LoadRouter(r)

	if err := customBindValidator.Register(); err != nil {
		log.Fatalf("Failed to register custom validator: %v", err)
	}

	return r
}

func getPort() string {
	if port, exists := os.LookupEnv("NODE_PORT"); exists {
		return port
	}
	return defaultPort
}
