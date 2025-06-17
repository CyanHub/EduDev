package auth

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type Claims struct {
	UserID uint     `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.StandardClaims
}



func GenerateToken(userID uint, roles []string) (string, error) {
	claims := Claims{
		UserID: userID,
		Roles:  roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("server.jwt_secret")))
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("server.jwt_secret")), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// HashPassword 对密码进行哈希处理
func HashPassword(password string) string {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        panic(err)
    }
    return string(bytes)
}

// CheckPasswordHash 验证密码哈希
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}