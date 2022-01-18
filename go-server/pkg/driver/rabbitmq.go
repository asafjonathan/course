package driver

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func RabbitChannel() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(os.Getenv("RABBIT_CONNECTION_STRING"))
	if err != nil {
		log.Fatal("There is no connection to rabbit...")
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	_, err = ch.QueueDeclare("ProductQueue", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	return conn, ch
}
