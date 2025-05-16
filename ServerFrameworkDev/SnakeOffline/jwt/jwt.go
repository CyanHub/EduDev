package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "牧牛火链",
		Subject:   "Jwt For Go Web",
		NotBefore: jwt.NewNumericDate(time.Now().Add(-10)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 60)),
	})
	signedString, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	return signedString
}

func ParseToken(token string) bool {
	claims, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return false
	}
	return claims.Valid
}
