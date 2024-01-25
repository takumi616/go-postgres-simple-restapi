package store

import (
	"context"
	"log"
	"database/sql"
	"golang-postgresql-restapi/entity"
)

type UpdateVocabulary struct {
	DbHandle  *sql.DB
}

func (u *UpdateVocabulary) UpdateVocabularyById(ctx context.Context, vocabulary *entity.Vocabulary, vocabularyId int) (entity.Vocabulary, error) {
	//Begin a transaction
	tx, err := u.DbHandle.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Failed to start a transaction: %v", err)
		return entity.Vocabulary{}, err
	}

	//Execute update query and fetch an updated record
	var updatedRow entity.Vocabulary
	query :="UPDATE vocabulary SET title = $2, sentence = $3, meaning = $4 WHERE vocabulary.Id = $1 RETURNING *"
	err = tx.QueryRowContext(ctx, query, vocabularyId, vocabulary.Title, vocabulary.Sentence, vocabulary.Meaning).Scan(&updatedRow.Id, &updatedRow.Title, &updatedRow.Sentence, &updatedRow.Meaning)
	if err != nil {
		//Execute roll back
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("Failed to rollback transaction: %v", rollbackErr)
		}
		log.Printf("Rolled back transaction: %v", err)
		return updatedRow, err
	}

	//Commit transaction
	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return updatedRow, err
	}

	return updatedRow, nil
}

