// @Title router.go
// @Description
// @Author Hunter 2024/9/3 17:10

package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-api-starter/internal/handler"
)

func LoadRouter(e *gin.Engine) {
	e.GET("/ping", handler.Ping)
}
