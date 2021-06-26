package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Keys     KeysConfiguration
}

func Init(env string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	var configuration *Configuration

	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("error on parsingn configuration file, %v", err)
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode config into struct, %v", err)
	}

	Config = configuration
}

func GetConfig() *Configuration {
	return Config
}
