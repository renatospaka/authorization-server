package repository

import (
	"context"
	"log"
	"time"

	"github.com/renatospaka/authorization-server/core/entity"
)

func (p *PostgresDatabase) processAuthorization(ctx context.Context, auth *entity.Authorization) error {
	log.Println("repository.authorizations.processAuthorization")

	var approvedAt, deniedAt, createdAt, updatedAt, deletedAt interface{}

	approvedAt = auth.ApprovedAt().Format(time.UnixDate)
	if auth.ApprovedAt().IsZero() {
		approvedAt = nil
	}

	deniedAt = auth.DeniedAt().Format(time.UnixDate)
	if auth.DeniedAt().IsZero() {
		deniedAt = nil
	}

	createdAt = auth.CreatedAt().Format(time.UnixDate)
	if auth.CreatedAt().IsZero() {
		createdAt = nil
	}

	updatedAt = auth.UpdatedAt().Format(time.UnixDate)
	if auth.UpdatedAt().IsZero() {
		updatedAt = nil
	}

	deletedAt = auth.DeletedAt().Format(time.UnixDate)
	if auth.DeletedAt().IsZero() {
		deletedAt = nil
	}

	query := `
	INSERT INTO authorizations
		(id, client_id, transaction_id, status, value, approved_at, denied_at, created_at, updated_at, deleted_at) 
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		auth.GetID(),
		auth.GetClientID(),
		auth.GetTransactionID(),
		auth.GetStatus(),
		auth.GetValue(),
		approvedAt,
		deniedAt,
		createdAt,
		updatedAt,
		deletedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
