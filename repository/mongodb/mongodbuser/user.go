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
func (d DB) InsertUser(user entity.User) (entity.User, error) {
	result, err := d.UserCollection().InsertOne(context.Background(), &user)
	if err != nil {
		return entity.User{}, err
	}

	objectID := result.InsertedID.(primitive.ObjectID).Hex()
	user.ID = objectID

	return user, err
}
func (d DB) UpdateUserByModel(user entity.User) error {
	objID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}
	// user id skip, because operation UpdateOne error : "PerForming an update on the path `_id` would modify the immutable field `_id` "
	user.ID = ""

	_, err = d.UserCollection().UpdateOne(context.Background(), bson.D{{"_id", objID}}, bson.D{{"$set", user}})
	if err != nil {
		return err
	}
	return nil

}
func (d DB) DeleteUserByID(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = d.UserCollection().DeleteOne(context.Background(), bson.D{{"_id", objID}})
	if err != nil {
		return err
	}

	return nil
}
func (d DB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	var user entity.User
	err := d.UserCollection().FindOne(context.Background(), bson.D{{"phone_number", phoneNumber}}).Decode(&user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
