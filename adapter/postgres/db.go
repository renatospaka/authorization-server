package repository

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/renatospaka/authorization-server/core/entity"
)

type PostgresDatabase struct {
	DB *sql.DB
}

var (
	ctx context.Context
)

func NewPostgresDatabase(db *sql.DB) *PostgresDatabase {
	ctx = context.Background()
	return &PostgresDatabase{
		DB: db,
	}
}


// Persists whatever happens during the process in the database
func (p *PostgresDatabase) Process(auth *entity.Authorization) error {
	return p.processAuthorization(ctx, auth)
}

// Finds a specific authorization by it id
func (p *PostgresDatabase) FindById(id string) (*entity.Authorization, error) {
	return p.findAuthorizationById(ctx, id)
}
