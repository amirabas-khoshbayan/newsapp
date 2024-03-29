package authservice

import (
	"github.com/golang-jwt/jwt"
	"newsapp/entity"
	"time"
)

type Config struct {
	ExpireAt time.Duration `yaml:"expire_at"`
}

type Service struct {
	config Config
}

func New(cfg Config) Service {
	return Service{config: cfg}
}
func (s Service) CreateToken(user entity.User) (string, error) {
	token, err := s.createToken(user.ID, user.PhoneNumber, s.config.ExpireAt)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s Service) createToken(userID, phoneNumber string, expireAt time.Duration) (string, error) {

	claims := entity.JwtClaims{
		PhoneNumber:    phoneNumber,
		UserID:         userID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expireAt.Milliseconds()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signingString, err := token.SigningString()
	if err != nil {
		return "", err
	}

	return signingString, nil
}
