package userservice

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"newsapp/entity"
	"newsapp/param/userparam"
	"strconv"
	"time"
)

type Repository interface {
	InsertUser(user entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, userID string) (entity.User, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (entity.User, error)
	GetUsers(ctx context.Context) ([]entity.User, error)
	UpdateUserByModel(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id string) error
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

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (s Service) GetUsers(ctx context.Context) ([]entity.User, error) {
	userList, err := s.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return userList, nil
}
func (s Service) CreateNewUser(req userparam.CreateNewUserRequest) (userparam.CreateNewUserResponse, error) {

	user := entity.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
		Password:     getMD5Hash(req.Password),
		RegisterDate: time.Now(),
		Role:         entity.Role(req.Role),
	}

	userRes, err := s.repo.InsertUser(user)
	if err != nil {
		return userparam.CreateNewUserResponse{}, err
	}

	return userparam.CreateNewUserResponse{userparam.UserInfo{
		ID:          strconv.Itoa(int(userRes.ID)),
		PhoneNumber: userRes.PhoneNumber,
		FirstName:   userRes.FirstName,
		LastName:    userRes.LastName,
		Email:       userRes.Email,
	}}, nil
}
func (s Service) Login(ctx context.Context, req userparam.LoginRequest) (userparam.LoginResponse, error) {
	user, err := s.repo.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		return userparam.LoginResponse{}, err
	}

	if user.Password != getMD5Hash(req.Password) {
		return userparam.LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}

	tokenStr, err := s.auth.CreateToken(user)
	if err != nil {
		return userparam.LoginResponse{}, err
	}

	return userparam.LoginResponse{User: userparam.UserInfo{
		ID:          strconv.Itoa(int(user.ID)),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
	}, Token: tokenStr}, nil

}
func (s Service) GiveAdminRole(ctx context.Context, userID string) error {

	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	user.Role = entity.AdminRole

	_, err = s.repo.UpdateUserByModel(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
func (s Service) EditUser(ctx context.Context, req userparam.EditUserRequest) (userparam.EditUserResponse, error) {
	userRes, err := s.repo.GetUserByID(ctx, strconv.Itoa(int(req.ID)))
	if err != nil {
		return userparam.EditUserResponse{}, err
	}
	if req.FirstName != "" {
		userRes.FirstName = req.FirstName
	}
	if req.LastName != "" {
		userRes.LastName = req.LastName
	}
	if userRes.Email != "" {
		userRes.Email = req.Email
	}
	if req.Password != "" {
		userRes.Password = req.Password
	}
	if req.AvatarFileName != "" {
		userRes.AvatarFileName = req.AvatarFileName
	}

	user, err := s.repo.UpdateUserByModel(ctx, userRes)
	if err != nil {
		return userparam.EditUserResponse{}, err
	}

	return userparam.EditUserResponse{User: userparam.UserInfo{
		ID:             strconv.Itoa(int(user.ID)),
		PhoneNumber:    user.PhoneNumber,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		AvatarFileName: user.AvatarFileName,
	}}, nil
}
func (s Service) DeleteUser(ctx context.Context, id string) error {
	err := s.repo.DeleteUser(ctx, id)

	return err
}
