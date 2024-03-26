package mongodbuser

import (
	"go.mongodb.org/mongo-driver/mongo"
	"newsapp/config"
	"newsapp/repository/mongodb"
)

type DB struct {
	conn *mongodb.MongoDB
}

func New(conn *mongodb.MongoDB) *DB {
	return &DB{conn: conn}
}
func (d DB) UserCollection() *mongo.Collection {
	return d.conn.Connection().Database(config.AppName).Collection("user")
}
