package config

import (
	"github.com/dhikaroofi/simple-rest-api/pkg/viper"
	"log"
)

type Config struct {
	AppPort string
	DB      DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func LoadConfigFromFile(path string) *Config {
	conf := new(Config)
	if err := viper.LoadYAMLToStruct(path, conf); err != nil {
		log.Fatalf("failed to load yaml config: %s", err.Error())
	}

	return conf
}
