package jwtio

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type JWTMethod int8

const (
	HMAC JWTMethod = iota
	Ed25519
	ECDSA
	RSA
	RSAPSS
)

func NewJWT(method jwt.SigningMethod, key []byte, isClaim bool, claims jwt.MapClaims) (string, error) {
	if isClaim {
		token := jwt.NewWithClaims(method, claims)

		tokenString, err := token.SignedString(key)
		if err != nil {
			return "", nil
		}

		return tokenString, nil
	}

	token := jwt.New(method)

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JWTParse(method JWTMethod, tokenString string, key []byte, isClaim bool, claims jwt.MapClaims) (*jwt.Token, error) {
	if isClaim {
		switch method {
		case HMAC:
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return key, nil
			})
			return token, err
		case Ed25519:
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return key, nil
			})
			return token, err
		case ECDSA:
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return key, nil
			})
			return token, err
		case RSA:
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return key, nil
			})
			return token, err
		case RSAPSS:
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSAPSS); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return key, nil
			})
			return token, err
		default:
			return &jwt.Token{}, fmt.Errorf("unknown method %v", method)
		}
	}

	switch method {
	case HMAC:
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		})
		return token, err
	case Ed25519:
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		})
		return token, err
	case ECDSA:
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		})
		return token, err
	case RSA:
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		})
		return token, err
	case RSAPSS:
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSAPSS); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		})
		return token, err
	default:
		return &jwt.Token{}, fmt.Errorf("unknown method %v", method)
	}
}

func JWTValidate(method JWTMethod, tokenString string, key []byte) (jwt.MapClaims, error) {
	token, err := JWTParse(method, tokenString, key, false, jwt.MapClaims{})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if ok {
			return claims, nil
		} else {
			return jwt.MapClaims{}, fmt.Errorf("invalid token")
		}
	} else {
		return jwt.MapClaims{}, err
	}
}
