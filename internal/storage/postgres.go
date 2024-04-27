package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	DB *sql.DB
}

func NewPostgreSQL(connString string) (*PostgreSQL, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return &PostgreSQL{DB: db}, nil
}
