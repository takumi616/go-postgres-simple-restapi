package handler

import (
	"encoding/json"
	"net/http"
	"golang-postgresql-restapi/entity"
	"golang-postgresql-restapi/store"
)

//Post a new vocabulary
func PostVocabulary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Decode http request
	var vocabulary entity.Vocabulary
	err := json.NewDecoder(r.Body).Decode(&vocabulary)
	if err != nil {
		errorRes := ErrorResponse{
			Message: "Failed to decode request: " + err.Error(),
		}
		writeJsonResponse(w, http.StatusInternalServerError, errorRes)
		return
	}

	//Insert a new record
	inserted, err := store.InsertVocabulary(r.Context(), &vocabulary)
	if err != nil {
		errorRes := ErrorResponse{
			Message: "Failed to add a record: " + err.Error(),
		}
		writeJsonResponse(w, http.StatusInternalServerError, errorRes)
		return 
	}

	writeJsonResponse(w, http.StatusCreated, inserted)
}