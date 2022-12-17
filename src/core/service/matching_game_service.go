package service

import (
	"context"
	"math/rand"
	"time"

	"jlpt/src/core/models"
	"jlpt/src/core/repository"
	"jlpt/src/utils"
)

type MatchingGameService interface {
	GetIDRange(ctx context.Context) (int, int, error)
	GetRandomIds(min, max int) models.VocabWordIds
	GetWordsForAQuestion(ctx context.Context, ids models.VocabWordIds) ([]*models.VocabWord, error)
}

type MatchingGameServiceImpl struct {
	level models.JLPTLevel
	repo  repository.VocabularyRepo
}

func NewMatchingGameService(level models.JLPTLevel, repo repository.VocabularyRepo) MatchingGameService {
	return &MatchingGameServiceImpl{level, repo}
}

func (mgs *MatchingGameServiceImpl) GetIDRange(ctx context.Context) (int, int, error) {
	return mgs.repo.GetMatchingQuizVocabularyIDBoundaries(ctx, utils.MapEnumToLevel(mgs.level))
}

func (mgs *MatchingGameServiceImpl) GetRandomIds(min, max int) models.VocabWordIds {
	rand.Seed(time.Now().UnixNano()) // important for randomizing
	var res models.VocabWordIds
	res[0] = rand.Intn(max-min) + min
	res[1] = rand.Intn(max-min) + min
	res[2] = rand.Intn(max-min) + min
	res[3] = rand.Intn(max-min) + min
	return res
}

func (mgs *MatchingGameServiceImpl) GetWordsForAQuestion(ctx context.Context, ids models.VocabWordIds) ([]*models.VocabWord, error) {
	return mgs.repo.GetMatchingQuizVocabulary(ctx, ids)
}
