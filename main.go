package main

import (
	"log"
	"net/http"
	"golang-postgresql-restapi/config"
	"golang-postgresql-restapi/handler"
	"github.com/gorilla/mux"
)

func main() {
	//Get config
	config := config.GetConfig()

	//Set up routing
	router := mux.NewRouter()
	router.HandleFunc("/api/vocabularies", handler.PostVocabulary).Methods("POST")

	//Set up server struct
	server := &http.Server{
		Addr: ":" + config.Port,
		Handler: router,
	}

	//Run http server
	log.Fatal(server.ListenAndServe())
}