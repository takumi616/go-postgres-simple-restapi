package handler

import (
	"net/http"
)

type GetVocabulary struct {
	Service  VocabularyFetcher
}

//Get all vocabularies
func (g *GetVocabulary) GetAllVocabularies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Select all records
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