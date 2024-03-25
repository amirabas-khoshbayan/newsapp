package userservice

import "newsapp/entity"

type Repository interface {
	RegisterUser(u entity.User) (entity.User, error)
}
type Service struct {
	repo Repository
}
