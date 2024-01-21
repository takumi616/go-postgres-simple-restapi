package main

import (
	"github.com/gorilla/mux"
	"golang-postgresql-restapi/config"
	"golang-postgresql-restapi/handler"
	"golang-postgresql-restapi/service"
	"golang-postgresql-restapi/store"
)

func setUpRouting(config *config.Config) *mux.Router {
	//Get database handle 
	dbHandle := store.ConnectToPostgres(config)

	//Embed structs which implement interface
	gh := &handler.GetVocabulary{
		Service: &service.FetchVocabulary{
			Store: &store.SelectVocabulary{
				DbHandle: dbHandle,
			},
		},
	}

	ph := &handler.PostVocabulary{
		Service: &service.AddVocabulary{
			Store: &store.InsertVocabulary{
				DbHandle: dbHandle,
			},
		},
	}

	dh := &handler.DeleteVocabulary{
		Service: &service.RemoveVocabulary{
			Store: &store.DeleteVocabulary{
				DbHandle: dbHandle,
			},
		},
	}

	//Set up routing
	router := mux.NewRouter()
	router.HandleFunc("/api/vocabularies", gh.GetAllVocabularies).Methods("GET")
	router.HandleFunc("/api/vocabularies/{id}", gh.GetVocabularyById).Methods("GET")
	router.HandleFunc("/api/vocabularies", ph.PostVocabulary).Methods("POST")
	router.HandleFunc("/api/vocabularies/{id}", dh.DeleteVocabularyById).Methods("DELETE")

	return router
}