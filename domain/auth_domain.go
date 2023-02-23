// Package domain is main root of service
package domain

import (
	"context"
	"github.com/MochamadAkbar/go-grpc-microservices/domain/internal/repository"
	"github.com/MochamadAkbar/go-grpc-microservices/domain/internal/usecase"
	"github.com/MochamadAkbar/go-grpc-microservices/pkg/orm"

	"github.com/MochamadAkbar/go-grpc-microservices/domain/internal/handler"
	authv1 "github.com/MochamadAkbar/go-grpc-microservices/stubs/auth/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func RegisterAuthServiceHandler(ctx context.Context, sv *runtime.ServeMux, conn *grpc.ClientConn) {
	err := authv1.RegisterAuthServiceHandler(ctx, sv, conn)
	if err != nil {
		panic(err)
	}
}

func RegisterAuthServiceServer(ctx context.Context, db *orm.Provider, sv *grpc.Server) {
	repo := repository.NewAuthRepository(db)
	uc := usecase.NewAuthUsecase(repo)
	srv := handler.NewAuthHandler(uc)
	authv1.RegisterAuthServiceServer(sv, srv)
}
