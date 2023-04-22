package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project_3/utils"
)

type API struct {
	Message string `json:"message"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		message := API{Message: "Message from Api"}

		outbyte, err := json.Marshal(message)
		utils.CkeckError(err)

		fmt.Fprint(w, string(outbyte))
	})
	
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		message := []User{
			{ID: 1, FirstName: "Asliddin", LastName: "Tukhtasinov"},
			{ID: 2, FirstName: "Asliddin", LastName: "Tukhtasinov"},
			{ID: 3, FirstName: "Jumong", LastName: "Sasiyona"},
		}

		outbyte, err := json.Marshal(message)
		utils.CkeckError(err)

		fmt.Fprint(w, string(outbyte))
	})

	http.ListenAndServe(":8080", nil)
	fmt.Println("server port:8080")
}
