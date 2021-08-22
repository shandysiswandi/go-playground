package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":8008", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	// some long process
	go doInBackground()

	w.Write([]byte("ok"))
}

func doInBackground() {
	log.Println("start")
	time.Sleep(3 * time.Second)
	log.Println("finish")
}
