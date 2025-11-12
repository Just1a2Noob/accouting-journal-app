package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, msg string, err error, code int) {
	response := struct {
		ErrorMSG    string
		Description string
	}{
		ErrorMSG:    msg,
		Description: err.Error(),
	}

	errResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error marshalling json: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Write(errResponse)
}
