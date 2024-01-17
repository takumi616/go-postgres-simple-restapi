package store

import (
	"context"
	"log"
	"database/sql"
	"golang-postgresql-restapi/entity"
)

type SelectVocabulary struct {
	DbHandle  *sql.DB
}

//Select all records
func (s *SelectVocabulary) SelectAllVocabularies(ctx context.Context) ([]entity.Vocabulary, error) {
	//Execute sql
	query := "SELECT * FROM vocabulary" 
	rows, err := s.DbHandle.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Failed to select records: %v", err)
		return nil, err
	}

	//Scan selected rows into go struct
	var selected []entity.Vocabulary
	for rows.Next() {
		vocabulary := entity.Vocabulary{}
		err = rows.Scan(&vocabulary.Id, &vocabulary.Title, &vocabulary.Sentence, &vocabulary.Meaning)
		if err != nil {
			log.Printf("Failed to scan a row into go struct: %v", err)
			return nil, err
		}
		selected = append(selected, vocabulary)
	}

	return selected, nil
}