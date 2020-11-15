package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Database struct {
	Username string `json:"user"`
	Password string `json:"pass"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type Configuration struct {
	Database Database `json:"database"`
}

func SetupEnvironment() *Configuration {
	var db = map[string]interface{}{}

	if os.Getenv("ENV") == "prod" {
		viper.AutomaticEnv()

		db = map[string]interface{}{
			"user": viper.Get("DB_USER"),
			"pass": viper.Get("DB_PASS"),
			"name": viper.Get("DB_NAME"),
			"host": viper.Get("DB_HOST"),
			"port": viper.Get("DB_PORT"),
		}
	} else {
		viper.SetConfigName("config.dev")
		viper.SetConfigType("json")
		viper.AddConfigPath("./pkg/config") // for build
		viper.AddConfigPath(".")            // for test
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
			return nil
		}

		db = viper.GetStringMap("database")
	}

	conf, err := newConfiguration(db)
	if err != nil {
		panic(fmt.Errorf("Error retrieving Configuration: %s \n", err))
	}

	return conf
}

func newConfiguration(database map[string]interface{}) (*Configuration, error) {
	unparsed := map[string]interface{}{
		"database": database,
	}
	config := Configuration{}
	bytes, err := json.Marshal(unparsed)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
