package service

import (
	"context"
	"golang-postgresql-restapi/entity"
)

type VocabularySelecter interface {
	SelectAllVocabularies(ctx context.Context) ([]entity.Vocabulary, error)
	SelectVocabularyById(ctx context.Context, vocabularyId int) (entity.Vocabulary, error)
}

type VocabularyInserter interface {
	InsertVocabulary(ctx context.Context, vocabulary *entity.Vocabulary) (entity.Vocabulary, error)
}

type VocabularyDeleter interface {
	DeleteVocabularyById(ctx context.Context, vocabularyId int) (entity.Vocabulary, error)
}