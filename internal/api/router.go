// @Title router.go
// @Description
// @Author Hunter 2024/9/3 17:10

package api

import (
	"github.com/gin-gonic/gin"
	api "go-gin-api-starter/internal/api/handler"
	v1 "go-gin-api-starter/internal/api/v1"
	"go-gin-api-starter/internal/database"
	"go-gin-api-starter/internal/middleware"
	"go-gin-api-starter/internal/repository"
	"go-gin-api-starter/internal/service"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	middleware.SetupMiddleware(r)

	LoadRouter(r)

	return r
}

func LoadRouter(e *gin.Engine) {
	r := e.Group("/api")
	{
		r.GET("/ping", api.Ping)
		loadUserRouter(r)
	}
}

func loadUserRouter(e *gin.RouterGroup) {
	userRepo := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepo)
	userHandler := v1.NewUserHandler(userService)

	r := e.Group("/v1/user")
	{
		r.POST("/auth", userHandler.Login)
		r.PUT("/auth", userHandler.RefreshToken)
		r.DELETE("/auth", userHandler.Logout)
		r.GET("/:id", userHandler.GetUserInfo)
	}
}
