package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "secret"

// JWT GENERATING
func GenerateJWT(issuer string) (string, error) {

	// JWT initializations
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	return claims.SignedString([]byte(SecretKey))
}

// PARSING JWT
func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}
