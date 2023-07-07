package cjwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("abc123")
var validDuration = 3 * time.Hour

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(validDuration)

	claims := MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(token string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token is invalid")
	}
}
