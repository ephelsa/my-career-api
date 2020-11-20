package database

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func NewMockDatabase(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("error opening stub database %s", err)
	}

	return db, mock
}
