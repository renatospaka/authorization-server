package service

import (
	"context"
	"log"

	"github.com/renatospaka/authorization-server/adapter/grpc/pb"
	"github.com/renatospaka/authorization-server/core/dto"
)

// Reprocess an existing authorization request and return to the gRPC caller
func (a *AuthorizationService) Reprocess(ctx context.Context, in *pb.AuthorizationReprocessRequest) (*pb.AuthorizationReprocessResponse, error) {
	log.Println("service.authorizations.reprocess")

	auth := &dto.AuthorizationReprocessDto{
		AuthorizationID: in.AuthorizationId,
		ClientID:        in.ClientId,
		TransactionID:   in.TransactionId,
		Value:           in.Value,
	}
	authResponse := &pb.AuthorizationReprocessResponse{}

	response, err := a.usecases.ReprocessPendingAuthorization(auth)
	authResponse = &pb.AuthorizationReprocessResponse{
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
