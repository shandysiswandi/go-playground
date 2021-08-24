package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/streadway/amqp"
)

func main() {
	amqpConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalln("dial", err)
	}
	defer amqpConn.Close()

	amqpChann, err := amqpConn.Channel()
	if err != nil {
		log.Fatalln("open channel", err)
	}
	defer amqpChann.Close()

	que, err := amqpChann.QueueDeclare("test-queue", false, false, false, false, nil)
	if err != nil {
		log.Fatalln("QueueDeclare", err)
	}

	http.HandleFunc("/send", func(rw http.ResponseWriter, r *http.Request) {
		count := 1
		if qp, ok := r.URL.Query()["count"]; ok && qp[0] != "" {
			count, _ = strconv.Atoi(qp[0])
		}

		data := "default message"
		if qp, ok := r.URL.Query()["message"]; ok && qp[0] != "" {
			data = qp[0]
		} else {
			log.Println(r.URL.Query()["message"])
		}

		for i := 1; i <= count; i++ {
			// publish
			str := fmt.Sprintf("%s - %d", data, i)
			err = amqpChann.Publish("", que.Name, false, false, amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(str),
			})
			if err != nil {
				log.Fatalln("Publish", err)
			}
		}

		// write
		rw.Write([]byte("ok"))
	})

	log.Println("listen on 8088")
	http.ListenAndServe(":8088", nil)
}
