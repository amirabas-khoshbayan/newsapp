package entity

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	PhoneNumber string `json:"phone_number"`
	UserID      string `json:"user_id"`
	jwt.StandardClaims
}
