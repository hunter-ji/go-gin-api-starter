// @Title authorize.go
// @Description
// @Author Hunter 2024/9/4 10:41

package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go-gin-api-starter/config"
	"go-gin-api-starter/internal/database"
	"go-gin-api-starter/pkg/util/response"
)

const (
	ErrMsgUnauthorized = "unauthorized access"
)

type QueryToken struct {
	Token string `form:"token" binding:"required"`
}

// isItOnTheWhiteList
// @Description: checks if the request path and method are in the whitelist
// @param path request path
// @param method request method
// @return bool true if in whitelist, false otherwise
func isItOnTheWhiteList(path, method string) bool {
	if allowedMethod, ok := config.RouterWhiteList[path]; ok {
		return allowedMethod == "*" || allowedMethod == method
	}
	return false
}

// respondWithError
// @Description: error response
// @param c gin context
// @param statusCode custom status code
// @param message error message
func respondWithError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"message": message,
	})
}

// validateToken
// @Description: checks the token and gets user info from Redis
// @param c gin context
// @param rdb Redis client
// @return *ContextUserInfo user info if token is valid
// @return error
func validateToken(c *gin.Context, rdb *redis.Client) (*ContextUserInfo, error) {
	var queryToken QueryToken
	if err := c.ShouldBindQuery(&queryToken); err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	assembledToken := fmt.Sprintf("%s-auth:%s", config.CommonSplicePrefix, queryToken.Token)

	// Check if token exists in Redis
	exists, err := rdb.Exists(c, assembledToken).Result()
	if err != nil || exists == 0 {
		return nil, fmt.Errorf("token not found")
	}

	// Get user info from Redis
	var userInfo ContextUserInfo
	if err := rdb.HGetAll(c, assembledToken).Scan(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to get user info")
	}

	return &userInfo, nil
}

// authorize
// @Description: middleware for access control, checks whitelist and validates token
// @return gin.HandlerFunc
func authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow OPTIONS requests
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": response.StatusOK})
			return
		}

		// Check if request is in whitelist
		if !isItOnTheWhiteList(c.Request.URL.Path, c.Request.Method) {
			userInfo, err := validateToken(c, database.RDB)
			if err != nil {
				respondWithError(c, response.StatusUnauthorized, ErrMsgUnauthorized)
				return
			}

			// Set user info in context
			c.Set("userInfo", *userInfo)
			fmt.Printf("userInfo: %v\n", userInfo)
		}

		c.Next()
	}
}
