// Package repository is describe layer to database
package repository

import (
	"context"

	"github.com/MochamadAkbar/go-grpc-microservices/stubs/auth/v1/entity"
	"github.com/jinzhu/gorm"
)

type AuthRepository interface {
	Register(ctx context.Context, r *entity.UserEntity) (entity.UserEntity, error)
	Login(ctx context.Context, r *entity.UserEntity) (entity.UserEntity, error)
}

type AuthRepositoryImpl struct {
	Conn *gorm.DB
}

func (repository *AuthRepositoryImpl) Register(ctx context.Context, r *entity.UserEntity) (entity.UserEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *AuthRepositoryImpl) Login(ctx context.Context, r *entity.UserEntity) (entity.UserEntity, error) {
	//TODO implement me
	panic("implement me")
}

