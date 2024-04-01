package authenticationservice

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"newsapp/entity"
	"strings"
	"time"
)

type Config struct {
	ExpireDuration time.Duration `yaml:"expire_duration"`
	SignKey        string        `yaml:"sign_key"`
}

type Service struct {
	Config Config
}

func New(cfg Config) Service {
	return Service{Config: cfg}
}
func (s Service) CreateToken(user entity.User) (string, error) {
	token, err := s.createToken(user.ID, user.PhoneNumber, user.Role, s.Config.ExpireDuration)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s Service) createToken(userID, phoneNumber string, role entity.Role, expireDuration time.Duration) (string, error) {

	claims := &JwtClaims{
		PhoneNumber:      phoneNumber,
		UserID:           userID,
		Role:             role,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration))},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signingString, err := token.SignedString([]byte(s.Config.SignKey))
	if err != nil {
		return "", err
	}

	return signingString, nil
}
func (s Service) ParseToken(headerToken string) (*JwtClaims, error) {
	tokenStr := strings.Replace(headerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config.SignKey), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok {
		fmt.Printf("%v %v", claims.UserID, claims.RegisteredClaims.ExpiresAt)
		return claims, nil
	} else {
		return nil, err
	}

}
