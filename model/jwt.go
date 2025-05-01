package model

import "github.com/golang-jwt/jwt/v4"

type GoShopClaims struct {
	jwt.RegisteredClaims
	BaseClaims
}
type BaseClaims struct {
	Username string `json:"username"`
	UserId   uint64 `json:"userId"`
	RoleId   uint64 `json:"roleId"`
}
