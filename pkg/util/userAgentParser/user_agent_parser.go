// @Title user_agent_parser.go
// @Description
// @Author Hunter 2024/9/3 17:32

package userAgentParser

import (
	"errors"

	"github.com/mileusna/useragent"
)

type Device struct {
	Browser        string `db:"browser" json:"browser"`
	BrowserVersion string `db:"browserVersion" json:"browserVersion"`
	OS             string `db:"os" json:"os"`
	OSVersion      string `db:"osVersion" json:"osVersion"`
}

// GetDeviceInfo
// @Description: parse user-agent and get request's device info
// @param userAgent user-agent
// @return d device info
// @return isBot the request was initiated by a script this time
// @return err
func GetDeviceInfo(userAgent string) (d Device, isBot bool, err error) {
	ua := useragent.Parse(userAgent)

	d.Browser = ua.Name
	d.BrowserVersion = ua.Version
	d.OS = ua.OS
	d.OSVersion = ua.OSVersion

	isBot = ua.Bot
	if isBot {
		return
	}

	if d.Browser == "" || d.BrowserVersion == "" || d.OS == "" || d.OSVersion == "" {
		err = errors.New("failed to parse user-agent")
		return
	}

	return
}
