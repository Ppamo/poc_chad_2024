package main

import (
	"fmt"
	"log"
	"net/http"
)

const address = ":8080"

func main() {
	log.Printf("> starting server at %s", address)
	http.HandleFunc("/", Handler)
	http.ListenAndServe(address, nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("> Handeling request from agent %s", r.UserAgent())
	fmt.Fprintf(w, `{"chad": "ok"}`)
}
