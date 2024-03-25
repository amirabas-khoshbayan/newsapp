package mongodbuser

import (
	"newsapp/repository/mongodb"
)

type DB struct {
	conn *mongodb.MongoDB
}

func New(conn *mongodb.MongoDB) *DB {
	return &DB{conn: conn}
}
