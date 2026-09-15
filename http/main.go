package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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
	router := mux.NewRouter()
	router.HandleFunc("/api/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		json.NewEncoder(w).Encode(vars["name"])

	}).Methods("GET")
	http.ListenAndServe(":8000", router)

}
