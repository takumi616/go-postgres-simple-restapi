package service

import (
	"context"
	"log"
	"strconv"
	"golang-postgresql-restapi/entity"
)

type RemoveVocabulary struct {
	Store  VocabularyDeleter
}

func (r *RemoveVocabulary) RemoveVocabularyById(ctx context.Context, id string) (entity.Vocabulary, error) {
	//Convert string type into int
	vocabularyId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Failed to get a vocabularyId: %v", err)
	}

	//Call a store layer's method via interface
	return r.Store.DeleteVocabularyById(ctx, vocabularyId)
}