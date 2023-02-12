// Package repository is describe layer to database
package repository

import (
	"context"
	"github.com/MochamadAkbar/go-grpc-microservices/pkg/orm"

	"github.com/MochamadAkbar/go-grpc-microservices/stubs/auth/v1/entity"
)

type AuthRepository interface {
	Register(ctx context.Context, r *entity.UserEntity) (entity.UserEntity, error)
	Login(ctx context.Context, r *entity.UserEntity) (entity.UserEntity, error)
}

type AuthRepositoryImpl struct {
	Conn *orm.Provider
}

func (repository *AuthRepositoryImpl) Register(ctx context.Context, r *entity.UserEntity) (entity.UserEntity, error) {
	return entity.UserEntity{
		Id: "test",
	}, nil
}

func (repository *AuthRepositoryImpl) Login(ctx context.Context, r *entity.UserEntity) (entity.UserEntity, error) {
	return entity.UserEntity{
		Id: "test",
	}, nil
}
