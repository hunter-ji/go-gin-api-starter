// @Title consumer.go
// @Description
// @Author Hunter 2024/9/3 18:06

package rabbitMQ

import (
	"fmt"
	"log"
	"strings"

	"github.com/streadway/amqp"
	"go-gin-api-starter/config"
	"go-gin-api-starter/pkg/util/fastTime"
	"go-gin-api-starter/pkg/util/uuid"
)

var clientID = fmt.Sprintf("%s-%s-%d", config.CommonSplicePrefix, uuid.GenerateUUID(), fastTime.UnixTimestamp())

func Consumer(exchangeName, exchangeType, queueName, routingKey string, messageHandler MessageHandler) {
	c := &consumerClient{
		conn:    nil,
		channel: nil,
		tag:     clientID,
	}

	var err error
	c.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		log.Fatalf("connect failed : %s", err)
	}
	defer c.conn.Close()

	c.channel, err = c.conn.Channel()
	if err != nil {
		log.Fatalf("ch error : %s", err)
	}
	defer c.channel.Close()

	if !strings.HasPrefix(exchangeName, "amq.") {
		err = c.channel.ExchangeDeclare(
			exchangeName,
			exchangeType,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatalf("exchange declare error : %s", err)
		}
	}

	queue, err := c.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Queue Declare: %s", err)
	}

	if err = c.channel.QueueBind(
		queue.Name,
		routingKey,
		exchangeName,
		false,
		nil,
	); err != nil {
		log.Fatalf("Queue Bind: %s", err)
	}

	deliveries, err := c.channel.Consume(
		queue.Name,
		c.tag,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Queue Consume: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range deliveries {
			log.Printf("Received a message: %s\n", d.Body)

			err := messageHandler(d.Body)
			if err != nil {
				fmt.Printf("failed to handle message in custom exchange: %s, error: %s\n", exchangeName, err)
			}

			if err == nil {
				d.Ack(false)
			}
		}
	}()

	log.Printf(" [consumerClient] Waiting for messages. To exit press CTRL+C")
	<-forever
}
