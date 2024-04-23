package userservice

import (
	"context"
	"fmt"
	"newsapp/entity"
	"newsapp/param/userparam"
	"testing"
	"time"
)

//TODO - Refactor and clean Mock Repository and authentication

type mockRepo struct{}
type mockAuth struct{}

func (m *mockAuth) CreateToken(user entity.User) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockRepo) InsertUser(user entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockRepo) GetUserByID(ctx context.Context, userID int) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockRepo) GetUsers(ctx context.Context) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockRepo) UpdateUserByModel(ctx context.Context, user entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockRepo) DeleteUser(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockRepo) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (entity.User, error) {
	if phoneNumber == "" {
		return entity.User{}, fmt.Errorf("no user found")
	}

	return entity.User{
		ID:           1,
		FirstName:    "test",
		LastName:     "test",
		PhoneNumber:  phoneNumber,
		Email:        "test@gmail.com",
		Username:     "test",
		Password:     "123456",
		RegisterDate: time.Now(),
	}, nil

}

func TestLogin(t *testing.T) {
	t.Run("no user found", func(t *testing.T) {
		mRepo := new(mockRepo)
		mAuth := new(mockAuth)
		userSvc := New(mRepo, mAuth)
		_, err := userSvc.Login(context.Background(), userparam.LoginRequest{
			PhoneNumber: "",
			Password:    "123456",
		})

		if err == nil {
			t.Fatal("error is nil")
		}

	})

	t.Run("password is not correct", func(t *testing.T) {
		mRepo := new(mockRepo)
		mAuth := new(mockAuth)
		userSvc := New(mRepo, mAuth)
		_, err := userSvc.Login(context.Background(), userparam.LoginRequest{
			PhoneNumber: "0915888888",
			Password:    "123456",
		})

		if err == nil {
			t.Fatal("error is nil")
		}
	})

	t.Run("token is failed", func(t *testing.T) {
		mRepo := new(mockRepo)
		mAuth := new(mockAuth)
		userSvc := New(mRepo, mAuth)
		_, err := userSvc.Login(context.Background(), userparam.LoginRequest{
			PhoneNumber: "0915888888",
			Password:    "123456",
		})

		if err == nil {
			t.Fatal("error is nil")
		}
	})

	t.Run("successful login", func(t *testing.T) {
		mRepo := new(mockRepo)
		mAuth := new(mockAuth)
		userSvc := New(mRepo, mAuth)
		_, err := userSvc.Login(context.Background(), userparam.LoginRequest{
			PhoneNumber: "0915888888",
			Password:    "123456",
		})

		if err != nil {
			t.Fatal("error is not nil")
		}
	})
}
