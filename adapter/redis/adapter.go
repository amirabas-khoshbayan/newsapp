package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Adapter struct {
	redisClient *redis.Client
}

func New(cfg Config) Adapter {
	return Adapter{redisClient: redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})}
}

func (a Adapter) Client() *redis.Client {
	return a.redisClient
}
