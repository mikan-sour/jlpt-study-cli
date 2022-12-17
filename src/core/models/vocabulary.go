package models

import (
	"database/sql"
)

type VocabWordIds [4]int

type VocabWord struct {
	ID          int
	Foreign1    string
	Foreign2    string
	Definitions string
	Level       string
}

type VocabWordRes struct {
	ID          int      `json:"id"`
	Foreign1    string   `json:"foreign1"`
	Foreign2    string   `json:"foreign2"`
	Definitions []string `json:"definitions"`
	Level       string   `json:"level"`
}

func (p *VocabWord) DbCreateWord(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO jlpt.words(foreign1,foreign2,definitions,level) VALUES($1, $2, $3, $4) RETURNING id",
		p.Foreign1, p.Foreign2, p.Definitions, p.Level).Scan(&p.ID)

	if err != nil {
		return err
	}
	return nil
}
