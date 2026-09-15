package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string
}

func handler(w http.ResponseWriter, r *http.Request) {
	msg := Message{Text: "Hello go"}
	json.NewEncoder(w).Encode(msg)

}
func userCreate(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("hello")

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/api/user", userCreate)
	http.ListenAndServe(":8000", nil)

}
