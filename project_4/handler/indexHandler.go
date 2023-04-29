package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project_4/utils"

	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type response struct {
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
	}

	urlParams := mux.Vars(r)
	id := urlParams["id"]

	byteOutput, err := json.Marshal(
		response{
			StatusCode: http.StatusOK,
			Message:    fmt.Sprintf("ID: %s of Url Params", id),
		},
	)
	utils.CheckError(err)

	w.WriteHeader(http.StatusOK)
	w.Write(byteOutput)
	return
}
