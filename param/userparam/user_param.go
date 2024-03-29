package userparam

type CreateNewUserRequest struct {
	FirstName   string `json:"first_name" `
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email"`
	Password    string `json:"password" validate:"required"`
}

type UserInfo struct {
	ID          string
	PhoneNumber string
	Email       string
	FirstName   string
	LastName    string
}

type CreateNewUserResponse struct {
	UserInfo UserInfo
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User  UserInfo `json:"user"`
	Token string
}
