// @Title router_white_list.go
// @Description
// @Author Hunter 2024/9/4 22:05

package config

var RouterWhiteList = map[string]string{
	"/api/ping":         "GET",
	"/api/v1/user/auth": "POST,PUT",
}
