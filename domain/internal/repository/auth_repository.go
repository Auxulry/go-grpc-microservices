// Package repository is describe layer to database
package repository

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/MochamadAkbar/go-grpc-microservices/pkg/orm"
	"github.com/MochamadAkbar/go-grpc-microservices/stubs/auth/v1/entity"
)

type AuthRepository interface {
	Register(ctx context.Context, r *entity.UserEntity) (entity.UserEntity, error)
	FindByEmail(ctx context.Context, email string) (entity.UserEntity, error)
}

type AuthRepositoryImpl struct {
	DB *orm.Provider
}

func NewAuthRepository(conn *orm.Provider) AuthRepository {
	return &AuthRepositoryImpl{DB: conn}
}

func (repository *AuthRepositoryImpl) Register(ctx context.Context, r *entity.UserEntity) (entity.UserEntity, error) {
	var data entity.UserEntity

	payload := &entity.UserEntityORM{
		Name:     r.GetName(),
		Email:    r.GetEmail(),
		Password: r.GetPassword(),
	}

	res := repository.DB.WithContext(ctx).Create(payload)
	if err := res.Error; err != nil {
		return data, status.Errorf(codes.Internal, "Internal Server Error: %v", err)
	}

	data.Id = payload.Id

	return data, nil
}

func (repository *AuthRepositoryImpl) FindByEmail(ctx context.Context, email string) (entity.UserEntity, error) {
	var user entity.UserEntityORM
	var data entity.UserEntity

	res := repository.DB.WithContext(ctx).Debug().Where(&entity.UserEntityORM{Email: email}).First(&user)
	if err := res.Error; err != nil {
		return data, err
	}

	data, err := user.ToPB(ctx)
	if err != nil {
		return data, err
	}

	return data, nil
}