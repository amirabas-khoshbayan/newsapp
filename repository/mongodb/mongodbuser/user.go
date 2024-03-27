package mongodbuser

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"newsapp/entity"
)

func (d DB) GetUserByID(id string) (entity.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.User{}, err
	}

	var user entity.User
	err = d.UserCollection().FindOne(context.Background(), bson.D{{"_id", objID}}).Decode(&user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
func (d DB) GetUsers() ([]entity.User, error) {

	cursor, err := d.UserCollection().Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	var users []entity.User
	err = cursor.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
func (d DB) InsertUser(user entity.User) (string, error) {
	result, err := d.UserCollection().InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	objectID := result.InsertedID.(primitive.ObjectID).Hex()

	return objectID, err
}
