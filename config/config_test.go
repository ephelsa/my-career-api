package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
)

func (conf Configuration) helperAssertConfig(t *testing.T, database map[string]interface{}) {
	assert.Equal(t, conf.Database.Username, database["user"])
	assert.Equal(t, conf.Database.Password, database["pass"])
	assert.Equal(t, conf.Database.Name, database["name"])
	assert.Equal(t, conf.Database.Host, database["host"])
	assert.Equal(t, conf.Database.Port, database["port"])
}

func Test_newConfiguration(t *testing.T) {
	db := map[string]interface{}{
		"user": "USER",
		"pass": "TOP-SECRET-PASSWORD",
		"name": "SOME-NAME",
		"host": "far-far-away.com",
		"port": "80280",
	}

	conf, err := newConfiguration(db)
	if err != nil {
		t.Errorf("An error has been occur %s", err)
	}

	conf.helperAssertConfig(t, db)
}

func Test_SetupEnvironment_Dev(t *testing.T) {
	file, err := ioutil.ReadFile("config.dev.json")
	if err != nil {
		t.Errorf("Can't open file: %s", err)
	}

	configuration := Configuration{}
	err = json.Unmarshal(file, &configuration)
	if err != nil {
		t.Errorf("Parse json error: %s", err)
	}

	db := map[string]interface{}{
		"user": configuration.Database.Username,
		"pass": configuration.Database.Password,
		"name": configuration.Database.Name,
		"host": configuration.Database.Host,
		"port": configuration.Database.Port,
	}

	SetupEnvironment().helperAssertConfig(t, db)
}

func Test_SetupEnvironment_Prod(t *testing.T) {
	// Prod
	_ = os.Setenv("ENV", "prod")

	// Env keys
	db := map[string]interface{}{
		"user": "USER",
		"pass": "TOP-SECRET-PASSWORD",
		"name": "SOME-NAME",
		"host": "far-far-away.com",
		"port": ":80280",
	}
	_ = os.Setenv("DB_USER", "USER")
	_ = os.Setenv("DB_PASS", "TOP-SECRET-PASSWORD")
	_ = os.Setenv("DB_NAME", "SOME-NAME")
	_ = os.Setenv("DB_HOST", "far-far-away.com")
	_ = os.Setenv("DB_PORT", ":80280")

	SetupEnvironment().helperAssertConfig(t, db)
}
