package userparam

type CreateNewUserRequest struct {
	FirstName   string `json:"first_name" `
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email"`
	Password    string `json:"password" validate:"required"`
	Role        string `json:"role"`
}

type UserInfo struct {
	ID             string `json:"id"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	AvatarFileName string `json:"avatar_file_name"`
}

type CreateNewUserResponse struct {
	UserInfo UserInfo `json:"user_info"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User  UserInfo `json:"user"`
	Token string
}

type EditUserResponse struct {
	User  UserInfo `json:"user"`
	Token string
}

type EditUserRequest struct {
	ID             int    `json:"id"`
	FirstName      string `json:"first_name" `
	LastName       string `json:"last_name"`
	PhoneNumber    string `json:"phone_number" validate:"required"`
	Email          string `json:"email"`
	Password       string `json:"password" validate:"required"`
	Role           string `json:"role"`
	AvatarFileName string `json:"avatar_file_name"`
}
