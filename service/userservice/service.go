package userservice

import (
	"fmt"
	"newsapp/entity"
)

type Repository interface {
	InsertUser(user entity.User) (string, error)
	GetUserByID(id string) (entity.User, error)
	GetUsers() ([]entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
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
		Age:         24,
		PhoneNumber: "123456",
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(insertUser)

	return userList, nil
}
