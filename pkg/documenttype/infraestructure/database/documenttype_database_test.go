package database

import (
	"context"
	"ephelsa/my-career/pkg/documenttype/domain"
	"ephelsa/my-career/test/database"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var dts = []domain.DocumentType{
	{
		Id:   "1",
		Name: "First Document Type",
	},
	{
		Id:   "2",
		Name: "Second Document Type",
	},
}

func TestPostgresDocumentTypeRepo_FetchAll(t *testing.T) {
	db, mock := database.NewMockDatabase(t)
	repo := postgresDocumentTypeRepo{
		Connection: db,
	}
	defer func() {
		_ = repo.Connection.Close()
	}()

	query := "SELECT id, value FROM document_type"
	rows := sqlmock.NewRows([]string{"id", "value"}).AddRow(dts[0].Id, dts[0].Name).AddRow(dts[1].Id, dts[1].Name)
	mock.ExpectQuery(query).WillReturnRows(rows)

	result, err := repo.FetchAll(context.Background())
	assert.NotNil(t, result)
	assert.NoError(t, err)
}

func TestPostgresDocumentTypeRepo_FetchAll_Error(t *testing.T) {
	db, mock := database.NewMockDatabase(t)
	repo := postgresDocumentTypeRepo{
		Connection: db,
	}
	defer func() {
		_ = repo.Connection.Close()
	}()

	query := "SELECT id, value FROM document_type"
	rows := sqlmock.NewRows([]string{"id", "value"})
	mock.ExpectQuery(query).WillReturnRows(rows)

	result, err := repo.FetchAll(context.Background())
	assert.Empty(t, result)
	assert.Error(t, err)
}

func TestPostgresDocumentTypeRepo_FetchByID(t *testing.T) {
	db, mock := database.NewMockDatabase(t)
	repo := postgresDocumentTypeRepo{
		Connection: db,
	}
	defer func() {
		_ = repo.Connection.Close()
	}()

	query := `SELECT id, value FROM document_type WHERE id = ?`
	rows := sqlmock.NewRows([]string{"id", "value"}).AddRow(dts[0].Id, dts[0].Name).AddRow(dts[1].Id, dts[1].Name)
	mock.ExpectQuery(query).WithArgs(dts[1].Id).WillReturnRows(rows)

	result, err := repo.FetchByID(context.Background(), dts[1].Id)

	assert.NotEmpty(t, result)
	assert.NoError(t, err)
}

func TestPostgresDocumentTypeRepo_FetchByID_Error(t *testing.T) {
	db, mock := database.NewMockDatabase(t)
	repo := postgresDocumentTypeRepo{
		Connection: db,
	}
	defer func() {
		_ = repo.Connection.Close()
	}()

	query := `SELECT id, value FROM document_type WHERE id = ?`
	rows := sqlmock.NewRows([]string{"id", "value"})
	mock.ExpectQuery(query).WithArgs(dts[1].Id).WillReturnRows(rows)

	result, err := repo.FetchByID(context.Background(), dts[1].Id)

	assert.Empty(t, result)
	assert.Error(t, err)
}
