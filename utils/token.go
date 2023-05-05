package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(ttl time.Duration, payload interface{}, privateKey string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ExtractToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("Invalid header")
	}

	parts := strings.SplitN(header, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.New("Invalid header")
	}

	return parts[1], nil
}

func ValidateToken(tokenString string, publicKey string) error {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		return err
	}

	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return err
	}

	return nil
}
