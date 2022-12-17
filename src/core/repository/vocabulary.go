package repository

import (
	"context"
	"database/sql"
	"jlpt/src/core/models"

	pg "github.com/lib/pq"
	"github.com/pkg/errors"
)

type VocabularyRepo interface {
	GetMatchingQuizVocabulary(ctx context.Context, ids models.VocabWordIds) ([]*models.VocabWord, error)
	GetMatchingQuizVocabularyIDBoundaries(ctx context.Context, level string) (int, int, error)
}

type VocabularyRepoImpl struct {
	db *sql.DB
}

func NewVocabRepo(db *sql.DB) VocabularyRepo {
	return &VocabularyRepoImpl{db}
}

const getMatchingQuizVocabularyQuery = `
SELECT id, foreign1, foreign2, definitions FROM jlpt.words WHERE id = ANY($1);
`

const GetMatchingQuizVocabularyIDBoundariesMinQuery = `
SELECT MIN(id) from jlpt.words WHERE level = $1;
`
const GetMatchingQuizVocabularyIDBoundariesMaxQuery = `
SELECT MAX(id) from jlpt.words WHERE level = $1;
`

func (v *VocabularyRepoImpl) GetMatchingQuizVocabulary(ctx context.Context, ids models.VocabWordIds) ([]*models.VocabWord, error) {
	rows, err := v.db.QueryContext(ctx, getMatchingQuizVocabularyQuery, pg.Array(ids))
	if err != nil {
		errors.Wrapf(err, "GetMatchingQuizVocabulary > error querying the DB: %s", err.Error())
		return nil, err
	}
	defer rows.Close()

	var result []*models.VocabWord

	for rows.Next() {
		var word models.VocabWord
		err := rows.Scan(&word.ID, &word.Foreign1, &word.Foreign2, &word.Definitions)
		if err != nil {
			errors.Wrapf(err, "GetMatchingQuizVocabulary > error scanning: %s", err.Error())
			return nil, err
		}
		result = append(result, &word)
	}

	return result, nil
}

func (v *VocabularyRepoImpl) GetMatchingQuizVocabularyIDBoundaries(ctx context.Context, level string) (int, int, error) {
	var (
		low, high int
		err       error
	)

	err = v.db.QueryRowContext(ctx, GetMatchingQuizVocabularyIDBoundariesMinQuery, level).Scan(&low)
	if err != nil {
		return 0, 0, err
	}

	err = v.db.QueryRowContext(ctx, GetMatchingQuizVocabularyIDBoundariesMaxQuery, level).Scan(&high)
	if err != nil {
		return 0, 0, err
	}

	return low, high, nil

}
