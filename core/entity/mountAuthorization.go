package entity

import (
	"time"

	pkgEntity "github.com/renatospaka/authorization-server/utils/entity"
)

type AuthorizationExisting struct {
	ReceivedAt time.Time
	DeniedAt   time.Time
	ApprovedAt time.Time
	*pkgEntity.TrailDate
	ID            string
	ClientId      string
	TransactionId string
	Status        string
	Value         float32
}


// Recreate an existing authorization
func MountAuthorization(auth *AuthorizationExisting) (*Authorization, error) {
	authorizationUuid, err := pkgEntity.Parse(auth.ID)
	if err != nil {
		return nil, ErrInvalidID
	}

	clientUuid, err := pkgEntity.Parse(auth.ClientId)
	if err != nil {
		return nil, ErrInvalidClientID
	}

	transactionUuid, err := pkgEntity.Parse(auth.TransactionId)
	if err != nil {
		return nil, ErrInvalidTransactionID
	}

	status := auth.Status
	if status!= TR_APPROVED &&
	status!= TR_DELETED &&
	status!= TR_DENIED &&
	status!= TR_PENDING {
		status = TR_PENDING
	}

	authorization := &Authorization{
		id:            authorizationUuid,
		clientId:      clientUuid,
		transactionId: transactionUuid,
		value:         auth.Value,
		status:        status,
		deniedAt:      time.Time{},
		approvedAt:    time.Time{},
		TrailDate:     &pkgEntity.TrailDate{},
		valid:         false,
	}
	authorization.TrailDate.SetCreationToDate(auth.CreatedAt())
	authorization.TrailDate.SetAlterationToToday()
	if !auth.DeletedAt().IsZero() {
		authorization.TrailDate.SetDeletionToDate(auth.DeletedAt())
	}

	// deliver the new authorization validated
	err = authorization.Validate()
	if err != nil {
		return nil, err
	}
	return authorization, nil
}
