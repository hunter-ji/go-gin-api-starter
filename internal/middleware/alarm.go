// @Title alarm.go
// @Description
// @Author Hunter 2024/9/4 10:26

package middleware

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"go-gin-api-starter/config"
)

// errorString
// @Description: error message
type errorString struct {
	s string
}

// errorInfo
// @Description: error detail message
type errorInfo struct {
	Time     string `json:"time"`      // time
	Alarm    string `json:"alarm"`     // alarm level
	Message  string `json:"message"`   // message
	Filename string `json:"filename"`  // error file name
	Line     int    `json:"line"`      // error line number
	FuncName string `json:"func_name"` // error function name
}

// Error
// @Description: return error message
// @receiver e errorString
// @return string error message
func (e *errorString) Error() string {
	return e.s
}

// Info
// @Description: create general error
// @param text error message
// @return error
func Info(text string) error {
	alarm("INFO", text, 2)
	return &errorString{text}
}

// Panic
// @Description: panic error level
// @param text error message
// @return error
func Panic(text string) error {
	alarm("PANIC", text, 5)
	return &errorString{text}
}

// alarm
// @Description: error method
// @param level error level
// @param str error message
// @param skip number of stack frames，0：current function，1：previous layer function，......
func alarm(level string, str string, skip int) {
	// 当前时间
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// 定义 文件名、行号、方法名
	functionName := "?"

	pc, fileName, line, ok := runtime.Caller(skip)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		functionName = filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName, ".")
	}

	var msg = errorInfo{
		Time:     currentTime,
		Alarm:    level,
		Message:  str,
		Filename: fileName,
		Line:     line,
		FuncName: functionName,
	}

	messageJson, errs := json.Marshal(msg)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}

	errorJsonInfo := string(messageJson)
	fmt.Println(errorJsonInfo)

	if level == "INFO" {
		// do something
	} else if level == "PANIC" {
		if config.NodeEnv == "production" {
			// publish message
			/*
				message := messageQueue.Message{
					Node:    "BrightWord",
					Status:  "error",
					Title:   msg.Filename,
					Content: errorJsonInfo,
					RunTime: msg.Time,
				}
				if err := message.Publish(); err != nil {
					fmt.Println("failed to publish message : ", err)
				}
			*/
		}
	}
}
