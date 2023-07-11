package repository

import (
	"context"
	"log"

	"github.com/renatospaka/authorization-server/core/entity"
)

func (p *PostgresDatabase) processNewAuthorization(ctx context.Context, auth *entity.Authorization) error {
	log.Println("repository.authorizations.processNewAuthorization")

	approvedAt := formatDateToSQL(auth.ApprovedAt())
	deniedAt := formatDateToSQL(auth.DeniedAt())
	createdAt := formatDateToSQL(auth.CreatedAt())
	updatedAt := formatDateToSQL(auth.UpdatedAt())
	deletedAt := formatDateToSQL(auth.DeletedAt())

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
