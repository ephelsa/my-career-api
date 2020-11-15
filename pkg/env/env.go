package env

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

type Config struct {
	Database Database `json:"database"`
	Server   Server   `json:"server"`
}

func Envs() Environments {
	return Environments{
		Production: EnvironmentFile{
			Name:      "config.prod",
			Extension: "json",
		},
		Development: EnvironmentFile{
			Name:      "config.dev",
			Extension: "json",
		},
	}
}

type EnvironmentFile struct {
	Name      string
	Extension string
}

type Environments struct {
	Production  EnvironmentFile
	Development EnvironmentFile
}

func Setup() *Config {
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
		err := viperReadConfigFile(Envs().Production)
		if err != nil {
			panic(err)
		}

		server = viper.GetStringMap("server")
	} else {
		err := viperReadConfigFile(Envs().Development)
		if err != nil {
			panic(err)
		}

		db = viper.GetStringMap("database")
		server = viper.GetStringMap("server")
	}

	conf, err := newConfig(db, server)
	if err != nil {
		panic(fmt.Errorf("Error retrieving Config: %s \n", err))
	}

	return conf
}

func viperReadConfigFile(file EnvironmentFile) error {
	viper.SetConfigName(file.Name)
	viper.SetConfigType(file.Extension)
	viper.AddConfigPath("./pkg/env") // for build
	viper.AddConfigPath(".")         // for test

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Fatal error env file: %s \n", err)
	} else {
		return nil
	}
}

func newConfig(database map[string]interface{}, server map[string]interface{}) (*Config, error) {
	unparsed := map[string]interface{}{
		"database": database,
		"server":   server,
	}
	config := Config{}
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
