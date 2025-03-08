package main

import (
	"context"
	crowdapiv1 "github.com/Flak34/crowd-api/internal/app/crowd/api/v1"
	"github.com/Flak34/crowd-api/internal/entrypoint"
	crowd_api_v1 "github.com/Flak34/crowd-api/internal/pb/crowd-api-v1"
	"github.com/Flak34/crowd-api/internal/pgqueue"
	project_repository "github.com/Flak34/crowd-api/internal/project/repository"
	task_repository "github.com/Flak34/crowd-api/internal/task/repository"
	task_service "github.com/Flak34/crowd-api/internal/task/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
	"github.com/rs/zerolog"
	"log/slog"
	"os"
	"os/signal"
	"riverqueue.com/riverui"
	"syscall"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"time"
)

const (
	dsn               = "postgres://postgres:postgres@localhost:5434/crowd-db"
	grpcServerAddress = "localhost:7002"
	httpServerAddress = ":7000"
	riverUIAddress    = ":7005"
	pgqWorkerTimeout  = 2 * time.Minute
	pgqWorkersCount   = 5
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set output log format for development
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	// db connection initialization
	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize pg pool")
	}
	defer dbpool.Close()
	ep := entrypoint.New(dbpool)

	// pgqueue client initialization
	pgqClient := setupPgQ(dbpool)
	log.Info().Msg("starting pgq")
	err = pgqClient.Start(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start pgq client")
	}

	// Initializing repositories
	taskRepo := task_repository.New()
	projectRepo := project_repository.New()

	// Initializing services
	taskService := task_service.New(ep, taskRepo, projectRepo, pgqClient)

	//Init riverUI server
	riverUI := setupRiverUIServer(ctx, pgqClient, dbpool)
	go func() {
		log.Info().Msg("starting riverUI server")
		if http.ListenAndServe(riverUIAddress, riverUI) != nil {
			log.Error().Err(err).Msg("failed to start river UI server")
		}
	}()

	// Setup and start gRPC server
	grpcServer, listener := setupGRPCServer(taskService)
	go func() {
		log.Info().Msg("starting gRPC server")
		if err = grpcServer.Serve(listener); err != nil {
			log.Fatal().Err(err).Msg("failed to serve grpc")
		}
	}()

	// Setup and start gRPC gateway server
	httpServer := setupHTTPServer(ctx)
	go func() {
		log.Info().Msg("starting HTTP server")
		if err = httpServer.ListenAndServe(); err != nil {
			log.Fatal().Err(err).Msg("failed to serve http")
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	// TODO graceful shutdown (http, grpc, river client with StopAndCancel() method)
	sig := <-sigCh
	log.Warn().Stringer("signal", sig).Msg("Gracefully shutting down")
}

func setupGRPCServer(taskService *task_service.Service) (*grpc.Server, net.Listener) {
	server := grpc.NewServer()
	listener, err := net.Listen("tcp", grpcServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start net listener")
	}
	crowdAPIV1Service := crowdapiv1.NewCrowdAPIV1(taskService)
	crowd_api_v1.RegisterCrowdAPIV1Server(server, crowdAPIV1Service)
	reflection.Register(server)
	return server, listener
}

func setupRiverUIServer(ctx context.Context, riverClient *river.Client[pgx.Tx], dbPool *pgxpool.Pool) *riverui.Server {
	opts := &riverui.ServerOpts{
		Client: riverClient,
		DB:     dbPool,
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		Prefix: "/riverui",
	}
	server, err := riverui.NewServer(opts)
	if err != nil {
		log.Error().Err(err).Msg("Failed to initialize riverui")
	}
	// Start the server to initialize background processes for caching and periodic queries:
	err = server.Start(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to start riverui server")
	}

	return server
}

func setupHTTPServer(ctx context.Context) *http.Server {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	if err := crowd_api_v1.RegisterCrowdAPIV1HandlerFromEndpoint(ctx, mux, grpcServerAddress, opts); err != nil {
		log.Fatal().Err(err).Msg("failed to register http handlers")
	}
	server := &http.Server{
		Addr:    httpServerAddress,
		Handler: mux,
	}
	return server
}

func setupPgQ(dbPool *pgxpool.Pool) *river.Client[pgx.Tx] {
	pgqWorkers := river.NewWorkers()
	river.AddWorker(pgqWorkers, &pgqueue.AnnotationDeadlineHandler{})
	pgqClient, err := river.NewClient(riverpgxv5.New(dbPool), &river.Config{
		JobTimeout: pgqWorkerTimeout,
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: pgqWorkersCount},
		},
		Workers: pgqWorkers,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize pgq client")
	}
	return pgqClient
}
