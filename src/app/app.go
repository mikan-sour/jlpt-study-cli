package app

import (
	"jlpt/src/config"
	"jlpt/src/core/models"
	"jlpt/src/core/repository"
	"jlpt/src/core/service"
	"jlpt/src/database"
	"jlpt/src/utils"
)

type App interface {
	SelectLevel() error
	GetLevelEnum() models.JLPTLevel
	GetLevelString() string
	GetMatchingGameService() service.MatchingGameService
}

type AppImpl struct {
	Level               models.JLPTLevel
	MatchingGameService service.MatchingGameService
}

func New(cfg config.Config) App {
	db, err := database.New(cfg)
	if err != nil {
		panic(err)
	}

	app := &AppImpl{}

	err = app.SelectLevel()
	if err != nil {
		panic(err)
	}

	app.MatchingGameService = service.NewMatchingGameService(
		app.Level,
		repository.NewVocabRepo(db),
	)

	return app
}

func (a *AppImpl) GetLevelEnum() models.JLPTLevel {
	return a.Level
}

func (a *AppImpl) GetLevelString() string {
	return utils.MapEnumToLevel(a.Level)
}

func (a *AppImpl) GetMatchingGameService() service.MatchingGameService {
	return a.MatchingGameService
}
