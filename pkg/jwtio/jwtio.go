// Package jwtio is shared pkg of json web token
package jwtio

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JSONWebToken struct {
	privateKey []byte
	publicKey  []byte
}

func NewJSONWebToken(private []byte, public []byte) *JSONWebToken {
	return &JSONWebToken{
		privateKey: private,
		publicKey:  public,
	}
}

func (j *JSONWebToken) Generate(duration time.Duration, payload interface{}) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		return "", err
	}
	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["dat"] = payload                  // Our custom data.
	claims["exp"] = now.Add(duration).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()               // The time at which the token was issued.
	claims["nbf"] = now.Unix()               // The time before which the token must be disregarded.

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j *JSONWebToken) Validate(tokenString string) (jwt.MapClaims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.publicKey)
	if err != nil {
		return nil, err
	}
	signing := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	}

	token, err := jwt.Parse(tokenString, signing)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}