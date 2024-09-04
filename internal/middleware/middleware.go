// @Title middleware.go
// @Description
// @Author Hunter 2024/9/4 10:34

package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-gin-api-starter/config"
)

func SetupMiddleware(r *gin.Engine) {
	r.Use(cors.Default())

	r.Use(limitIP(20))

	r.Use(authorize())

	// global recover
	r.Use(globalRecover())

	// print http request and response in development and test environments
	if config.NodeEnv != config.Production {
		r.Use(printHTTPLog())
	}
}
