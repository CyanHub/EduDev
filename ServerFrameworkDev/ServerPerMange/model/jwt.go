package model

import "github.com/golang-jwt/jwt/v5"

type GoShopClaims struct {
    jwt.RegisteredClaims
    BaseClaims
}
type BaseClaims struct {
    Username string `json:"username"`
    UserId   uint64 `json:"userId"`
}