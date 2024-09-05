// @Title user_test.go
// @Description
// @Author Hunter 2024/9/4 20:16

package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-gin-api-starter/internal/api"
	"go-gin-api-starter/internal/model"
	"go-gin-api-starter/pkg/util/response"
	"go-gin-api-starter/pkg/util/typeConversion"
)

func TestUser(t *testing.T) {
	router := api.SetUpRouter()

	// Test login
	loginRequest := model.LoginRequest{
		AccountName: "hunter",
		Password:    "5d41402abc4b2a76b9719d911017c592",
	}

	w, req := createLoginRequest(t, loginRequest)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var loginResp response.CustomResponse[model.AuthResponse]
	err := parseResponse(w.Body.String(), &loginResp)
	assert.NoError(t, err)

	assert.Equal(t, response.StatusOK, loginResp.Code)

	// Test refresh token
	refreshTokenRequest := model.RefreshTokenRequest{
		RefreshToken: loginResp.Data.RefreshToken,
	}

	w, req = createRefreshTokenRequest(t, refreshTokenRequest)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var newLoginResp response.CustomResponse[model.AuthResponse]

	err = parseResponse(w.Body.String(), &newLoginResp)
	assert.NoError(t, err)

	assert.Equal(t, response.StatusOK, newLoginResp.Code)

	// Test get user info
	userID := 1

	w, req = createGetUserInfoRequest(t, uint64(userID), newLoginResp.Data.AccessToken)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var getUserInfoResp response.CustomResponse[model.User]

	err = parseResponse(w.Body.String(), &getUserInfoResp)

	assert.NoError(t, err)

	assert.Equal(t, response.StatusOK, getUserInfoResp.Code)

	// Test logout
	w, req = createLogoutRequest(t, newLoginResp.Data.RefreshToken)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var logoutResp response.CustomResponse[string]

	err = parseResponse(w.Body.String(), &logoutResp)

	assert.NoError(t, err)

	assert.Equal(t, response.StatusOK, logoutResp.Code)

	// Test refresh token after logout
	w, req = createRefreshTokenRequest(t, refreshTokenRequest)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	err = parseResponse(w.Body.String(), &newLoginResp)

	assert.NoError(t, err)

	assert.Equal(t, response.StatusErr, newLoginResp.Code)

	// Test login with invalid password
	loginRequest = model.LoginRequest{
		AccountName: "hunter",
		Password:    "hello",
	}

	w, req = createLoginRequest(t, loginRequest)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	err = parseResponse(w.Body.String(), &loginResp)

	assert.NoError(t, err)

	assert.Equal(t, response.StatusErr, loginResp.Code)
}

func createLoginRequest(t *testing.T, loginRequest model.LoginRequest) (*httptest.ResponseRecorder, *http.Request) {
	loginRequestReader, err := typeConversion.StructToReader(loginRequest)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/user/auth", loginRequestReader)
	assert.NoError(t, err)

	return httptest.NewRecorder(), req
}

func createRefreshTokenRequest(t *testing.T, refreshTokenRequest model.RefreshTokenRequest) (*httptest.ResponseRecorder, *http.Request) {
	refreshTokenRequestReader, err := typeConversion.StructToReader(refreshTokenRequest)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/user/auth", refreshTokenRequestReader)
	assert.NoError(t, err)

	return httptest.NewRecorder(), req
}

func createLogoutRequest(t *testing.T, token string) (*httptest.ResponseRecorder, *http.Request) {
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/user/auth", nil)
	assert.NoError(t, err)

	req.Header.Set("Authorization", "Bearer "+token)

	return httptest.NewRecorder(), req
}

func createGetUserInfoRequest(t *testing.T, userID uint64, token string) (*httptest.ResponseRecorder, *http.Request) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/user/%d", userID), nil)
	assert.NoError(t, err)

	req.Header.Set("Authorization", "Bearer "+token)

	return httptest.NewRecorder(), req
}

func parseResponse(respBody string, target interface{}) error {
	return typeConversion.StringToStruct(respBody, target)
}
