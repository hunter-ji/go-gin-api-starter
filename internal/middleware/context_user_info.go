// @Title context_user_info.go
// @Description Get userinfo from context and format it.
// @Author Hunter 2024/9/4 10:42

package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go-gin-api-starter/pkg/auth"
)

// GetContextUserInfo
// @Description: get userinfo from context and format it
// @param c gin.Context
// @return contextUserInfo formatted userinfo
// @return err return when error, otherwise nil
func GetContextUserInfo(c *gin.Context) (contextUserInfo auth.ContextUserInfo, exists bool, err error) {
	userInfoInterface, exists := c.Get("userInfo")
	if !exists {
		return
	}

	contextUserInfo, ok := userInfoInterface.(auth.ContextUserInfo)
	if !ok {
		err = errors.New("failed to convert contextUserInfo format")
		return
	}

	return
}
