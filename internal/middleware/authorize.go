// @Title authorize.go
// @Description
// @Author Hunter 2024/9/4 10:41

package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"go-gin-api-starter/config"
	"go-gin-api-starter/internal/database"
	"go-gin-api-starter/pkg/auth"
	"go-gin-api-starter/pkg/util/response"
)

const (
	ErrMsgUnauthorized = "unauthorized access"
	ErrMsgExpiredToken = "expired token"
)

// isItOnTheWhiteList
// @Description: checks if the request path and method are in the whitelist
// @param path request path
// @param method request method
// @return bool true if in whitelist, false otherwise
func isItOnTheWhiteList(path, method string) bool {
	if allowedMethod, ok := config.RouterWhiteList[path]; ok {
		if allowedMethod == "*" {
			return true
		}

		allowedMethods := strings.Split(allowedMethod, ",")
		if slices.Contains(allowedMethods, method) {
			return true
		}
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
// @return string new access token if token is expired
// @return bool true if token is expired
// @return error
func validateToken(c *gin.Context) (*ContextUserInfo, string, bool, error) {
	// Get the Authorization header
	authHeader := c.GetHeader("Authorization")

	// Check if the header is empty
	if authHeader == "" {
		return nil, "", false, fmt.Errorf("failed to get access token")
	}

	// Check if the token starts with "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, "", false, fmt.Errorf("invalid access token")
	}

	// remove prefix
	token := strings.TrimPrefix(authHeader, "Bearer ")

	claims, newAccessToken, expired, err := auth.ValidateAccessTokenAndRefresh(token, database.RDB)
	if err != nil {
		return nil, "", expired, err
	}

	var userInfo ContextUserInfo
	userInfo.UserID = claims.UserID

	return &userInfo, newAccessToken, false, nil
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
			userInfo, newAccessToken, _, err := validateToken(c)
			if err != nil {
				respondWithError(c, response.StatusUnauthorized, ErrMsgUnauthorized)
				return
			}

			// Set new access token in header only if access token is expired
			if newAccessToken != "" {
				c.Header("Authorization", newAccessToken)
			}

			// Set user info in context
			c.Set("userInfo", *userInfo)
		}

		c.Next()
	}
}
