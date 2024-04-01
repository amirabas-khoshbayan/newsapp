package authenticationservice

import (
	"github.com/golang-jwt/jwt/v5"
	"newsapp/entity"
)

type JwtClaims struct {
	PhoneNumber string      `json:"phone_number"`
	UserID      string      `json:"user_id"`
	Role        entity.Role `json:"role"`
	jwt.RegisteredClaims
}

func (j JwtClaims) Valid() error {
	_, err := j.RegisteredClaims.GetExpirationTime()
	return err
}
