package server

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_NewServer(t *testing.T) {
	port := ":6969"
	got := NewServer(port)

	if got.App == nil {
		t.Error("App is nil")
	}
	assert.Equal(t, got.Port, port)
}
