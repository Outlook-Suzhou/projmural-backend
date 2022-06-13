package http

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	MicrosoftId string
	jwt.StandardClaims
}

func GenerateJWT(MicrosoftId string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(JWT_EXPIRE_SECOND * time.Second)
	claims := Claims{
		MicrosoftId: MicrosoftId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    JWT_ISSUER,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JWT_SECRET))
	return token, err
}

func ParseJWT(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
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
