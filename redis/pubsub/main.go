package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	ctx            = context.Background()
	channel string = "SOME_PUB_SUB"
	redConn *redis.Client
	pubSub  *redis.PubSub
)

func main() {
	log.Println("open connection")
	redConn = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	log.Println("check connection")
	if _, err := redConn.Ping(ctx).Result(); err != nil {
		log.Fatalln("redis ping", err)
	}

	log.Println("pubsub channel")
	pubSub = redConn.Subscribe(ctx, channel)
	if _, err := pubSub.Receive(ctx); err != nil {
		log.Fatalln("pub/sub Receive", err)
	}

	log.Println("send / publish message to topic")
	if err := publish(); err != nil {
		log.Fatalln("error subscribe", err)
	}

	log.Println("get / pull / subscibe message from topic")
	if err := subscribe(); err != nil {
		log.Fatalln("error subscribe", err)
	}

	log.Println("finish operation")
	pubSub.Close()
	redConn.Close()
	log.Println("success")
}

func publish() error {
	counter := 0
	for counter < 10 {
		counter++

		s := fmt.Sprintf("%d", counter)
		if err := redConn.Publish(ctx, channel, "message "+s).Err(); err != nil {
			return nil
		}

		// time.Sleep(5 * time.Second)
	}

	return nil
}

func subscribe() error {
	ch := pubSub.Channel()
	counter := 0
	for msg := range ch {
		fmt.Println("data |", msg.Channel, "|", msg.Pattern, "|", msg.Payload, "|", msg.PayloadSlice, "|", msg.String())

		counter++
		if counter == 10 {
			break
		}
	}

	return nil
}
