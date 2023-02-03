// Package handler is described all func to deliver method like rpc or rest
package handler

import (
	"context"
	"net/http"

	authv1 "github.com/MochamadAkbar/go-grpc-microservices/stubs/auth/v1"
)

type AuthHandler struct {
	authv1.UnimplementedAuthServiceServer
}

func (h *AuthHandler) Check(ctx context.Context, req *authv1.HealthCheckRequest) (*authv1.HealthCheckResponse, error) {
	return &authv1.HealthCheckResponse{Message: http.StatusText(http.StatusOK)}, nil
}

func (h *AuthHandler) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	return &authv1.LoginResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data: &authv1.TokenData{
			UserId:    "123",
			Token:     "test",
			ExpiresIn: 0,
		},
	}, nil
}
