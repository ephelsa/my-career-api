package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
)

func (conf Configuration) helperAssertConfig(t *testing.T, database map[string]interface{}, server map[string]interface{}) {
	assert.Equal(t, conf.Database.Username, database["user"])
	assert.Equal(t, conf.Database.Password, database["pass"])
	assert.Equal(t, conf.Database.Name, database["name"])
	assert.Equal(t, conf.Database.Host, database["host"])
	assert.Equal(t, conf.Database.Port, database["port"])
	assert.Equal(t, conf.Server.Port, server["port"])
}

func Test_newConfiguration(t *testing.T) {
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
	conf, err := newConfiguration(db, server)
	if err != nil {
		t.Errorf("An error has been occur %s", err)
	}

	conf.helperAssertConfig(t, db, server)
}

func openConfigFile(fileName string) (*Configuration, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Can't open file: %s \n", err)
	}

	config := Configuration{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("Parse json error: %s \n", err)
	}

	return &config, nil
}

func Test_SetupEnvironment_Dev(t *testing.T) {
	configuration, err := openConfigFile("config.dev.json")
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
	SetupEnvironment().helperAssertConfig(t, db, server)
}

func Test_SetupEnvironment_Prod(t *testing.T) {
	// Prod
	_ = os.Setenv("ENV", "prod")

	configuration, err := openConfigFile("config.prod.json")
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
		"pass": os.Getenv("DB_PASS"),
		"name": os.Getenv("DB_NAME"),
		"host": os.Getenv("DB_HOST"),
		"port": os.Getenv("DB_PORT"),
	}
	server := map[string]interface{}{
		"port": configuration.Server.Port,
	}

	SetupEnvironment().helperAssertConfig(t, db, server)
}
