package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
)

func (conf Config) helperAssertConfig(t *testing.T, database map[string]interface{}, server map[string]interface{}) {
	assert.Equal(t, conf.Database.Username, database["user"])
	assert.Equal(t, conf.Database.Password, database["pass"])
	assert.Equal(t, conf.Database.Name, database["name"])
	assert.Equal(t, conf.Database.Host, database["host"])
	assert.Equal(t, conf.Database.Port, database["port"])
	assert.Equal(t, conf.Server.Port, server["port"])
}

func Test_newConfig(t *testing.T) {
	db := map[string]interface{}{
		"user": "USER",
		"pass": "TOP-SECRET-PASSWORD",
		"name": "SOME-NAME",
		"host": "far-far-away.com",
		"port": "80280",
	}
	server := map[string]interface{}{
		"port": "6969",
	}
	conf, err := newConfig(db, server)
	if err != nil {
		t.Errorf("An error has been occur %s", err)
	}

	conf.helperAssertConfig(t, db, server)
}

func openConfigFile(fileName string) (*Config, error) {
	env, err := godotenv.Read(fileName)
	if err != nil {
		return nil, fmt.Errorf("Can't open %s with the error: %s \n", fileName, err)
	}

	config := Config{
		Database: Database{
			Username: env["DB_USER"],
			Password: env["DB_PASSWORD"],
			Name:     env["DB_NAME"],
			Host:     env["DB_HOST"],
			Port:     env["DB_PORT"],
		},
		Server: Server{
			Port: env["SERVER_PORT"],
		},
	}

	return &config, nil
}

func Test_Setup_Development(t *testing.T) {
	configuration, err := openConfigFile(fmt.Sprintf("%s.%s", Envs().Development.Name, Envs().Development.Extension))
	if err != nil {
		t.Error(err)
	}

	db := map[string]interface{}{
		"user": configuration.Database.Username,
		"pass": configuration.Database.Password,
		"name": configuration.Database.Name,
		"host": configuration.Database.Host,
		"port": configuration.Database.Port,
	}
	server := map[string]interface{}{
		"port": configuration.Server.Port,
	}
	Setup().helperAssertConfig(t, db, server)
}

func Test_Setup_Production(t *testing.T) {
	// Prod
	_ = os.Setenv("ENV", "prod")

	configuration, err := openConfigFile(fmt.Sprintf("%s.%s", Envs().Production.Name, Envs().Production.Extension))
	if err != nil {
		t.Error(err)
	}

	// Env keys
	_ = os.Setenv("DB_USER", "USER")
	_ = os.Setenv("DB_PASS", "TOP-SECRET-PASSWORD")
	_ = os.Setenv("DB_NAME", "SOME-NAME")
	_ = os.Setenv("DB_HOST", "far-far-away.com")
	_ = os.Setenv("DB_PORT", ":80280")

	db := map[string]interface{}{
		"user": os.Getenv("DB_USER"),
		"pass": os.Getenv("DB_PASSWORD"),
		"name": os.Getenv("DB_NAME"),
		"host": os.Getenv("DB_HOST"),
		"port": os.Getenv("DB_PORT"),
	}
	server := map[string]interface{}{
		"port": configuration.Server.Port,
	}

	Setup().helperAssertConfig(t, db, server)
}
