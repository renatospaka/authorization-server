package repository

import "github.com/renatospaka/authorization-server/core/entity"

type AuthorizationInterface interface {
	Process(auth *entity.Authorization) error
	FindById(id string) (*entity.Authorization, error)
}
