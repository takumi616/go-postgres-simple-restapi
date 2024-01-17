package service

import (
	"context"
	"golang-postgresql-restapi/entity"
)

type FetchVocabulary struct {
	Store  VocabularyFetcher
}

//Fetch all vocabularies
func (f *FetchVocabulary) FetchAllVocabularies(ctx context.Context) ([]entity.Vocabulary, error) {
	return f.Store.SelectAllVocabularies(ctx)
}
