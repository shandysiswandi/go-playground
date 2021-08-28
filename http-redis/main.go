package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var (
	redConn *redis.Client
)

const (
	keyCounter = "counter"
)

func main() {
	ctx := context.Background()

	// open connection
	redConn = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := redConn.Ping(ctx).Result(); err != nil {
		log.Fatalln("check connection", err)
	}

	if err := redConn.FlushAll(ctx).Err(); err != nil {
		log.Fatalln("reset redis", err)
	}

	// http route path
	http.HandleFunc("/incr", Incr)
	http.HandleFunc("/decr", Decr)

	// http serve
	log.Println("listen on 8088")
	http.ListenAndServe(":8088", nil)
}

// Incr is example use below
//
// curl http://localhost:8088/incr
// curl http://localhost:8088/incr?by=2
func Incr(w http.ResponseWriter, r *http.Request) {
	// URL QUERY
	var qVal int64 = 0
	q, ok := r.URL.Query()["by"]
	if ok && len(q) > 0 {
		inte, _ := strconv.Atoi(q[0])
		qVal = int64(inte)
	}

	// increment by 1
	var value int64 = 0
	var err error
	if qVal < 1 {
		value, err = redConn.Incr(r.Context(), keyCounter).Result()
		if err != nil {
			w.Write([]byte("err" + err.Error()))
		}
	} else {
		value, err = redConn.IncrBy(r.Context(), keyCounter, qVal).Result()
		if err != nil {
			w.Write([]byte("err" + err.Error()))
		}
	}

	val := fmt.Sprintf("value of %s = %d", keyCounter, value)
	w.Write([]byte(val))
}

// Decr is example use below
//
// curl http://localhost:8088/decr
// curl http://localhost:8088/decr?by=2
func Decr(w http.ResponseWriter, r *http.Request) {
	// URL QUERY
	var qVal int64 = 0
	q, ok := r.URL.Query()["by"]
	if ok && len(q) > 0 {
		inte, _ := strconv.Atoi(q[0])
		qVal = int64(inte)
	}

	// decrement by 1
	var value int64 = 0
	var err error
	if qVal < 1 {
		value, err = redConn.Decr(r.Context(), keyCounter).Result()
		if err != nil {
			w.Write([]byte("err" + err.Error()))
		}
	} else {
		value, err = redConn.DecrBy(r.Context(), keyCounter, qVal).Result()
		if err != nil {
			w.Write([]byte("err" + err.Error()))
		}
	}

	val := fmt.Sprintf("value of %s = %d", keyCounter, value)
	w.Write([]byte(val))
}
