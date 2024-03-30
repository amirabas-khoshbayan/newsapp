package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"newsapp/entity"
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
	token, err := s.createToken(user.ID, user.PhoneNumber, s.Config.ExpireDuration)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s Service) createToken(userID, phoneNumber string, expireDuration time.Duration) (string, error) {

	claims := &entity.JwtClaims{
		PhoneNumber:      phoneNumber,
		UserID:           userID,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration))},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signingString, err := token.SignedString([]byte(s.Config.SignKey))
	if err != nil {
		return "", err
	}

	return signingString, nil
}
