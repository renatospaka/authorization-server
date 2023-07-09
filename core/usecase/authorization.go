package usecase

import (
	"github.com/renatospaka/authorization-server/core/repository"
)

type AuthorizationUsecase struct {
	repo repository.AuthorizationInterface
}

func NewAuthorizationUsecase(repo repository.AuthorizationInterface) *AuthorizationUsecase {
	return &AuthorizationUsecase{
		repo: repo,
	}
}
