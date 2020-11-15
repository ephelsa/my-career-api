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

type Server struct {
	Port string `json:"port"`
}

type Configuration struct {
	Database Database `json:"database"`
	Server   Server   `json:"server"`
}

func SetupEnvironment() *Configuration {
	var db = map[string]interface{}{}
	var server = map[string]interface{}{}

	if os.Getenv("ENV") == "prod" {
		viper.AutomaticEnv()
		db = map[string]interface{}{
			"user": viper.Get("DB_USER"),
			"pass": viper.Get("DB_PASS"),
			"name": viper.Get("DB_NAME"),
			"host": viper.Get("DB_HOST"),
			"port": viper.Get("DB_PORT"),
		}
		err := viperReadConfigFile("config.prod", "json")
		if err != nil {
			panic(err)
		}
		server = viper.GetStringMap("server")
	} else {
		err := viperReadConfigFile("config.dev", "json")
		if err != nil {
			panic(err)
		}

		db = viper.GetStringMap("database")
		server = viper.GetStringMap("server")
	}

	conf, err := newConfiguration(db, server)
	if err != nil {
		panic(fmt.Errorf("Error retrieving Configuration: %s \n", err))
	}

	return conf
}

func viperReadConfigFile(name string, extension string) error {
	viper.SetConfigName(name)
	viper.SetConfigType(extension)
	viper.AddConfigPath("./pkg/config") // for build
	viper.AddConfigPath(".")            // for test

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	} else {
		return nil
	}
}

func newConfiguration(database map[string]interface{}, server map[string]interface{}) (*Configuration, error) {
	unparsed := map[string]interface{}{
		"database": database,
		"server":   server,
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
