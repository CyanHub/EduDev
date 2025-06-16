package utils

import (
	"errors"
	"time"

	"FileSystem/global"
	"FileSystem/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrTokenExpired     = errors.New("令牌已过期")
	ErrTokenNotValidYet = errors.New("令牌尚未生效")
	ErrTokenMalformed   = errors.New("非法令牌格式")
	ErrTokenInvalid     = errors.New("无效令牌")
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(global.CONFIG.Jwt.Secret),

	}
}

// CreateToken 创建Token
func (j *JWT) CreateToken(claims model.BaseClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j.createClaims(claims))
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
		return nil, err
	}

	if token != nil {
		if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, ErrTokenInvalid
}

// createClaims 创建Claims
func (j *JWT) createClaims(baseClaims model.BaseClaims) model.CustomClaims {
	return model.CustomClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    global.CONFIG.Jwt.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.CONFIG.Jwt.ExpireTime) * time.Second)),
		},
	}
}

// GetUserID 从Gin上下文获取用户ID
func GetUserID(c *gin.Context) uint64 {
	if claims, exists := c.Get("claims"); exists {
		return claims.(*model.CustomClaims).UserID
	}
	if claims, exists := c.Get("userID"); exists {
		return claims.(uint64)
	}
	return 0
}

// package utils

// import (
// 	"errors"
// 	"time"

// 	"FileSystem/global"
// 	"FileSystem/model"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v4"
// )

// // 定义错误变量
// // 修改错误变量命名，遵循 ErrXxx 规范
// var (
// 	ErrTokenExpired     = errors.New("令牌过期，请重新登录")
// 	ErrTokenNotValidYet = errors.New("令牌尚未生效，请稍后再试")
// 	ErrTokenMalformed   = errors.New("非法的令牌")
// 	ErrTokenInvalid     = errors.New("无效令牌")
// )

// type Jwt struct {
// 	signingKey []byte
// }

// func NewJwt() *Jwt {
// 	return &Jwt{signingKey: []byte(global.CONFIG.Jwt.Secret)}
// }

// func (j *Jwt) CreateToken(claims model.BaseClaims) (string, error) {
// 	goShopClaims := j.CreateClaims(claims)
// 	return j.GenerateToken(&goShopClaims)
// }

// func GetUserID(c *gin.Context) uint64 {
// 	if claims, exists := c.Get("claims"); exists {
// 		return claims.(*model.BaseClaims).UserId
// 	}
// 	if claims, exists := c.Get("userID"); exists {
// 		return claims.(uint64)
// 	}
// 	return 0
// }

// func (j *Jwt) CreateClaims(baseClaims model.BaseClaims) model.GoShopClaims {
// 	claims := model.GoShopClaims{
// 		BaseClaims: baseClaims,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			Issuer:    global.CONFIG.Jwt.Issuer,
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.CONFIG.Jwt.ExpireTime) * time.Second)),
// 		},
// 	}
// 	return claims
// }

// func (j *Jwt) GenerateToken(claims *model.GoShopClaims) (string, error) {
// 	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.signingKey)
// }

// func (j *Jwt) ParseToken(tokenString string) (*model.GoShopClaims, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &model.GoShopClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return j.signingKey, nil
// 	})
// 	if err != nil {
// 		var ve *jwt.ValidationError
// 		if errors.As(err, &ve) {
// 			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
// 				return nil, ErrTokenMalformed
// 			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
// 				return nil, ErrTokenExpired
// 			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
// 				return nil, ErrTokenNotValidYet
// 			} else {
// 				return nil, ErrTokenInvalid
// 			}
// 		}
// 		return nil, err
// 	}
// 	if token != nil {
// 		if claims, ok := token.Claims.(*model.GoShopClaims); ok && token.Valid {
// 			return claims, nil
// 		}
// 		return nil, ErrTokenInvalid
// 	} else {
// 		return nil, ErrTokenInvalid
// 	}
// }
