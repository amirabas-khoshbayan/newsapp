package config

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
}

var AppConfig Config
