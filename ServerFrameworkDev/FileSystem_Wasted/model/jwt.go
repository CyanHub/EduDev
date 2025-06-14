package model

import "github.com/golang-jwt/jwt/v4"

type BaseClaims struct {
	UserID   uint64
	Username string
	RoleID   uint64
}

type CustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

// Valid implements jwt.Claims.
func (c CustomClaims) Valid() error {
	return c.RegisteredClaims.Valid()
}


// package model

// import "github.com/golang-jwt/jwt/v4"

// type BaseClaims struct {
// 	UserId   uint64
// 	Username string
// 	RoleId   uint64
// }

// type GoShopClaims struct {
// 	BaseClaims
// 	jwt.RegisteredClaims
// }

// // 显式实现 Valid 方法，解决歧义问题
// func (g *GoShopClaims) Valid() error {
// 	// 调用 RegisteredClaims 的 Valid 方法
// 	return g.RegisteredClaims.Valid()
// }
