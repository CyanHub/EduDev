package utils

import (
	"errors"
	"time"

	"github.com/CyanHub/EduDev/global"
	"github.com/CyanHub/EduDev/model"
	"github.com/golang-jwt/jwt/v4"
)

type Jwt struct {
	signingKey []byte
}

func NewJwt() *Jwt {
	return &Jwt{signingKey: []byte(global.CONFIG.Jwt.Secret)}
}

func (j *Jwt) CreateClaims(baseClaims model.BaseClaims) model.GoShopClaims {
	claims := model.GoShopClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    global.CONFIG.Jwt.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.CONFIG.Jwt.ExpireTime) * time.Second)),
		},
	}
	return claims
}

func (j *Jwt) GenerateToken(claims *model.GoShopClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.signingKey)
}

var (
	TokenExpired     = errors.New("令牌过期，请重新登录")
	TokenNotValidYet = errors.New("令牌尚未生效，请稍后再试")
	TokenMalformed   = errors.New("非法的令牌")
	TokenInvalid     = errors.New("无效令牌")
)

func (j *Jwt) ParseToken(tokenString string) (*model.GoShopClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.GoShopClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.signingKey, nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*model.GoShopClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
