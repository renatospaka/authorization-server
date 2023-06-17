package service

import (
	"context"

	"github.com/renatospaka/authorization-server/adapter/grpc/pb"
	"github.com/renatospaka/authorization-server/core/dto"
	"github.com/renatospaka/authorization-server/core/usecase"
)

// handle it as it were a controller
type AuthorizationService struct {
	usecases *usecase.AuthorizationUsecase
	pb.UnimplementedAuthorizationServiceServer
}

func NewAuthorizationService(usecases *usecase.AuthorizationUsecase) *AuthorizationService {
	return &AuthorizationService{
		usecases: usecases,
	}
}

// Process the authorization request and return to the gRPC caller
func (a *AuthorizationService) Process(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	auth := &dto.AuthorizationProcessDto{
		Value: in.Value,
	}
	authResponse := &pb.AuthorizationResponse{}

	result, err := a.usecases.ProcessAuthorization(auth)
	if err != nil {
		authResponse = &pb.AuthorizationResponse{
			AuthorizationId: result.ID,
			ClientId:        "",
			TransactionId:   "",
			Status:          result.Status,
			Value:           float32(result.Value),
		}
		return authResponse, err
	}

	authResponse = &pb.AuthorizationResponse{
		AuthorizationId: result.ID,
		ClientId:        "",
		TransactionId:   "",
		Status:          result.Status,
		Value:           result.Value,
	}
	return authResponse, nil
}
