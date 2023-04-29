package main

import (
	"log"
	"net/http"
	"project_4/handler"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/index/{id}", handler.IndexHandler).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
