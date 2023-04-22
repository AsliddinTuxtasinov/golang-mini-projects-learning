package main

import (
	"fmt"
	"net/http"
	"project_2/utils"
)

type Human struct {
	FirstName, LastName string
	Age                 int
}

func (hum Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hum.FirstName = "Asliddin"
	hum.LastName = "Tukhtasinov"
	hum.Age = 23

	// Parse the form data
	err := r.ParseForm()
	utils.CheckError(err)

	// Get the form data
	fmt.Println(r.Form)

	// URL path information
	fmt.Println("path: ", r.URL.Path)

	// Write HTML response
	fmt.Fprint(w, "<h1>"+hum.FirstName+"<h1>")
}

func main() {
	var hum Human
	err := http.ListenAndServe("localhost:8080", &hum)
	utils.CheckError(err)
}
