// Package usecase describe business logic
package usecase

import (
	"context"
	"github.com/MochamadAkbar/go-grpc-microservices/domain/internal/repository"
	"github.com/MochamadAkbar/go-grpc-microservices/pkg/jwtio"
	"github.com/MochamadAkbar/go-grpc-microservices/ssl"
	authv1 "github.com/MochamadAkbar/go-grpc-microservices/stubs/auth/v1"
	"github.com/MochamadAkbar/go-grpc-microservices/stubs/auth/v1/entity"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type AuthUsecase interface {
	Register(ctx context.Context, user *entity.UserEntity) (*authv1.TokenData, error)
	Login(ctx context.Context, user *entity.UserEntity) (*authv1.TokenData, error)
}

type AuthUsecaseImpl struct {
	Repository   repository.AuthRepository
	JsonWebToken *jwtio.JSONWebToken
}

func NewAuthUsecase(repository repository.AuthRepository) AuthUsecase {
	privateKey, err := ssl.SSLCreds.ReadFile("keys/id_rsa.key")
	if err != nil {
		panic(err)
	}

	publicKey, err := ssl.SSLCreds.ReadFile("keys/id_rsa.pub")
	if err != nil {
		panic(err)
	}

	jsonwebtoken := jwtio.NewJSONWebToken(privateKey, publicKey)

	return &AuthUsecaseImpl{Repository: repository, JsonWebToken: jsonwebtoken}
}

func (usecase *AuthUsecaseImpl) Register(ctx context.Context, user *entity.UserEntity) (*authv1.TokenData, error) {
	res, err := usecase.Repository.FindByEmail(ctx, user.GetEmail())
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "Conflict : email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server : %s", err.Error())
	}

	user.Password = string(hash)

	res, err = usecase.Repository.Register(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server : %s", err.Error())
	}

	dat := make(map[string]string)
	dat["userId"] = res.Id

	token, err := usecase.JsonWebToken.Generate(time.Minute, dat)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "JWT: %s", err.Error())
	}

	return &authv1.TokenData{
		AccessToken: token,
	}, nil
}

func (usecase *AuthUsecaseImpl) Login(ctx context.Context, user *entity.UserEntity) (*authv1.TokenData, error) {
	res, err := usecase.Repository.FindByEmail(ctx, user.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Not Found: %s", err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.GetPassword()), []byte(user.GetPassword()))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Unauthorized: %s", err.Error())
	}

	dat := make(map[string]string)
	dat["userId"] = res.Id

	token, err := usecase.JsonWebToken.Generate(time.Minute, dat)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "JWT: %s", err.Error())
	}

	return &authv1.TokenData{
		AccessToken: token,
	}, nil
}