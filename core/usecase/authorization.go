package usecase

import (
	"log"

	"github.com/renatospaka/authorization-server/core/dto"
	"github.com/renatospaka/authorization-server/core/entity"
	"github.com/renatospaka/authorization-server/core/repository"
	"github.com/renatospaka/authorization-server/utils/dateTime"
)

type AuthorizationUsecase struct {
	repo repository.AuthorizationInterface
}

func NewAuthorizationUsecase(repo repository.AuthorizationInterface) *AuthorizationUsecase {
	return &AuthorizationUsecase{
		repo: repo,
	}
}

func (a *AuthorizationUsecase) ProcessAuthorization(auth *dto.AuthorizationProcessDto) (*dto.AuthorizationProcessResultDto, error) {
	log.Println("usecase.authorizations.process")

	authorization, err := entity.NewAuthorization(auth.Value)
	if err != nil {
		return nil, err
	}
	authorization.Process()

	err = a.repo.Process(authorization)
	if err != nil {
		return nil, err
	}

	response := &dto.AuthorizationProcessResultDto{
		ID:         authorization.GetID(),
		Status:     authorization.GetStatus(),
		Value:      authorization.GetValue(),
		DeniedAt:   dateTime.FormatDateToNull(authorization.DeniedAt()),
		ApprovedAt: dateTime.FormatDateToNull(authorization.ApprovedAt()),
	}
	return response, nil
}
