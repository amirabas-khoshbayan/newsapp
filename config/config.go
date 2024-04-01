package config

import (
	"newsapp/repository/mongodb"
	"newsapp/service/authenticationservice"
)

type HttpServer struct {
	Port               int  `yaml:"port"`
	UseCustomValidator bool `yaml:"use_custom_validator"`
}

type Config struct {
	HttpServer HttpServer                   `yaml:"http_server"`
	MongoDB    mongodb.Config               `yaml:"mongodb"`
	Auth       authenticationservice.Config `yaml:"auth"`
}

var AppConfig Config
