// @Title router.go
// @Description
// @Author Hunter 2024/9/3 17:10

package api

import (
	"github.com/gin-gonic/gin"
	api "go-gin-api-starter/internal/api/handler"
)

func LoadRouter(e *gin.Engine) {
	e.GET("/ping", api.Ping)
}
