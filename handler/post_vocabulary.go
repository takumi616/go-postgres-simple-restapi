package handler

import (
	"encoding/json"
	"net/http"
	"golang-postgresql-restapi/entity"
)

type PostVocabulary struct {
	Service  VocabularyAdder
}

//A Handler which responds to POST request 
func (p *PostVocabulary) PostVocabulary(w http.ResponseWriter, r *http.Request) {
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

	//Call a service layer's method via interface
	inserted, err := p.Service.AddVocabulary(r.Context(), &vocabulary)
	if err != nil {
		errorRes := ErrorResponse{
			Message: "Failed to add a record: " + err.Error(),
		}
		writeJsonResponse(w, http.StatusInternalServerError, errorRes)
		return 
	}

	writeJsonResponse(w, http.StatusCreated, inserted)
}