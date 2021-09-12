package rmq

import (
	"log"

	"github.com/streadway/amqp"
)

type Rabbit struct {
	Channel *amqp.Channel
	Name    string
}

type Rabbiter interface {
	Publish(string) error
	Consume() (<-chan amqp.Delivery, error)
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func NewChannelMQ(name string) *Rabbit {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	return &Rabbit{
		Channel: ch,
		Name:    q.Name,
	}
}

func (mq *Rabbit) Publish(str string) error {
	err := mq.Channel.Publish(
		"",      // exchange
		mq.Name, // routing key
		false,   // mandatory
		false,   // immediate
		//   Publishing store the client message sent to the MQ server.
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(str),
		})
	return err
}

func (mq *Rabbit) Consume() (<-chan amqp.Delivery, error) {
	msgs, err := mq.Channel.Consume(
		mq.Name, // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	return msgs, err
}
