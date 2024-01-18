package service

import (
	"context"
	"log"
	"strconv"
	"golang-postgresql-restapi/entity"
)

type FetchVocabulary struct {
	Store  VocabularySelecter
}

//Fetch all vocabularies
func (f *FetchVocabulary) FetchAllVocabularies(ctx context.Context) ([]entity.Vocabulary, error) {
	return f.Store.SelectAllVocabularies(ctx)
}

//Fetch a record by Id
func (f *FetchVocabulary) FetchVocabularyById(ctx context.Context, id string) (entity.Vocabulary, error) {
	//Convert string type into int
	vocabularyId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Failed to get a vocabularyId: %v", err)
	}

	return f.Store.SelectVocabularyById(ctx, vocabularyId)
}
