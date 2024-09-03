// @Title producer.go
// @Description
// @Author Hunter 2024/9/3 18:05

package rabbitMQ

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var connectionPool map[string]*connection

func createConnection(exchange, exchangeType string) *connection {
	c := &connection{
		exchange:     exchange,
		exchangeType: exchangeType,
		err:          make(chan error),
	}
	go c.Listen()
	return c
}

func (c *connection) Connect() error {
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		logrus.Errorf("failed to create connection, reason: %s", err)
		return err
	}

	c.conn = conn
	c.channels = map[string]*amqp.Channel{}

	go func() {
		<-c.conn.NotifyClose(make(chan *amqp.Error)) // Listen to NotifyClose
		c.err <- errors.New("connection closed")
	}()

	c.defaultChannel, err = c.conn.Channel()
	if err != nil {
		logrus.Errorf("failed to create channel: %s", err)
		return err
	}
	if err := c.defaultChannel.ExchangeDeclare(
		c.exchange,     // name
		c.exchangeType, // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // noWait
		nil,            // arguments
	); err != nil {
		logrus.Errorf("failed to create exchange declare: %s", err)
		return err
	}

	return nil
}

func publishAction(ch *amqp.Channel, exchangeName, routingKey string, message []byte) error {
	err := ch.Publish(
		exchangeName, // exchangeName
		routingKey,   // routing routingKey
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		logrus.Errorf("failed to publish message, body: %v, reason: %s", string(message), err)
		return err
	}
	logrus.Infof("success to publish message to %s, body: %v", exchangeName, string(message))
	return nil
}

func newClient(hash, exchangeName, exchangeType string) *connection {
	c, ok := connectionPool[hash]
	if !ok {
		if len(connectionPool) == 0 {
			connectionPool = make(map[string]*connection)
		}
		connectionPool[hash] = createConnection(exchangeName, exchangeType)
		c = connectionPool[hash]
		if err := c.Connect(); err != nil {
			logrus.Fatal(err)
		}
	}

	return c
}

func Publish(exchangeName, exchangeType, routingKey string, message []byte) error {
	hash := fmt.Sprintf("%s-%s-%s", exchangeName, exchangeType, routingKey)

	c := newClient(hash, exchangeName, exchangeType)

	if ch, ok := c.channels[hash]; ok {
		return publishAction(ch, exchangeName, routingKey, message)
	} else {
		channel, err := c.conn.Channel()
		if err != nil {
			return fmt.Errorf("channel: %s", err)
		}

		if !strings.HasPrefix(exchangeName, "amq.") {
			err = channel.ExchangeDeclare(
				exchangeName,
				exchangeType,
				true,
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				return fmt.Errorf("exchangeName Declare: %s", err)
			}
		}

		if reliable {
			// Reliable publisher confirms require confirm.select support from the connection.
			if err := channel.Confirm(false); err != nil {
				return fmt.Errorf("channel could not be put into confirm mode: %s", err)
			}

			confirms := channel.NotifyPublish(make(chan amqp.Confirmation, 1))

			defer confirmOne(confirms)
		}

		c.channels[hash] = channel
		return publishAction(channel, exchangeName, routingKey, message)
	}
}

// Reconnect reconnects the connection
func (c *connection) Reconnect() error {
	if err := c.Connect(); err != nil {
		return err
	}
	return nil
}

func (c *connection) Listen() {
	logrus.Infof("start listen channel: %s", c.exchange)
	for {
		if err := <-c.err; err != nil {
			err := c.Reconnect()
			if err != nil {
				logrus.Errorf("failed to reconnect mq, reason: %s", err)
				time.Sleep(5 * time.Second)
			}
		}
	}
}

func confirmOne(confirms <-chan amqp.Confirmation) {
	logrus.Infof("waiting for confirmation of one publishing")

	if confirmed := <-confirms; confirmed.Ack {
		logrus.Infof("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
	} else {
		logrus.Infof("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
	}
}
