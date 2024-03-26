package userservice

import "newsapp/entity"

type Repository interface {
	GetUsers() ([]entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
func (s Service) GetUsers() ([]entity.User, error) {
	userList := []entity.User{{
		FirstName:   "Abbas",
		LastName:    "khosh",
		Age:         34,
		PhoneNumber: "09158346511",
	},
		{
			FirstName:   "sss",
			LastName:    "gggg",
			Age:         34,
			PhoneNumber: "09158346561",
		}}

	return userList, nil
}
