package entity

import "time"

type User struct {
	ID             uint      `bson:"_id,omitempty" json:"id"`
	FirstName      string    `bson:"first_name" json:"first_name"`
	LastName       string    `bson:"last_name" json:"last_name"`
	PhoneNumber    string    `bson:"phone_number" json:"phone_number"`
	Email          string    `bson:"email" json:"email"`
	Username       string    `bson:"username" json:"username"`
	Password       string    `bson:"password" json:"-"` // Password always keep hashed password
	RegisterDate   time.Time `bson:"register_date" json:"register_date"`
	Role           Role      `bson:"role" json:"role"`
	AvatarFileName string    `bson:"avatar_name" json:"avatar_file_name"`
}
