package app

import (
	grpcapp "grpc/sso/internal/app/grpc"
	"grpc/sso/internal/services/auth"
	"grpc/sso/storage/sqlite"
	"log/slog"
	"time"
)

type App struct {
	log     *slog.Logger
	GRPCSrv *grpcapp.App
	port    int
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	store, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, store, store, store, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
