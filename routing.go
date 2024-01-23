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

	//Set up routing
	router := mux.NewRouter()
	router.HandleFunc("/api/vocabularies", m["get"].(*handler.GetVocabulary).GetAllVocabularies).Methods("GET")
	router.HandleFunc("/api/vocabularies/{id}", m["get"].(*handler.GetVocabulary).GetVocabularyById).Methods("GET")
	router.HandleFunc("/api/vocabularies", m["post"].(*handler.PostVocabulary).PostVocabulary).Methods("POST")
	router.HandleFunc("/api/vocabularies/{id}", m["put"].(*handler.PutVocabulary).PutVocabularyById).Methods("PUT")
	router.HandleFunc("/api/vocabularies/{id}", m["delete"].(*handler.DeleteVocabulary).DeleteVocabularyById).Methods("DELETE")

	return router
}

func getHandlerMap(dbHandle *sql.DB) (map[string]interface{}) {
	m := make(map[string]interface{})

	get := &handler.GetVocabulary{
		Service: &service.FetchVocabulary{
			Store: &store.SelectVocabulary{
				DbHandle: dbHandle,
			},
		},
	}
	m["get"] = get

	post := &handler.PostVocabulary{
		Service: &service.AddVocabulary{
			Store: &store.InsertVocabulary{
				DbHandle: dbHandle,
			},
		},
	}
	m["post"] = post

	put := &handler.PutVocabulary{
		Service: &service.EditVocabulary{
			Store: &store.UpdateVocabulary{
				DbHandle: dbHandle,
			},
		},
	}
	m["put"] = put

	delete := &handler.DeleteVocabulary{
		Service: &service.RemoveVocabulary{
			Store: &store.DeleteVocabulary{
				DbHandle: dbHandle,
			},
		},
	}
	m["delete"] = delete

	return m
}