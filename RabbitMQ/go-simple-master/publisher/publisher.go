package publisher

import (
	"os"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var C *amqp.Channel

func JoinNetWork() {
	conn, _ = amqp.Dial(os.Getenv("AMQP_URL"))
	C, _ = conn.Channel()
	C.ExchangeDeclare("ProcessImage", amqp.ExchangeHeaders, true, false, false, false, nil)
	C.ExchangeDeclare("UploadImage", amqp.ExchangeHeaders, true, false, false, false, nil)
}
