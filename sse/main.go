package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var messageChan chan string

// https://0marumaru.medium.com/server-sent-event-in-golang-fe1b7d2d11e5

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/stream", stream).Methods(http.MethodGet)
	r.HandleFunc("/send", send).Methods(http.MethodGet)

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("listen n serve on port 8080")
	log.Fatalln(s.ListenAndServe())
}

func stream(w http.ResponseWriter, r *http.Request) {
	//
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	messageChan = make(chan string)
	defer func() {
		close(messageChan)
		messageChan = nil
	}()

	flusher, _ := w.(http.Flusher)
	for {
		_, err := fmt.Fprintf(w, "data: %s\n\n", <-messageChan)
		if err != nil {
			log.Println(err)
		}
		flusher.Flush()
	}
}

func send(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("message")
	if msg == "" {
		msg = "random-message"
	}

	if messageChan != nil {
		messageChan <- msg
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}
