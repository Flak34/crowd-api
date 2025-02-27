package main

import (
	"context"
	crowdapiv1 "github.com/Flak34/crowd-api/internal/app/crowd/api/v1"
	crowd_api_v1 "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

const (
	dsn               = ""
	grpcServerAddress = "localhost:7002"
	httpServerAddress = ":7000"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Setup and start gRPC server
	go func() {
		grpcServer := grpc.NewServer()
		listener, err := net.Listen("tcp", grpcServerAddress)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		crowdAPIV1Service := crowdapiv1.NewCrowdAPIV1()
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
