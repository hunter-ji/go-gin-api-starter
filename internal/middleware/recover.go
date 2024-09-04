// @Title globalRecover.go
// @Description
// @Author Hunter 2024/9/4 10:25

package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-gin-api-starter/pkg/util/response"
)

// globalRecover
// @Description: global globalRecover
// @return gin.HandlerFunc
func globalRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				stack := debug.Stack()
				logrus.Errorf("Panic recovered: %v\nStack Trace:\n%s", r, string(stack))

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    response.StatusInternalServerError,
					"message": "Internal Server Error",
				})
			}
		}()

		c.Next()
	}
}
