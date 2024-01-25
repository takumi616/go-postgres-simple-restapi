package service

import (
	"context"
	"log"
	"strconv"
	"golang-postgresql-restapi/entity"
)

type EditVocabulary struct {
	Store  VocabularyUpdater
}

func (e *EditVocabulary) EditVocabularyById(ctx context.Context, vocabulary *entity.Vocabulary, id string) (entity.Vocabulary, error) {
	//Convert string type into int
	vocabularyId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Failed to get a vocabularyId: %v", err)
	}

	//Call a store layer's method via interface
	return e.Store.UpdateVocabularyById(ctx, vocabulary, vocabularyId)
}