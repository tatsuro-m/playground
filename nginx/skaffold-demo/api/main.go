package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("starting...")

	http.HandleFunc("/healthz", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello. This is %s api!", os.Getenv("API_NAME"))
}
