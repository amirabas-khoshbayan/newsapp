package config

import (
	"newsapp/adapter/redis"
	"newsapp/logger"
	"newsapp/repository/mongodb"
	"newsapp/repository/mysql"
	"newsapp/service/authenticationservice"
	"newsapp/service/publishservice"
)

type HttpServer struct {
	Port               int  `yaml:"port"`
	UseCustomValidator bool `yaml:"use_custom_validator"`
}

type Config struct {
	HttpServer     HttpServer                   `yaml:"http_server"`
	MongoDB        mongodb.Config               `yaml:"mongodb"`
	MySQL          mysql.Config                 `yaml:"mysql"`
	Auth           authenticationservice.Config `yaml:"auth"`
	ZapLogger      logger.Config                `yaml:"zap_logger"`
	Redis          redis.Config                 `yaml:"redis"`
	PublishService publishservice.Config        `yaml:"publish_service"`
}

var AppConfig Config
