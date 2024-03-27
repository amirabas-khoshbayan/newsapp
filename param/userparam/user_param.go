package userparam

type CreateNewUserRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
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
