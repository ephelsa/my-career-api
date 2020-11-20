package database

import (
	"context"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
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
	tests := []struct {
		description string

		mockQuery string
		mockRows  *sqlmock.Rows

		expectedError     bool
		expectedErrorType error
		expectedLenResult int
	}{
		{
			description:       "fetch all",
			mockQuery:         "SELECT id, value FROM study_level",
			mockRows:          sqlmock.NewRows([]string{"id", "value"}).AddRow(sts[0].Id, sts[0].Name).AddRow(sts[1].Id, sts[1].Name),
			expectedError:     false,
			expectedLenResult: 2,
		},
		{
			description:       "fetch all with empty resource error",
			mockQuery:         "SELECT id, value FROM study_level",
			mockRows:          sqlmock.NewRows([]string{"id", "value"}),
			expectedError:     true,
			expectedErrorType: sharedDomain.ResourcesEmpty,
			expectedLenResult: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			db, mock := database.NewMockDatabase(t)
			repo := postgresStudyLevelRepo{
				Connection: db,
			}
			defer func() {
				_ = db.Close()
			}()

			mock.ExpectQuery(test.mockQuery).WillReturnRows(test.mockRows)

			result, err := repo.FetchAll(context.Background())
			assert.Equalf(t, test.expectedLenResult, len(result), test.description)
			assert.Equalf(t, test.expectedError, err != nil, test.description)
			if test.expectedError {
				assert.Equalf(t, test.expectedErrorType, err, test.description)
			}
		})
	}
}
