package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	DB *sql.DB
}

func NewPostgreSQL(connString string) (*PostgreSQL, error) {
	db, err := sql.Open("postgres", "host=localhost user=postgres port=5432 password=local sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &PostgreSQL{DB: db}, nil
}
