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
