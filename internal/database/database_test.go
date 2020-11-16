package database

import (
	"ephelsa/my-career/internal/env"
	"github.com/joho/godotenv"
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
)

func Test_NewDatabase(t *testing.T) {
	envDb := env.Database{
		Username: "test-user",
		Password: "TopSecretPassword",
		Name:     "gopher",
		Host:     "127.0.0.1",
		Port:     "5432",
	}
	got := NewDatabase(envDb)
	assert.Equal(t, got.URI, "postgres://test-user:TopSecretPassword@127.0.0.1:5432/gopher")
}

func Test__Database(t *testing.T) {
	err := godotenv.Load("../env/dev.env")
	if err != nil {
		t.Errorf("Error opening dev.env file %s \n", err)
	}

	newDb := NewDatabase(env.Database{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	})

	newDb.Connect()
}
