package pkg

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Roles string `json:"roles"`
	jwt.StandardClaims
}
