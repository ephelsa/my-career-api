package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Database struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     string
}

type Configuration struct {
	Database Database
}

func EnvironmentConfig() *Configuration {
	if os.Getenv("ENV") == "prod" {
		viper.AutomaticEnv()

		db := map[string]interface{}{
			"user": viper.Get("DB_USER"),
			"pass": viper.Get("DB_PASS"),
			"name": viper.Get("DB_NAME"),
			"host": viper.Get("DB_HOST"),
			"port": viper.Get("DB_PORT"),
		}

		return newConfiguration(db)
	} else {
		viper.SetConfigName("config.dev")
		viper.SetConfigType("json")
		viper.AddConfigPath("./config/")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
			return nil
		}

		return newConfiguration(viper.GetStringMap("database"))
	}
}

func newConfiguration(database map[string]interface{}) *Configuration {
	return &Configuration{
		Database: Database{
			Username: database["user"].(string),
			Password: database["pass"].(string),
			Name:     database["name"].(string),
			Host:     database["host"].(string),
			Port:     database["port"].(string),
		},
	}
}
