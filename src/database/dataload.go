package database

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"jlpt/src/core/models"

	_ "github.com/lib/pq"
)

const (
	jlpt1 = "./seed/jlpt_n1.csv"
	jlpt2 = "./seed/jlpt_n2.csv"
	jlpt3 = "./seed/jlpt_n3.csv"
	jlpt4 = "./seed/jlpt_n4.csv"
	jlpt5 = "./seed/jlpt_n5.csv"
)

var CSVs = []string{
	jlpt1,
	jlpt2,
	jlpt3,
	jlpt4,
	jlpt5,
}

func onFailure(DB *sql.DB, warn string, err error) {
	if err != nil {
		log.Fatal(warn+"\n", err)
	}
}

func (d *DBClientImpl) Dataload(db *sql.DB) error {

	if CheckDB(db) {
		fmt.Println("DB already has data, so we're skipping the dataload")
		return nil
	}

	fmt.Println("CREATING TABLE")
	_, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS words (id SERIAL PRIMARY KEY,foreign1 VARCHAR(255) NOT NULL, foreign2 VARCHAR(255) NOT NULL,definitions TEXT NOT NULL,level  VARCHAR(255) NOT NULL)")

	if err != nil {
		fmt.Println("skip the dataload, the table exists")
		return nil
	}

	for _, csvFile := range CSVs {
		openFile, err := os.Open(csvFile)
		onFailure(db, "Error opening file", err)

		r := csv.NewReader(openFile)
		r.LazyQuotes = true

		fileData, err := r.ReadAll()
		if err != nil {
			return fmt.Errorf("error in csvReadAll func: %s", err.Error())
		}

		fmt.Printf("STARTING IMPORT FOR %s\n", csvFile)

		for i, record := range fileData {
			if i == 0 {
				continue
			} else {
				word := models.VocabWord{
					Foreign1:    record[0],
					Foreign2:    record[1],
					Definitions: record[2],
					Level:       record[3],
				}

				word.DbCreateWord(db)
			}
		}
	}

	fmt.Println("SUCCESSFUL INITIAL DB LOAD")

	return nil
}

func CheckDB(DB *sql.DB) bool {
	var (
		err    error
		exists = true
		id     int8
	)

	err = DB.QueryRow("SELECT id FROM jlpt.words WHERE id = 1").Scan(&id)

	if err != nil || id != 1 {
		exists = false
	}

	return exists
}
