package store 

import (
	"context"
	"log"
	"database/sql"
	"golang-postgresql-restapi/entity"
)

type InsertVocabulary struct {
	DbHandle  *sql.DB
}

func (i InsertVocabulary) InsertVocabulary(ctx context.Context, vocabulary *entity.Vocabulary) (entity.Vocabulary, error) {
	//Begin a transaction
	tx, err := i.DbHandle.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Failed to start a transaction: %v", err)
		return entity.Vocabulary{}, err
	}

	//Execute insert query and fetch a new inserted record
	query := "INSERT INTO vocabulary (title, sentence, meaning) VALUES($1, $2, $3) RETURNING *"
	var inserted entity.Vocabulary 
	err = tx.QueryRowContext(ctx, query, vocabulary.Title, vocabulary.Sentence, vocabulary.Meaning).Scan(&inserted.Id, &inserted.Title, &inserted.Sentence, &inserted.Meaning)
	if err != nil {
		//Execute roll back
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("Failed to rollback transaction: %v", rollbackErr)
		}
		log.Printf("Rolled back transaction: %v", err)
		return inserted, err
	}

	//Commit transaction
	if err := tx.Commit(); err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}

	return inserted, nil
}






