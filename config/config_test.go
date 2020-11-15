package config

import (
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
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

	newConfiguration(db).helperAssertConfig(t, db)
}

func Test_EnvironmentConfig_Prod(t *testing.T) {
	// Prod
	os.Setenv("ENV", "prod")

	// Env keys
	db := map[string]interface{}{
		"user": "USER",
		"pass": "TOP-SECRET-PASSWORD",
		"name": "SOME-NAME",
		"host": "far-far-away.com",
		"port": "80280",
	}
	os.Setenv("DB_USER", "USER")
	os.Setenv("DB_PASS", "TOP-SECRET-PASSWORD")
	os.Setenv("DB_NAME", "SOME-NAME")
	os.Setenv("DB_HOST", "far-far-away.com")
	os.Setenv("DB_PORT", "80280")

	EnvironmentConfig().helperAssertConfig(t, db)
}

//func Test_EnvironmentConfig_Dev(t *testing.T) {
//	file, err := os.Open("config.dev.json")
//	if err != nil {
//		t.Error("Can't open file")
//	}
//
//	file.
//}
