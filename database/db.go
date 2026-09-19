package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect(databaseURL string) (*sql.DB, error) {
	pool, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}