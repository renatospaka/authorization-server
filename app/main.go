package main

import (
	"context"
	"database/sql"
	"log"

	// "net/http"

	// httpServer "github.com/renatospaka/authorization-server/adapter/httpServer"
	postgres "github.com/renatospaka/authorization-server/adapter/postgres"
	"github.com/renatospaka/authorization-server/adapter/postgres/grpcServer"
	// "github.com/renatospaka/authorization-server/adapter/rest/controller"
	// "github.com/renatospaka/authorization-server/core/usecase"
	"github.com/renatospaka/authorization-server/utils/configs"
)

func main() {
	log.Println("iniciando a aplicação")
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Panic(err)
	}

	ctx := context.Background()
	
	//open connection to the database
	log.Println("iniciando conexão com o banco de dados")
	conn := "postgresql://" + configs.DBUser + ":" + configs.DBPassword + "@" + configs.DBHost + ":" + configs.DBPort + "/" + configs.DBName + "?sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	postgres.NewPostgresDatabase(db)
	// usecases := usecase.NewTransactionUsecase(repo)
	// controllers := controller.NewTransactionController(usecases)
	// webServer := httpServer.NewHttpServer(ctx, controllers)
	grpcSrv := grpcServer.NewGrpcServer(ctx)

	//start web server
	log.Println("autorizador de transações escutando porta:", configs.GRPCServerPort)
	grpcSrv.Connect( configs.GRPCServerPort)
}
