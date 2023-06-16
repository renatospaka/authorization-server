package grpcServer

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	ctx    context.Context
	Server *grpc.Server
}

func NewGrpcServer(ctx context.Context) *GrpcServer {
	log.Println("iniciando servidor gRPC")

	srv := &GrpcServer{
		ctx: ctx,
	}
	return srv
}

func (g *GrpcServer) Connect(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Panic(err)
	}
	defer lis.Close()

	g.Server = grpc.NewServer()
	if err := g.Server.Serve(lis); err != nil {
		log.Panic(err)
	}
}
