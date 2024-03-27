package config

import "newsapp/repository/mongodb"

type HttpServer struct {
	Port               int  `yaml:"port"`
	UseCustomValidator bool `yaml:"use_custom_validator"`
}

type Config struct {
	HttpServer HttpServer     `yaml:"http_server"`
	MongoDB    mongodb.Config `yaml:"mongodb"`
}

var AppConfig Config
