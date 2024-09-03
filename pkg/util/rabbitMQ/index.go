// @Title index.go
// @Description
// @Author Hunter 2024/9/3 18:04

package rabbitMQ

import "go-gin-api-starter/config"

var amqpURI string

var reliable bool

func init() {
	amqpURI = config.MessageQueueConfig.Uri
	reliable = false
}
