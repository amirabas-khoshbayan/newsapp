package userservice

import (
	"fmt"
	"newsapp/entity"
	"newsapp/param/userparam"
	"time"
)

type Repository interface {
	InsertUser(user entity.User) (entity.User, error)
	GetUserByID(id string) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
	GetUsers() ([]entity.User, error)
}
type AuthGenerator interface {
	CreateToken(user entity.User) (string, error)
}

type Service struct {
	repo Repository
	auth AuthGenerator
}

func New(repo Repository, authGenerator AuthGenerator) Service {
	return Service{repo: repo, auth: authGenerator}
}
func (s Service) GetUsers() ([]entity.User, error) {
	userList, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}
	insertUser, err := s.repo.InsertUser(entity.User{
		ID:          "",
		FirstName:   "abbas",
		LastName:    "Khoshbayan",
		PhoneNumber: "123456",
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(insertUser)

	return userList, nil
}
func (s Service) CreateNewUser(req userparam.CreateNewUserRequest) (userparam.CreateNewUserResponse, error) {
	//TODO - hash the password
	user := entity.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
		Password:     req.Password,
		RegisterDate: time.Now(),
	}

	userRes, err := s.repo.InsertUser(user)
	if err != nil {
		return userparam.CreateNewUserResponse{}, err
	}

	return userparam.CreateNewUserResponse{userparam.UserInfo{
		ID:          userRes.ID,
		PhoneNumber: userRes.PhoneNumber,
		FirstName:   userRes.FirstName,
		LastName:    userRes.LastName,
		Email:       userRes.Email,
	}}, nil
}
func (s Service) Login(req userparam.LoginRequest) (userparam.LoginResponse, error) {
	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return userparam.LoginResponse{}, err
	}

	tokenStr, err := s.auth.CreateToken(user)
	if err != nil {
		return userparam.LoginResponse{}, err
	}

	return userparam.LoginResponse{User: userparam.UserInfo{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
	}, Token: tokenStr}, nil

}
