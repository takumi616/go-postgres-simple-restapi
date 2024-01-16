package service

import (
	"context"
	"golang-postgresql-restapi/entity"
)

type AddVocabulary struct {
	Store  VocabularyInserter	
}

//Add a new vocabulary
func (a *AddVocabulary) AddVocabulary(ctx context.Context, vocabulary *entity.Vocabulary) (entity.Vocabulary, error) {
	return a.Store.InsertVocabulary(ctx, vocabulary)
}