package utils

import (
	"crawl_data/conf"
	"crawl_data/pkg/model"
	"log"

	"github.com/streadway/amqp"
)

// Here we set the way error messages are displayed in the terminal.
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func Connect() *amqp.Connection {
	conn, err := amqp.Dial(conf.LoadEnv().AMQP_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func Message(productPreview []model.ProductView) {

	CreateQueue("vascara_crawl", productPreview)

}

func CreateQueue(nameQueue string, url []model.ProductView) {
	conn := Connect()
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// We create a Queue to send the message to.
	q, err := ch.QueueDeclare(
		nameQueue, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// We set the payload for the message.
	for _, v := range url {
		body := v.Detaillink
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		// If there is an error publishing the message, a log will be displayed in the terminal.
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Congrats, sending message: %s", body)
	}
}
