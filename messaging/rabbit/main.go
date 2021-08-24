package main

import (
	"log"

	"github.com/streadway/amqp"
)

var (
	err       error
	amqpConn  *amqp.Connection
	amqpChann *amqp.Channel
)

func main() {
	amqpConn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalln("dial", err)
	}
	defer amqpConn.Close()

	amqpChann, err = amqpConn.Channel()
	if err != nil {
		log.Fatalln("open channel", err)
	}
	defer amqpChann.Close()

	que, err := amqpChann.QueueDeclare("test-queue", false, false, false, false, nil)
	if err != nil {
		log.Fatalln("QueueDeclare", err)
	}

	err = amqpChann.Publish("", "test-queue", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("kuda nill 2"),
	})
	if err != nil {
		log.Fatalln("Publish", err)
	}

	//
	log.Println("done publish...", que)
}
