package main

import (
	"net/http"
	"project_5/handler"
)

func main() {
	http.HandleFunc("/", handler.IndexHandler)
	http.ListenAndServe(":8080", nil)
}
