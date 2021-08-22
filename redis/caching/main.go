package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx     = context.Background()
	redConn *redis.Client
)

type data struct {
	Message string `json:"message"`
}

func main() {
	// open connection
	redConn = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// check connection
	if _, err := redConn.Ping(ctx).Result(); err != nil {
		log.Fatalln("redis ping", err)
	}

	// set data
	if err := setData("go-playground-1", `{"message":"some value"}`, time.Hour); err != nil {
		log.Fatalln("redis set", err)
	}

	// get data
	var theData data
	if err := getData("go-playground-1", &theData); err != nil {
		log.Fatalln("redis get", err)
	}

	// finish operation
	redConn.Close()
	log.Println("success", theData)
}

func getData(key string, data interface{}) error {
	val, err := redConn.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	if err == redis.Nil {
		return errors.New("key does nor exist")
	}

	return json.Unmarshal([]byte(val), data)
}

func setData(key string, data interface{}, exp time.Duration) error {
	return redConn.Set(ctx, key, data, exp).Err()
}
