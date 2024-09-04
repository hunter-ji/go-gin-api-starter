// @Title context_user_info.go
// @Description
// @Author Hunter 2024/9/4 10:42

package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// ContextUserInfo
// @Description: userinfo in context, update the fields according to the actual project
type ContextUserInfo struct {
	UserID uint64 `redis:"id"`
}

// GetContextUserInfo
// @Description: get userinfo from context and format it
// @param c gin.Context
// @return contextUserInfo formatted userinfo
// @return err return when error, otherwise nil
func GetContextUserInfo(c *gin.Context) (contextUserInfo ContextUserInfo, exists bool, err error) {
	userInfoInterface, exists := c.Get("userInfo")
	if !exists {
		return
	}

	contextUserInfo, ok := userInfoInterface.(ContextUserInfo)
	if !ok {
		err = errors.New("failed to convert contextUserInfo format")
		return
	}

	return
}
