package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type MongoDB struct {
	config Config
	client *mongo.Client
}

func (m MongoDB) Connect() *mongo.Client {
	return m.client

}

func New(cfg Config) *MongoDB {
	ctx := context.TODO()

	uri := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err.Error())
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err.Error())
	}

	return &MongoDB{config: cfg, client: client}
}
