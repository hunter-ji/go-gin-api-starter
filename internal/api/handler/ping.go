// @Title ping.go
// @Description
// @Author Hunter 2024/9/3 18:36

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
	return
}
