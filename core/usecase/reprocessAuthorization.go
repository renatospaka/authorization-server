package usecase

import (
	"log"

	"github.com/renatospaka/authorization-server/core/dto"
	"github.com/renatospaka/authorization-server/core/entity"
	"github.com/renatospaka/authorization-server/utils/dateTime"
)

func (a *AuthorizationUsecase) ReprocessAuthorization(auth *dto.AuthorizationReprocessDto) (*dto.AuthorizationReprocessResultDto, error) {
	log.Println("usecase.authorizations.reprocessAuthorization")

	response := &dto.AuthorizationReprocessResultDto{
		ClientID:      auth.ClientID,
		TransactionID: auth.TransactionID,
		Value:         auth.Value,
	}

	authorization, err := entity.NewAuthorization(auth.ClientID, auth.TransactionID, auth.Value)
	if err != nil {
		response.ErrorMessage = err.Error()
		return response, err
	}

	authorization.Process()
	response.AuthorizationID = authorization.GetID()
	response.Status = authorization.GetStatus()
	response.DeniedAt = dateTime.FormatDateToNull(authorization.DeniedAt())
	response.ApprovedAt = dateTime.FormatDateToNull(authorization.ApprovedAt())
	if !authorization.IsValid() {
		response.ErrorMessage = err.Error()
		return response, err
	}

	err = a.repo.Process(authorization)
	if err != nil {
		response.ErrorMessage = err.Error()
		return response, err
	}
	return response, nil
}
