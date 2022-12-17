package database

import (
	"database/sql"
	"fmt"
)

func (db *DBClientImpl) Initialize() (*sql.DB, error) {

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db.host, db.port, db.user, db.password, db.databaseName)

	DB, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("DB CONNECTED")

	return DB, err
}
