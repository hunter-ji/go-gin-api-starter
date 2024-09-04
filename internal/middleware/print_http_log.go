// @Title print_http_log.go
// @Description
// @Author Hunter 2024/9/4 11:07

package middleware

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanity-io/litter"
)

const (
	maxLogLength = 200 // Maximum length for request/response logging
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func truncateString(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength] + "..."
}

// LogEntry represents a structured log entry
type LogEntry struct {
	Timestamp  string      `json:"timestamp"`
	Method     string      `json:"method"`
	URL        string      `json:"url"`
	StatusCode int         `json:"status_code"`
	UserInfo   interface{} `json:"user_info,omitempty"`
	Request    string      `json:"request"`
	Response   string      `json:"response"`
}

func printHTTPLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Capture request body
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := io.ReadAll(tee)
		c.Request.Body = io.NopCloser(&buf)

		// Capture response body
		blw := &bodyLogWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		// Prepare log entry
		logEntry := LogEntry{
			Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
			Method:     c.Request.Method,
			URL:        c.Request.URL.String(),
			StatusCode: c.Writer.Status(),
			Request:    truncateString(string(body), maxLogLength),
			Response:   truncateString(blw.body.String(), maxLogLength),
		}

		// Log user info if available
		userInfo, exists, err := GetContextUserInfo(c)
		if exists {
			if err != nil {
				logEntry.UserInfo = map[string]string{"error": err.Error()}
			} else {
				logEntry.UserInfo = userInfo
			}
		}

		fmt.Println()
		litter.Dump(logEntry)
	}
}
