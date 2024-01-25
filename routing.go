package main

import (
	"github.com/gorilla/mux"
	"database/sql"
	"golang-postgresql-restapi/config"
	"golang-postgresql-restapi/handler"
	"golang-postgresql-restapi/service"
	"golang-postgresql-restapi/store"
)

func setUpRouting(config *config.Config) *mux.Router {
	//Get database handle 
	dbHandle := store.ConnectToPostgres(config)

	//Get handler map
	m := getHandlerMap(dbHandle)
	handleGet := m["get"].(*handler.GetVocabulary)
	handlePost := m["post"].(*handler.PostVocabulary)
	handlePut := m["put"].(*handler.PutVocabulary)
	handleDelete := m["delete"].(*handler.DeleteVocabulary)

	//Set up routing
	router := mux.NewRouter()
	router.HandleFunc("/api/vocabularies", handleGet.GetAllVocabularies).Methods("GET")
	router.HandleFunc("/api/vocabularies/{id}", handleGet.GetVocabularyById).Methods("GET")
	router.HandleFunc("/api/vocabularies", handlePost.PostVocabulary).Methods("POST")
	router.HandleFunc("/api/vocabularies/{id}", handlePut.PutVocabularyById).Methods("PUT")
	router.HandleFunc("/api/vocabularies/{id}", handleDelete.DeleteVocabularyById).Methods("DELETE")

	return router
}

//Set up embedded struct 
func getHandlerMap(dbHandle *sql.DB) (map[string]interface{}) {
	m := make(map[string]interface{})

	m["get"] = &handler.GetVocabulary{
		Service: &service.FetchVocabulary{
			Store: &store.SelectVocabulary{
				DbHandle: dbHandle,
			},
		},
	}

	m["post"] = &handler.PostVocabulary{
		Service: &service.AddVocabulary{
			Store: &store.InsertVocabulary{
				DbHandle: dbHandle,
			},
		},
	}

	m["put"] = &handler.PutVocabulary{
		Service: &service.EditVocabulary{
			Store: &store.UpdateVocabulary{
				DbHandle: dbHandle,
			},
		},
	}

	m["delete"] = &handler.DeleteVocabulary{
		Service: &service.RemoveVocabulary{
			Store: &store.DeleteVocabulary{
				DbHandle: dbHandle,
			},
		},
	}

	return m
}