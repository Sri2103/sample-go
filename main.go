package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sample go -2 app here-2")
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8500", nil)
}
