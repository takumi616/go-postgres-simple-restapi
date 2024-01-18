package main

import (
	"log"
	"net/http"
	"golang-postgresql-restapi/config"
	"golang-postgresql-restapi/handler"
	"golang-postgresql-restapi/service"
	"golang-postgresql-restapi/store"
	"github.com/gorilla/mux"
)

func main() {
	//Get config
	config := config.GetConfig()
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

	//Set up routing
	router := mux.NewRouter()
	router.HandleFunc("/api/vocabularies", gh.GetAllVocabularies).Methods("GET")
	router.HandleFunc("/api/vocabularies/{id}", gh.GetVocabularyById).Methods("GET")
	router.HandleFunc("/api/vocabularies", ph.PostVocabulary).Methods("POST")

	//Set up server struct
	server := &http.Server{
		Addr: ":" + config.Port,
		Handler: router,
	}

	//Run http server
	log.Fatal(server.ListenAndServe())
}