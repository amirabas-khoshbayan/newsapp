package entity

type User struct {
	ID          string `bson:"_id,omitempty"`
	FirstName   string `bson:"first_name"`
	LastName    string `bson:"last_name"`
	Age         int    `bson:"age"`
	PhoneNumber string `bson:"phone_number"`
}
