package service

import (
	"log"

	"github.com/renatospaka/authorization-server/adapter/grpc/pb"
	"github.com/renatospaka/authorization-server/core/usecase"
)

// handle it as it were a controller
type AuthorizationService struct {
	usecases *usecase.AuthorizationUsecase
	pb.UnimplementedAuthorizationServiceServer
}

func NewAuthorizationService(usecases *usecase.AuthorizationUsecase) *AuthorizationService {
	log.Println("iniciando serviços gRPC")

	return &AuthorizationService{
		usecases: usecases,
	}
}
