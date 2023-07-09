package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/renatospaka/authorization-server/core/entity"
	pkgEntity "github.com/renatospaka/authorization-server/utils/entity"
)

func (p *PostgresDatabase) findAuthorizationById(ctx context.Context, id string) (*entity.Authorization, error) {
	log.Println("repository.authorizations.findAuthorizationById")

	query := `
	SELECT a.id as authorization_id, 
				a.client_id, 
				a.transaction_id,
				a.status, 
				a.approved_at, 
				a.denied_at,
				a.value,
				a.created_at,
				a.updated_at,
				a.deleted_at
	FROM authorizations a 
	WHERE a.id = $1`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows := stmt.QueryRow(id)
	var (
		authorization_id sql.NullString
		transaction_id   sql.NullString
		client_id        sql.NullString
		status           sql.NullString
		approved_at      sql.NullTime
		denied_at        sql.NullTime
		value            sql.NullFloat64
		created_at       sql.NullTime
		updated_at       sql.NullTime
		deleted_at       sql.NullTime
	)
	err = rows.Scan(&authorization_id, &transaction_id, &client_id, &status, &approved_at, &denied_at, &value, &created_at, &updated_at, &deleted_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrAuthorizationIDNotFound
		}
		return nil, err
	}

	authData := &entity.AuthorizationExisting{
		DeniedAt:      denied_at.Time,
		ApprovedAt:    approved_at.Time,
		TrailDate:     pkgEntity.NewTrailDate(),
		ID:            authorization_id.String,
		ClientId:      client_id.String,
		TransactionId: transaction_id.String,
		Status:        status.String,
		Value:         float32(value.Float64),
	}

	authorization, err := entity.MountAuthorization(authData)
	if err != nil {
		return nil, err
	}
	return authorization, nil
}
