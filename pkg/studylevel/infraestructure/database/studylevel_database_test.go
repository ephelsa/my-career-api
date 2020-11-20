package database

import (
	"context"
	"ephelsa/my-career/pkg/studylevel/domain"
	"ephelsa/my-career/test/database"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sts = []domain.StudyLevel{
	{
		Id:   1,
		Name: "Private",
	},
	{
		Id:   2,
		Name: "Public",
	},
}

func TestPostgresStudyLevelRepo_FetchAll(t *testing.T) {
	db, mock := database.NewMockDatabase(t)
	repo := postgresStudyLevelRepo{
		Connection: db,
	}
	defer func() {
		_ = db.Close()
	}()

	query := `SELECT id, value FROM study_level`
	rows := sqlmock.NewRows([]string{"id", "value"}).AddRow(sts[0].Id, sts[0].Name).AddRow(sts[1].Id, sts[1].Name)
	mock.ExpectQuery(query).WillReturnRows(rows)

	result, err := repo.FetchAll(context.Background())
	assert.NotEmpty(t, result)
	assert.NoError(t, err)
	assert.Len(t, result, len(sts))
}

func TestPostgresStudyLevelRepo_FetchAll_Error(t *testing.T) {
	db, mock := database.NewMockDatabase(t)
	repo := postgresStudyLevelRepo{
		Connection: db,
	}
	defer func() {
		_ = db.Close()
	}()

	query := `SELECT id, value FROM study_level`
	rows := sqlmock.NewRows([]string{"id", "value"})
	mock.ExpectQuery(query).WillReturnRows(rows)

	result, err := repo.FetchAll(context.Background())
	assert.Empty(t, result)
	assert.Error(t, err)
}
