// @Title response.go
// @Description
// @Author Hunter 2024/9/4 10:03

package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Custom response code
const (
	StatusOK                  = 20000 // request success
	StatusErr                 = 20001 // general error
	StatusNotModified         = 30004 // no change
	StatusBadRequest          = 40000 // bad request
	StatusUnauthorized        = 40001 // need to authenticate
	StatusInternalServerError = 50000 // internal server error
	StatusErrToken            = 50008 // token error
	StatusRepeatLogin         = 50012 // duplicate login
	StatusExpireToken         = 50014 // token expired
)

type CustomResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

// Data sends a successful response with the provided data
// @param c *gin.Context
// @param data interface{} - The data to be returned in the response
func Data[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, CustomResponse[T]{
		Code: StatusOK,
		Data: data,
	})
}

// Error sends an error response with the provided message
// @param c *gin.Context
// @param message string - The error message to be returned
func Error(c *gin.Context, message string) {
	c.JSON(http.StatusOK, CustomResponse[string]{
		Code:    StatusErr,
		Message: message,
		Data:    "",
	})
}

// Code sends a response with a specific status code and its corresponding message
// @param c *gin.Context
// @param code int - The custom status code
func Code(c *gin.Context, code int) {
	codeMap := map[int]string{
		StatusOK:           "",
		StatusNotModified:  "",
		StatusBadRequest:   "Bad Request",
		StatusUnauthorized: "Authentication Required",
		StatusErrToken:     "Token Error",
		StatusRepeatLogin:  "Duplicate Login",
		StatusExpireToken:  "Token Expired",
	}

	c.JSON(http.StatusOK, CustomResponse[string]{
		Code:    code,
		Message: codeMap[code],
		Data:    "",
	})
}
