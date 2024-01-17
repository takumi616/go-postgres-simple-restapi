package service

import (
	"context"
	"golang-postgresql-restapi/entity"
)

type VocabularyFetcher interface {
	SelectAllVocabularies(ctx context.Context) ([]entity.Vocabulary, error)
}

type VocabularyInserter interface {
	InsertVocabulary(ctx context.Context, vocabulary *entity.Vocabulary) (entity.Vocabulary, error)
}