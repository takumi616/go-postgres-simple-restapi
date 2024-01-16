package store

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"golang-postgresql-restapi/config"
)

//Connect to Postgresql
func ConnectToPostgres(config *config.Config) *sql.DB {
	dataSourceName := "host=" + config.DatabaseHost + " port=" + config.DatabasePort + " user=" + config.DatabaseUser + " password=" + config.DatabasePassword + " dbname=" + config.DatabaseName + " sslmode=" + config.DatabaseSSLMODE
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	return db
}

