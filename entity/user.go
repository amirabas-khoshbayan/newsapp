package entity

import "time"

type User struct {
	ID           string    `bson:"_id,omitempty" json:"id"`
	FirstName    string    `bson:"first_name" json:"first_name"`
	LastName     string    `bson:"last_name" json:"last_name"`
	PhoneNumber  string    `bson:"phone_number" json:"phone_number"`
	Email        string    `bson:"email" json:"email"`
	Username     string    `bson:"username" json:"username"`
	Password     string    `bson:"password" json:"password"`
	RegisterDate time.Time `bson:"register_date" json:"register_date"`
	Role         Role      `bson:"role" json:"role"`
}
