package usecase

import (
	"errors"
	"log"

	"github.com/renatospaka/authorization-server/core/dto"
	"github.com/renatospaka/authorization-server/utils/dateTime"
)

// Reprocess a Transaction Process request.
// By default, the transaction cannot be approved nor denied, thus, there must not be an Authorization
// Status of the authorization, if any, MUST BE "pending"
func (a *AuthorizationUsecase) ReprocessPendingAuthorization(auth *dto.AuthorizationReprocessDto) (*dto.AuthorizationReprocessResultDto, error) {
	log.Println("usecase.authorizations.reprocessPendingAuthorization")

	response := &dto.AuthorizationReprocessResultDto{
		AuthorizationID: auth.AuthorizationID,
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

	if auth.AuthorizationID != authorization.GetID() {
		err = errors.New("authorization cannot be reprocessed, authorization id of the request does not match with existing")
		response.ErrorMessage = err.Error()
		return response, err
	}

	if auth.ClientID != authorization.GetClientID() {
		err = errors.New("authorization cannot be reprocessed, client id of the request does not match with existing")
		response.ErrorMessage = err.Error()
		return response, err
	}

	newStatus, err := authorization.Reprocess()
	if err != nil {
		response.Status = authorization.GetStatus()
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

	err = a.repo.Reprocess(authorization)
	if err != nil {
		response.ErrorMessage = err.Error()
		return response, err
	}
	return response, nil
}
