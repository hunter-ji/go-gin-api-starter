// @Title user.go
// @Description
// @Author Hunter 2024/9/4 16:10

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-gin-api-starter/internal/middleware"
	"go-gin-api-starter/internal/model"
	"go-gin-api-starter/internal/service"
	"go-gin-api-starter/pkg/util/response"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Login(c *gin.Context) {
	var loginReq model.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		response.Code(c, response.StatusBadRequest)
		return
	}

	resp, err := h.userService.Login(&loginReq)
	if err != nil {
		response.Error(c, "login failed")
		return
	}

	response.Data(c, resp)
}

func (h *UserHandler) Logout(c *gin.Context) {
	userinfo, _, err := middleware.GetContextUserInfo(c)
	if err != nil {
		logrus.Errorf("failed to get userinfo: %v\n", err)
		response.Error(c, "get userinfo failed")
		return
	}

	err = h.userService.Logout(userinfo.UserID)
	if err != nil {
		logrus.Errorf("failed to logout: %v\n", err)
		response.Error(c, "logout failed")
		return
	}

	response.Data(c, "logout success")
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	var refreshTokenReq model.RefreshTokenRequest
	if err := c.ShouldBindJSON(&refreshTokenReq); err != nil {
		response.Code(c, response.StatusBadRequest)
		return
	}

	newAccessToken, newRefreshToken, err := h.userService.RefreshToken(refreshTokenReq.RefreshToken)
	if err != nil {
		response.Error(c, "refresh token failed")
		return
	}

	response.Data(c, model.AuthResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	})
}

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	var getUserInfoReq model.GetUserInfoRequest
	if err := c.ShouldBindUri(&getUserInfoReq); err != nil {
		response.Code(c, response.StatusBadRequest)
		return
	}

	userinfo, _, err := middleware.GetContextUserInfo(c)
	if err != nil {
		response.Error(c, "get userinfo failed")
		return
	}

	if getUserInfoReq.ID != userinfo.UserID {
		response.Error(c, "only your own information can be obtained")
		return
	}

	resp, err := h.userService.GetUserByID(userinfo.UserID)
	if err != nil {
		logrus.Errorf("failed to get user info: %v", err)
		response.Error(c, "get user info failed")
		return
	}

	response.Data(c, resp)
}
