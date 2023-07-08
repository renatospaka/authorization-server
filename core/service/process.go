package service

import (
	"context"
	"log"

	"github.com/renatospaka/authorization-server/adapter/grpc/pb"
	"github.com/renatospaka/authorization-server/core/dto"
)

// Process the authorization request and return to the gRPC caller
func (a *AuthorizationService) Process(ctx context.Context, in *pb.AuthorizationProcessRequest) (*pb.AuthorizationProcessResponse, error) {
	log.Println("service.authorizations.process")

	auth := &dto.AuthorizationProcessDto{
		Value:         in.Value,
		ClientID:      in.ClientId,
		TransactionID: in.TransactionId,
	}
	authResponse := &pb.AuthorizationProcessResponse{}

	response, err := a.usecases.ProcessAuthorization(auth)
	authResponse = &pb.AuthorizationProcessResponse{
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