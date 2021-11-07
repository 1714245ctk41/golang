package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	//* Define RabbitMQ server URL.
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	//* CreateRabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	//* Opening a channel to our RabbitMQ instance over
	//* the connection we have already established.
	channelRabitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabitMQ.Close()

	//* subscribing to QueueService1 for getting messages.
	messages, err := channelRabitMQ.Consume(
		"QueueService1", //* queue name
		"",              //* consumer
		true,            //* auto-ack
		false,           //* exclusive
		false,           //* no local
		false,           //* no wait
		nil,             //* arguments
	)
	if err != nil {
		log.Panicln(err)
	}

	//* Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	//* Make a channel to receive messages into infinite loop.
	forerver := make(chan bool)
	go func() {
		for message := range messages {
			//* For example, show received message in a console.
			log.Println(" > Received message: %\n", message.Body)
		}
	}()
	<-forerver
}
