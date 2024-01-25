package handler

import (
	"net/http"
	"github.com/gorilla/mux"
)

type DeleteVocabulary struct {
	Service  VocabularyRemover
}

//A Handler which responds to DELETE request 
func (d *DeleteVocabulary) DeleteVocabularyById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Retrieve a map of route variables
	vars := mux.Vars(r)

	//Call a service layer's method via interface
	deleted, err := d.Service.RemoveVocabularyById(r.Context(), vars["id"])
	if err != nil {
		errorRes := ErrorResponse{
			Message: "Failed to remove a record: " + err.Error(),
		}
		writeJsonResponse(w, http.StatusInternalServerError, errorRes)
		return
	}

	writeJsonResponse(w, http.StatusOK, deleted)
}