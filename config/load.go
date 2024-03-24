package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func GetConfig() Config {

	file, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		panic(err)
	}

	return AppConfig
}
