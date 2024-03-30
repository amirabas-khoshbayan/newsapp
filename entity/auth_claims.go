package entity

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	PhoneNumber string `json:"phone_number"`
	UserID      string `json:"user_id"`
	jwt.RegisteredClaims
}

func (j JwtClaims) Valid() error {
	_, err := j.RegisteredClaims.GetExpirationTime()
	return err
}
