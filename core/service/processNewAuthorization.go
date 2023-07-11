package service

import (
	"context"
	"log"

	"github.com/renatospaka/authorization-server/adapter/grpc/pb"
	"github.com/renatospaka/authorization-server/core/dto"
)

// Process the authorization request and return to the gRPC caller
func (a *AuthorizationService) ProcessNewAuthorization(ctx context.Context, in *pb.AuthorizationProcessNewRequest) (*pb.AuthorizationProcessNewResponse, error) {
	log.Println("service.authorizations.processNewAuthorization")

	auth := &dto.AuthorizationProcessDto{
		Value:         in.Value,
		ClientID:      in.ClientId,
		TransactionID: in.TransactionId,
	}
	authResponse := &pb.AuthorizationProcessNewResponse{}

	response, err := a.usecases.ProcessNewAuthorization(auth)
	authResponse = &pb.AuthorizationProcessNewResponse{
		AuthorizationId: response.ID,
		ClientId:        response.ClientID,
		TransactionId:   response.TransactionID,
		Status:          response.Status,
		Value:           response.Value,
	}

	if err != nil {
		authResponse.ErrorMessage = response.ErrorMessage
		return authResponse, nil
	}
	return authResponse, nil
}
