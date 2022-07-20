package model

import "github.com/golang-jwt/jwt"

type JwtUserClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Roles string `json:"roles"`
	jwt.StandardClaims
}
