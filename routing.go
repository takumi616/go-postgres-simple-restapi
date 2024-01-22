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
	get := &handler.GetVocabulary{
		Service: &service.FetchVocabulary{
			Store: &store.SelectVocabulary{
				DbHandle: dbHandle,
			},
		},
	}

	post := &handler.PostVocabulary{
		Service: &service.AddVocabulary{
			Store: &store.InsertVocabulary{
				DbHandle: dbHandle,
			},
		},
	}

	put := &handler.PutVocabulary{
		Service: &service.EditVocabulary{
			Store: &store.UpdateVocabulary{
				DbHandle: dbHandle,
			},
		},
	}

	delete := &handler.DeleteVocabulary{
		Service: &service.RemoveVocabulary{
			Store: &store.DeleteVocabulary{
				DbHandle: dbHandle,
			},
		},
	}

	//Set up routing
	router := mux.NewRouter()
	router.HandleFunc("/api/vocabularies", get.GetAllVocabularies).Methods("GET")
	router.HandleFunc("/api/vocabularies/{id}", get.GetVocabularyById).Methods("GET")
	router.HandleFunc("/api/vocabularies", post.PostVocabulary).Methods("POST")
	router.HandleFunc("/api/vocabularies/{id}", put.PutVocabularyById).Methods("PUT")
	router.HandleFunc("/api/vocabularies/{id}", delete.DeleteVocabularyById).Methods("DELETE")

	return router
}