package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"golang-postgresql-restapi/config"
)

func runHTTPServer(ctx context.Context) {
	//Get environment variables
	config := config.GetConfig()
	//Get router
	router := setUpRouting(config)

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	//Run http server in another groutine
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM)
	defer stop()
	go func() {
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Printf("Failed to run http server: %v", err)
			stop()
		}
	}()

	//wait for ctx is canceled
	<-ctx.Done()
	if err := server.Shutdown(context.Background()); err != nil {
		log.Println("Failed to shutdown server.")
	} else {
		log.Println("Successfully shutdown server.")
	}
}
