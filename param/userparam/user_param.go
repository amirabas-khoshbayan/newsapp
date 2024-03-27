package userparam

type CreateNewUserRequest struct {
	FirstName   string `json:"first_name" `
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password" validate:"required"`
}
type UserInfo struct {
	ID          string
	UserName    string
	PhoneNumber string
	Email       string
	FirstName   string
	LastName    string
}
type CreateNewUserResponse struct {
	UserInfo UserInfo
}
