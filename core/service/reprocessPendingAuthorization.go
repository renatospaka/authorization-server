package service

import (
	"context"
	"log"

	"github.com/renatospaka/authorization-server/adapter/grpc/pb"
	"github.com/renatospaka/authorization-server/core/dto"
)

// Reprocess an existing authorization request and return to the gRPC caller
func (a *AuthorizationService) ReprocessPendingAuthorization(ctx context.Context, in *pb.AuthorizationReprocessPendingRequest) (*pb.AuthorizationReprocessPendingResponse, error) {
	log.Println("service.authorizations.reprocessPendingAuthorization")

	auth := &dto.AuthorizationReprocessPendingDto{
		ClientID:        in.ClientId,
		TransactionID:   in.TransactionId,
		Value:           in.Value,
	}
	authResponse := &pb.AuthorizationReprocessPendingResponse{}

	response, err := a.usecases.ReprocessPendingAuthorization(auth)
	authResponse = &pb.AuthorizationReprocessPendingResponse{
		AuthorizationId: response.AuthorizationID,
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
