package database

import (
	"ephelsa/my-career/internal/env"
	"github.com/magiconair/properties/assert"
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
