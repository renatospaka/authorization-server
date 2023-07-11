package repository

import "github.com/renatospaka/authorization-server/core/entity"

type AuthorizationInterface interface {
	Process(*entity.Authorization) error
	SaveReprocessPendingAuthorization(*entity.Authorization) error
	FindById(string) (*entity.Authorization, error)
	FindTransactionById(string) (*entity.Authorization, error)
}
