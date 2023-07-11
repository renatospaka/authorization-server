package usecase

import (
	"errors"
	"log"

	"github.com/renatospaka/authorization-server/core/dto"
	"github.com/renatospaka/authorization-server/utils/dateTime"
)

// Reprocess an Authorization for Transaction request.
// By default, the transaction cannot be approved nor denied, thus, there must not be an Authorization
// Status of the authorization, if any, MUST BE "pending"
func (a *AuthorizationUsecase) ReprocessPendingAuthorization(auth *dto.AuthorizationReprocessPendingDto) (*dto.AuthorizationReprocessPendingResultDto, error) {
	log.Println("usecase.authorizations.reprocessPendingAuthorization")

	response := &dto.AuthorizationReprocessPendingResultDto{
		ClientID:        auth.ClientID,
		TransactionID:   auth.TransactionID,
		Value:           auth.Value,
	}

	authorization, err := a.repo.FindTransactionById(auth.TransactionID)
	if err != nil {
		response.Status = "<not updated>"
		response.ErrorMessage = err.Error()
		return response, err
	}

	if auth.ClientID != authorization.GetClientID() {
		err = errors.New("authorization cannot be reprocessed, client id of the request does not match with existing")
		response.ErrorMessage = err.Error()
		return response, err
	}

	if auth.Value != authorization.GetValue() {
		err = errors.New("authorization cannot be reprocessed, value of the request does not match with existing")
		response.ErrorMessage = err.Error()
		return response, err
	}

	newStatus, err := authorization.Reprocess()
	if err != nil {
		response.Status = authorization.GetStatus()
		response.ErrorMessage = err.Error()
		return response, err
	}

	response.AuthorizationID = authorization.GetID()
	response.Status = newStatus
	response.DeniedAt = dateTime.FormatDateToNull(authorization.DeniedAt())
	response.ApprovedAt = dateTime.FormatDateToNull(authorization.ApprovedAt())
	if !authorization.IsValid() {
		response.ErrorMessage = err.Error()
		return response, err
	}

	err = a.repo.SaveReprocessPendingAuthorization(authorization)
	if err != nil {
		response.ErrorMessage = err.Error()
		return response, err
	}
	return response, nil
}
