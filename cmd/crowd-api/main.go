package main

import (
	"context"
	crowdapiv1 "github.com/Flak34/crowd-api/internal/app/crowd/api/v1"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	crowd_api_v1 "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	project_repository "github.com/Flak34/crowd-api/internal/project/repository"
	task_repository "github.com/Flak34/crowd-api/internal/task/repository"
	task_service "github.com/Flak34/crowd-api/internal/task/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

const (
	dsn               = "postgres://postgres:postgres@localhost:5434/crowd-db"
	grpcServerAddress = "localhost:7002"
	httpServerAddress = ":7000"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to initialize pg pool: %v", err)
	}
	defer dbpool.Close()
	ep := entrypoint.New(dbpool)
	taskRepo := task_repository.New()
	projectRepo := project_repository.New()
	taskService := task_service.New(ep, taskRepo, projectRepo)

	// Setup and start gRPC server
	go func() {
		grpcServer := grpc.NewServer()
		listener, err := net.Listen("tcp", grpcServerAddress)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		crowdAPIV1Service := crowdapiv1.NewCrowdAPIV1(taskService)
		crowd_api_v1.RegisterCrowdAPIV1Server(grpcServer, crowdAPIV1Service)
		reflection.Register(grpcServer)
		if err = grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Setup and start gRPC gateway server
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	if err := crowd_api_v1.RegisterCrowdAPIV1HandlerFromEndpoint(ctx, mux, grpcServerAddress, opts); err != nil {
		log.Fatalf("failed to register http handlers: %v", err)
	}
	if err := http.ListenAndServe(httpServerAddress, mux); err != nil {
		log.Fatalf("failed to serve http: %v", err)
	}
}
