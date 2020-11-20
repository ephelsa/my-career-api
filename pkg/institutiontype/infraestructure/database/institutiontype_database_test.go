package database

import (
	"context"
	"ephelsa/my-career/pkg/institutiontype/domain"
	"ephelsa/my-career/test/database"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var its = []domain.InstitutionType{
	{
		Id:   0,
		Name: "Type A",
	},
	{
		Id:   1,
		Name: "Type B",
	},
}

func TestPostgresInstitutionTypeRepo_FetchAll(t *testing.T) {
	db, mock := database.NewMockDatabase(t)
	repo := postgresInstitutionTypeRepo{
		Connection: db,
	}
	defer func() {
		_ = db.Close()
	}()

	query := `SELECT id, value FROM institution_type`
	rows := sqlmock.NewRows([]string{"id", "value"}).AddRow(its[0].Id, its[0].Name).AddRow(its[1].Id, its[1].Name)
	mock.ExpectQuery(query).WillReturnRows(rows)

	result, err := repo.FetchAll(context.Background())
	assert.NotEmpty(t, result)
	assert.NoError(t, err)
}

func TestPostgresInstitutionTypeRepo_FetchAll_Error(t *testing.T) {
	db, mock := database.NewMockDatabase(t)
	repo := postgresInstitutionTypeRepo{
		Connection: db,
	}
	defer func() {
		_ = db.Close()
	}()

	query := `SELECT id, value FROM institution_type`
	rows := sqlmock.NewRows([]string{"id", "value"})
	mock.ExpectQuery(query).WillReturnRows(rows)

	result, err := repo.FetchAll(context.Background())
	assert.Empty(t, result)
	assert.Error(t, err)
}
