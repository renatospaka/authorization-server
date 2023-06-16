package repository

import (
	"database/sql"

	_ "github.com/lib/pq"

)

type PostgresDatabase struct {
	DB *sql.DB
}

func NewPostgresDatabase(db *sql.DB) *PostgresDatabase {
	return &PostgresDatabase{
		DB: db,
	}
}
