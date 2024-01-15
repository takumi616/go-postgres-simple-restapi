package service

import (
	"context"
	"golang-postgresql-restapi/entity"
	"golang-postgresql-restapi/store"
)

//Add a new vocabulary
func AddVocabulary(ctx context.Context, vocabulary *entity.Vocabulary) (entity.Vocabulary, error) {
	return store.InsertVocabulary(ctx, vocabulary)
}