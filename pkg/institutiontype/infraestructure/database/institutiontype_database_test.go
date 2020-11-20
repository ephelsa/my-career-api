package database

import (
	"context"
	"ephelsa/my-career/pkg/institutiontype/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
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
	tests := []struct {
		description string

		mockQuery string
		mockRows  *sqlmock.Rows

		expectedError      bool
		expectedErrorType  error
		expectedLenResults int
	}{
		{
			description:        "fetch all",
			mockQuery:          "SELECT id, value FROM institution_type",
			mockRows:           sqlmock.NewRows([]string{"id", "value"}).AddRow(its[0].Id, its[0].Name).AddRow(its[1].Id, its[1].Name),
			expectedError:      false,
			expectedLenResults: 2,
		},
		{
			description:        "fetch all with empty resource error",
			mockQuery:          "SELECT id, value FROM institution_type",
			mockRows:           sqlmock.NewRows([]string{"id", "value"}),
			expectedError:      true,
			expectedErrorType:  sharedDomain.ResourcesEmpty,
			expectedLenResults: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			db, mock := database.NewMockDatabase(t)
			repo := postgresInstitutionTypeRepo{
				Connection: db,
			}
			defer func() {
				_ = db.Close()
			}()

			mock.ExpectQuery(test.mockQuery).WillReturnRows(test.mockRows)

			result, err := repo.FetchAll(context.Background())
			assert.Equalf(t, test.expectedLenResults, len(result), test.description)
			assert.Equalf(t, test.expectedError, err != nil, test.description)
			if test.expectedError {
				assert.Equalf(t, test.expectedErrorType, err, test.description)
			}
		})
	}
}
