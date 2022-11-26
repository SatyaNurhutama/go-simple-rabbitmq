package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Success connect to RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"CheckQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	logrus.Println(q)

	err = ch.Publish(
		"",
		"CheckQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Satya Here"),
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success to push message to queue")
}
