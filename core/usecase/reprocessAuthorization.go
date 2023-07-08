package usecase

import (
	"fmt"
	"log"

	"github.com/renatospaka/authorization-server/core/dto"
	"github.com/renatospaka/authorization-server/core/entity"
	"github.com/renatospaka/authorization-server/utils/dateTime"
)

func (a *AuthorizationUsecase) ReprocessAuthorization(auth *dto.AuthorizationReprocessDto) (*dto.AuthorizationReprocessResultDto, error) {
	log.Println("usecase.authorizations.reprocessAuthorization")

	// These should be relying on entity validation
	// if err := validateTransactionId(auth.TransactionID); err != nil {return nil, err}
	// if err := validateClientId(auth.ClientID); err != nil {return nil, err}
	// if err := validateAuthorizationId(auth.AuthorizationID); err != nil {return nil, err}

	response := &dto.AuthorizationReprocessResultDto{
		AuthorizationID: auth.AuthorizationID,
		ClientID:        auth.ClientID,
		TransactionID:   auth.TransactionID,
		Value:           auth.Value,
	}

	authorization, err := a.repo.FindById(auth.AuthorizationID)
	if err != nil {
		response.ErrorMessage = err.Error()
		return response, err
	}

	status := authorization.GetStatus()
	if status != entity.TR_PENDING {
		err := fmt.Errorf("cannot reprocess authorization with status %s", status)
		response.ErrorMessage = err.Error()
		return response, err
	}

	newStatus, err := authorization.Reprocess()
	if err != nil {
		response.ErrorMessage = err.Error()
		return response, err
	}

	response.Status = newStatus
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
