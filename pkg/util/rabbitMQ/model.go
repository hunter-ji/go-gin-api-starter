// @Title model.go
// @Description
// @Author Hunter 2024/9/3 18:06

package rabbitMQ

import "github.com/streadway/amqp"

type MessageHandler func(message []byte) error

type consumerClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	tag     string
	done    chan error
}

type connection struct {
	conn           *amqp.Connection
	channels       map[string]*amqp.Channel
	defaultChannel *amqp.Channel
	exchange       string
	exchangeType   string
	err            chan error
}
