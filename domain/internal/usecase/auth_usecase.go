// Package usecase describe business logic
package usecase

import (
	"context"

	"github.com/MochamadAkbar/go-grpc-microservices/domain/internal/repository"
	authv1 "github.com/MochamadAkbar/go-grpc-microservices/stubs/auth/v1"
	"github.com/MochamadAkbar/go-grpc-microservices/stubs/auth/v1/entity"
)

type AuthUsecase interface {
	Register(ctx context.Context, user *entity.UserEntity) (*authv1.TokenData, error)
	Login(ctx context.Context, user *entity.UserEntity) (*authv1.TokenData, error)
}

type AuthUsecaseImpl struct {
	Repository repository.AuthRepository
}

func (usecase *AuthUsecaseImpl) Register(ctx context.Context, user *entity.UserEntity) (*authv1.TokenData, error) {
	res, err := usecase.Repository.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return &authv1.TokenData{
		UserId:    res.Id,
		Token:     "test",
		ExpiresIn: 12345651,
	}, nil
}

func (usecase *AuthUsecaseImpl) Login(ctx context.Context, user *entity.UserEntity) (*authv1.TokenData, error) {
	res, err := usecase.Repository.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	return &authv1.TokenData{
		UserId:    res.Id,
		Token:     "test",
		ExpiresIn: 12345651,
	}, nil
}
