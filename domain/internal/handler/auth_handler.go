// Package handler is described all func to deliver method like rpc or rest
package handler

import (
	"context"
	"github.com/MochamadAkbar/go-grpc-microservices/domain/internal/usecase"
	"github.com/MochamadAkbar/go-grpc-microservices/stubs/auth/v1/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"

	authv1 "github.com/MochamadAkbar/go-grpc-microservices/stubs/auth/v1"
)

type AuthHandler struct {
	authv1.UnimplementedAuthServiceServer
	Usecase usecase.AuthUsecase
}

func NewAuthHandler(usecase usecase.AuthUsecase) authv1.AuthServiceServer {
	return &AuthHandler{Usecase: usecase}
}

func (handler *AuthHandler) Check(ctx context.Context, in *authv1.HealthCheckRequest) (*authv1.HealthCheckResponse, error) {
	return &authv1.HealthCheckResponse{Message: http.StatusText(http.StatusOK)}, nil
}

func (handler *AuthHandler) Register(ctx context.Context, in *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	user := entity.UserEntity{
		Name:     in.GetName(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}
	res, err := handler.Usecase.Register(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &authv1.RegisterResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   res,
	}, nil
}

func (handler *AuthHandler) Login(ctx context.Context, in *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	user := entity.UserEntity{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}

	res, err := handler.Usecase.Login(ctx, &user)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}

	return &authv1.LoginResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   res,
	}, nil
}