package handler

import (
	"context"
	"golang-postgresql-restapi/entity"
)

type VocabularyFetcher interface {
	FetchAllVocabularies(ctx context.Context) ([]entity.Vocabulary, error)
	FetchVocabularyById(ctx context.Context, id string) (entity.Vocabulary, error)
}

type VocabularyAdder interface {
	AddVocabulary(ctx context.Context, vocabulary *entity.Vocabulary) (entity.Vocabulary, error)
}