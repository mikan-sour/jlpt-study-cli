package database

import (
	"database/sql"
	"jlpt/src/config"
)

type DBClient interface {
	Initialize() (*sql.DB, error)
	Dataload() error
}

type DBClientImpl struct {
	host         string
	port         string
	user         string
	password     string
	databaseName string
}

func New(cfg config.Config) (*sql.DB, error) {

	client := &DBClientImpl{
		host:         cfg.DB_Host,
		port:         cfg.DB_Port,
		user:         cfg.DB_Username,
		password:     cfg.DB_Password,
		databaseName: cfg.DB_Name,
	}

	db, err := client.Initialize()
	if err != nil {
		return nil, err
	}

	err = client.Dataload(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
