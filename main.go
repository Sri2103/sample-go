package main

import (
	"log"
	"net/http"
	"time"

	"sample-jenkins/internal/handlers"
)

func main() {
	h := handlers.NewHandler()
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.Dummy)

	mux.HandleFunc("/ip", h.GetHome)

	mux.HandleFunc("/greet", h.Greet)

	server := &http.Server{
		Addr:              ":8500",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	log.Println("Server to start on :8500")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
