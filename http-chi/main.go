package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))
	r.Use(middleware.Heartbeat("/health"))
	r.Use(middleware.Compress(1))
	r.Use(middleware.CleanPath)
	r.Use(middleware.Throttle(1))

	r.Get("/", root)
	r.Get("/panic", panicHandler)
	r.Get("/timeout", timeoutHandler)
	r.Get("/rip", rip)

	log.Println("server running on port 8080")
	http.ListenAndServe(":8080", r)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to root"))
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("panic dong")
}

func timeoutHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rand.Seed(time.Now().Unix())
	processTime := time.Duration(rand.Intn(10)+1) * time.Second
	log.Println(processTime)

	select {
	case <-ctx.Done():
		w.Write([]byte("reach the limit of timeout duration"))
		return

	case <-time.After(processTime):
		w.Write([]byte("masih bisa di handel karena tidak melewati batas limit timeout"))
		return
	}
}

func rip(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("real ip " + r.RemoteAddr))
}
