package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgres(dbStr string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
