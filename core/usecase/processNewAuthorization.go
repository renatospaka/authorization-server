package usecase

import (
	"log"

	"github.com/renatospaka/authorization-server/core/dto"
	"github.com/renatospaka/authorization-server/core/entity"
	"github.com/renatospaka/authorization-server/utils/dateTime"
)

func (a *AuthorizationUsecase) ProcessNewAuthorization(auth *dto.AuthorizationProcessDto) (*dto.AuthorizationProcessResultDto, error) {
	log.Println("usecase.authorizations.processNewAuthorization")

	response := &dto.AuthorizationProcessResultDto{
		TransactionID: auth.TransactionID,
		ClientID: auth.ClientID,
		Value: auth.Value,
	}

	authorization, err := entity.NewAuthorization(auth.ClientID, auth.TransactionID, auth.Value)
	if err != nil {
		response.ID = ""
		response.ErrorMessage = err.Error()
		return response, err
	}
	
	authorization.Process()
	response.ID = authorization.GetID()
	response.Status = authorization.GetStatus()
	response.DeniedAt = dateTime.FormatDateToNull(authorization.DeniedAt())
	response.ApprovedAt = dateTime.FormatDateToNull(authorization.ApprovedAt())
	if !authorization.IsValid() {
		response.ErrorMessage = err.Error()
		return response, err
	}

	_, err = a.repo.FindTransactionById(auth.TransactionID)
	// The transaction already existis in the database and that cannot happen
	if err == nil {
		response.ID = ""
		response.ErrorMessage = entity.ErrTransactionIDAlreadyInUse.Error()
		return response, entity.ErrTransactionIDAlreadyInUse
	}
	// that's OK not finding the transaction, however, if it is another erro, then...
	if err != entity.ErrAuthorizationIDNotFound {
		response.ID = ""
		response.ErrorMessage = err.Error()
		return response, err
	}
	
	err = a.repo.SaveProcessNewAuthorization(authorization)
	if err != nil {
		response.ErrorMessage = err.Error()
		return response, err
	}
	return response, nil
}
