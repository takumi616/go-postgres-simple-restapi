package handler

import (
	"net/http"
	"github.com/gorilla/mux"
)

type DeleteVocabulary struct {
	Service  VocabularyRemover
}

func (d *DeleteVocabulary) DeleteVocabularyById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Retrieve a map of route variables
	vars := mux.Vars(r)

	//Remove a record by Id
	deleted, err := d.Service.RemoveVocabularyById(r.Context(), vars["id"])
	if err != nil {
		errorRes := ErrorResponse{
			Message: "Failed to delete a record: " + err.Error(),
		}
		writeJsonResponse(w, http.StatusInternalServerError, errorRes)
		return
	}

	writeJsonResponse(w, http.StatusOK, deleted)
}