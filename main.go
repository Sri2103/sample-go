package main

import (
	"log"
	"net/http"

	"sample-jenkins/internal/handlers"
)

func main() {
	h := handlers.NewHandler()
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.Dummy)

	mux.HandleFunc("/ip", h.GetHome)

	mux.HandleFunc("/greet", h.Greet)

	err := http.ListenAndServe(":8500", mux)
	if err != nil {
		log.Fatal(err)
	}
}
