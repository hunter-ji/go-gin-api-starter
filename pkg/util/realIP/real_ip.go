// @Title real_ip.go
// @Description
// @Author Hunter 2024/9/3 17:42

package realIP

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetRealIP(c *gin.Context) string {
	// First, try to get IP using c.ClientIP()
	ip := c.ClientIP()

	// If c.ClientIP() returns an internal IP, empty string, or invalid format,
	// then check X-Real-IP header
	if ip == "" || ip == "::1" || strings.HasPrefix(ip, "127.") || !isValidIP(ip) {
		realIP := c.GetHeader("X-Real-IP")
		if realIP != "" && isValidIP(realIP) {
			ip = realIP
		}
	}

	return ip
}

// Helper function: check if the IP is valid
func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
