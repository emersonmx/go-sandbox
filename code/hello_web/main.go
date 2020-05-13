package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Customers")
	})

	http.ListenAndServe(":8080", r)
}
