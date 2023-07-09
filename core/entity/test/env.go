package test

import (
	"time"

	"github.com/renatospaka/authorization-server/core/entity"
	pkgEntity "github.com/renatospaka/authorization-server/utils/entity"
)

// Prefix "c_" assigned to not be exported to the entire project,
// only to the tests cases
const (
	c_CLIENT_ID             = "ebf77a6e-e236-441a-8ef6-4a3afdc5d4b1"
	c_CLIENT_ID_FAKE        = "00000000-0000-0000-0000-000000000000"
	c_TRANSACTION_ID        = "5b1e904f-1503-48f1-95ce-2388761cb741"
	c_TRANSACTION_ID_FAKE   = "00000000-0000-0000-0000-000000000000"
	c_AUTHORIZATION_ID      = "1d0ac9af-026d-4e90-842a-1fece70cc24a"
	c_AUTHORIZATION_ID_FAKE = "00000000-0000-0000-0000-000000000000"
)

var (
	authData  *entity.AuthorizationExisting
)

func getTestData() *entity.AuthorizationExisting {
	return &entity.AuthorizationExisting{
		ReceivedAt:    time.Time{},
		DeniedAt:      time.Time{},
		ApprovedAt:    time.Time{},
		TrailDate:     pkgEntity.NewTrailDate(),
		ID:            c_AUTHORIZATION_ID,
		ClientId:      c_CLIENT_ID,
		TransactionId: c_TRANSACTION_ID,
		Status:        entity.TR_PENDING,
		Value:         200.00,
	}
}