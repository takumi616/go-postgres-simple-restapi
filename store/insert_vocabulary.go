package store 

import (
	"context"
	"log"
	"golang-postgresql-restapi/config"
	"golang-postgresql-restapi/entity"
)

//Insert a new record
func InsertVocabulary(ctx context.Context, vocabulary *entity.Vocabulary) (entity.Vocabulary, error) {
	config := config.GetConfig()

	dbHandle := connectToPostgres(config)

	//Begin a transaction
	tx, err := dbHandle.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Failed to start a transaction: %v", err)
		return entity.Vocabulary{}, err
	}

	//Insert a new record and fetch it
	query := "INSERT INTO vocabulary (title, sentence, meaning) VALUES($1, $2, $3) RETURNING *"
	var insertedRow entity.Vocabulary 
	err = tx.QueryRowContext(ctx, query, vocabulary.Title, vocabulary.Sentence, vocabulary.Meaning).Scan(&insertedRow.Id, &insertedRow.Title, &insertedRow.Sentence, &insertedRow.Meaning)
	if err != nil {
		//Execute roll back
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("Failed to rollback this transaction: %v", rollbackErr)
		}
		log.Printf("Rolled back this transaction: %v", err)
		return insertedRow, err
	}

	//Commit transaction
	if err := tx.Commit(); err != nil {
		log.Fatalf("Failed to commit this transaction: %v", err)
	}

	return insertedRow, nil
}






