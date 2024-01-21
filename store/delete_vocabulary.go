package store

import (
	"context"
	"log"
	"database/sql"
	"golang-postgresql-restapi/entity"
)

type DeleteVocabulary struct {
	DbHandle  *sql.DB
}

func (d *DeleteVocabulary) DeleteVocabularyById(ctx context.Context, vocabularyId int) (entity.Vocabulary, error) {
	tx, err := d.DbHandle.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Failed to start a transaction: %v", err)
		return entity.Vocabulary{}, err
	}

	var deletedRow entity.Vocabulary
	query := "DELETE FROM vocabulary " + "WHERE id = $1 RETURNING *" 
	err = tx.QueryRowContext(ctx, query, vocabularyId).Scan(&deletedRow.Id, &deletedRow.Title, &deletedRow.Sentence, &deletedRow.Meaning)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("Failed to rollback this transaction: %v", rollbackErr)
		}
		log.Printf("Rolled back this transaction: %v", err)
		return deletedRow, err
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("Failed to commit this transaction: %v", err)
		return deletedRow, err
	}

	return deletedRow, nil
}