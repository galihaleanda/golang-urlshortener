package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUrl string
	Port  string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot Lead Config")
	}

	return &Config{
		DBUrl: viper.GetString("DATABASE_URL"),
		Port:  viper.GetString("PORT"),
	}
}
