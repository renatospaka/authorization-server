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


// Persists new transaction's authorization into the database
func (p *PostgresDatabase) SaveProcessNewAuthorization(auth *entity.Authorization) error {
	return p.processNewAuthorization(ctx, auth)
}

// Persists reprocessed pending transaction's authorization into the database
func (p *PostgresDatabase) SaveReprocessPendingAuthorization(auth *entity.Authorization) error {
	return p.reprocessPendingAuthorization(ctx, auth)
}

// Finds a specific authorization by it id
func (p *PostgresDatabase) FindById(id string) (*entity.Authorization, error) {
	return p.findAuthorizationById(ctx, id)
}

// Finds a specific transaction by it id
func (p *PostgresDatabase) FindTransactionById(transactionid string) (*entity.Authorization, error) {
	return p.findTransationById(ctx, transactionid)
}
