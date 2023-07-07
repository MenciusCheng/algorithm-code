package cjwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func NewMyToken(jwtSecret []byte, validDuration time.Duration) *MyToken {
	return &MyToken{
		jwtSecret:     jwtSecret,
		validDuration: validDuration,
	}
}

type MyToken struct {
	jwtSecret     []byte
	validDuration time.Duration
}

func (m *MyToken) GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(m.validDuration)

	claims := MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(m.jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *MyToken) ParseToken(token string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return m.jwtSecret, nil
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
