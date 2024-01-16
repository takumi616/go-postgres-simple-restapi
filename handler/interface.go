package handler

import (
	"context"
	"golang-postgresql-restapi/entity"
)

type VocabularyAdder interface {
	AddVocabulary(ctx context.Context, vocabulary *entity.Vocabulary) (entity.Vocabulary, error)
}