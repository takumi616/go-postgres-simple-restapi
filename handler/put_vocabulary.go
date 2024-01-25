package handler

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"golang-postgresql-restapi/entity"
)

type PutVocabulary struct {
	Service  VocabularyEditter
}

//A Handler which responds to PUT request 
func (p *PutVocabulary) PutVocabularyById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Decode http request body
	var vocabulary entity.Vocabulary
	err := json.NewDecoder(r.Body).Decode(&vocabulary)
	if err != nil {
		errorRes := ErrorResponse{
			Message: "Failed to decode request: " + err.Error(),
		}
		writeJsonResponse(w, http.StatusInternalServerError, errorRes)
		return
	}

	//Retrieve a map of route variables
	vars := mux.Vars(r)
	//Call a service layer's method via interface
	updated, err := p.Service.EditVocabularyById(r.Context(), &vocabulary, vars["id"])
	if err != nil {
		errorRes := ErrorResponse{
			Message: "Failed to edit a record: " + err.Error(),
		}
		writeJsonResponse(w, http.StatusInternalServerError, errorRes)
		return 
	}

	writeJsonResponse(w, http.StatusOK, updated)
}