package repository

import (
	"context"
	"log"

	"github.com/renatospaka/authorization-server/core/entity"
)

func (p *PostgresDatabase) reprocessPendingAuthorization(ctx context.Context, auth *entity.Authorization) error {
	log.Println("repository.authorizations.reprocessPendingAuthorization")

	approvedAt := formatDateToSQL(auth.ApprovedAt())
	deniedAt := formatDateToSQL(auth.DeniedAt())
	createdAt := formatDateToSQL(auth.CreatedAt())
	updatedAt := formatDateToSQL(auth.UpdatedAt())
	deletedAt := formatDateToSQL(auth.DeletedAt())

	query := `
	UPDATE authorizations
	SET		status = $1,
				client_id = $2,
				id = $3,
				approved_at = $4, 
				denied_at = $5, 
				created_at = $6, 
				updated_at = $7, 
				deleted_at = $8
	WHERE	transaction_id = $9
	`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		auth.GetStatus(),
		auth.GetClientID(),
		auth.GetID(),
		approvedAt,
		deniedAt,
		createdAt,
		updatedAt,
		deletedAt,
		auth.GetTransactionID(),
	)
	if err != nil {
		return err
	}
	return nil
}
