package main

import (
	"fmt"
	"net/http"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("from middleware")
		next.ServeHTTP(w, r)

	})

}
func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Print("hello vro")
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", middleware(http.HandlerFunc(hello)))
	http.ListenAndServe(":8000", mux)
}
