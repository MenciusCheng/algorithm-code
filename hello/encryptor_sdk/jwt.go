package encryptor_sdk

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type KeyClaims struct {
	SecretKey string `json:"secretKey"` // 密钥
	jwt.RegisteredClaims
}

func generateJwtToken(secretKey string, jwtSecret []byte, duration time.Duration) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(duration)

	claims := KeyClaims{
		SecretKey: secretKey,
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

func parseJwtToken(token string, jwtSecret []byte) (*KeyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &KeyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*KeyClaims); ok && tokenClaims.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token is invalid")
	}
}
