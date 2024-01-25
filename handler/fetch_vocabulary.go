package handler

import (
	"net/http"
	"github.com/gorilla/mux"
)

type GetVocabulary struct {
	Service  VocabularyFetcher
}

//Get all vocabularies
func (g *GetVocabulary) GetAllVocabularies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Fetch all records
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

//Get a vocabulary by Id
func (g *GetVocabulary) GetVocabularyById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Retrieve a map of route variables
	vars := mux.Vars(r)
	//Fetch a record by Id
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