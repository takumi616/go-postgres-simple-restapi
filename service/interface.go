package service

import (
	"context"
	"golang-postgresql-restapi/entity"
)

type VocabularyInserter interface {
	InsertVocabulary(ctx context.Context, vocabulary *entity.Vocabulary) (entity.Vocabulary, error)
}