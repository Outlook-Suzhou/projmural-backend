package http

import (
	"projmural-backend/pkg/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	MicrosoftId string
	jwt.StandardClaims
}

func GenerateJWT(MicrosoftId string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.Jwt.ExpiredSeconds) * time.Second)
	claims := Claims{
		MicrosoftId: MicrosoftId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.Jwt.Issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Jwt.Secret))
	return token, err
}

func ParseJWT(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
