package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

//Encode and write json response into response writer
func writeJsonResponse(w http.ResponseWriter, statusCode int, body any) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Printf("Failed to encode and write error message: %v", err)
		return
	}
}