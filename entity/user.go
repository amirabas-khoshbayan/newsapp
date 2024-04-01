package entity

import "time"

type User struct {
	ID           string    `bson:"_id,omitempty"`
	FirstName    string    `bson:"first_name"`
	LastName     string    `bson:"last_name"`
	PhoneNumber  string    `bson:"phone_number"`
	Email        string    `bson:"email"`
	UserName     string    `bson:"user_name"`
	Password     string    `bson:"password"`
	RegisterDate time.Time `bson:"register_date"`
	Role         Role      `bson:"role"`
}
