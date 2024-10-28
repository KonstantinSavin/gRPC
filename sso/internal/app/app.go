package app

import (
	grpcapp "grpc/sso/internal/app/grpc"
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
	//TODO store

	//TODO auth

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
