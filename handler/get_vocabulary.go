package handler

import (
	"net/http"
	"github.com/gorilla/mux"
)

type GetVocabulary struct {
	Service  VocabularyFetcher
}

//A Handler which responds to GET request (Get all data)
func (g *GetVocabulary) GetAllVocabularies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Call a service layer's method via interface
	selected, err := g.Service.FetchAllVocabularies(r.Context())
	if err != nil {
		errorRes := ErrorResponse{
			Message: "Failed to fetch records: " + err.Error(),
		}
		writeJsonResponse(w, http.StatusInternalServerError, errorRes)
		return
	}
	
	writeJsonResponse(w, http.StatusOK, selected)
}

//A Handler which responds to GET request (Get single data by id)
func (g *GetVocabulary) GetVocabularyById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Retrieve a map of route variables
	vars := mux.Vars(r)
	//Call a service layer's method via interface
	selected, err := g.Service.FetchVocabularyById(r.Context(), vars["id"])
	if err != nil {
		errorRes := ErrorResponse{
			Message: "Failed to fetch a record: " + err.Error(),
		}
		writeJsonResponse(w, http.StatusInternalServerError, errorRes)
		return
	}

	writeJsonResponse(w, http.StatusOK, selected)
}