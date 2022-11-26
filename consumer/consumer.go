package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	message, err := ch.Consume(
		"CheckQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range message {
			fmt.Printf("Message reveived: %s\n", d.Body)
		}
	}()

	fmt.Println("Success connected to RabbitMQ")
	fmt.Println(" [*] - waiting for messages")
	<-forever

}
