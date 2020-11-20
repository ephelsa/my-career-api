package database

import (
	"context"
	"ephelsa/my-career/pkg/documenttype/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
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
			mockQuery:          "SELECT id, value FROM document_type",
			mockRows:           sqlmock.NewRows([]string{"id", "value"}).AddRow(dts[0].Id, dts[0].Name).AddRow(dts[1].Id, dts[1].Name),
			expectedError:      false,
			expectedLenResults: 2,
		},
		{
			description:        "fetch all with resource empty error",
			mockQuery:          "SELECT id, value FROM document_type",
			mockRows:           sqlmock.NewRows([]string{"id", "value"}),
			expectedErrorType:  sharedDomain.ResourcesEmpty,
			expectedError:      true,
			expectedLenResults: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			db, mock := database.NewMockDatabase(t)
			repo := postgresDocumentTypeRepo{
				Connection: db,
			}
			defer func() {
				_ = repo.Connection.Close()
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

func TestPostgresDocumentTypeRepo_FetchByID(t *testing.T) {
	tests := []struct {
		description string

		idArg string

		mockQuery string
		mockRows  *sqlmock.Rows

		expectedError     bool
		expectedErrorType error
		expectedEmpty     bool
	}{
		{
			description:   "find a result",
			idArg:         dts[1].Id,
			mockQuery:     "SELECT id, value FROM document_type WHERE id = ?",
			mockRows:      sqlmock.NewRows([]string{"id", "value"}).AddRow(dts[0].Id, dts[0].Name).AddRow(dts[1].Id, dts[1].Name),
			expectedError: false,
			expectedEmpty: false,
		},
		{
			description:       "resource not found",
			idArg:             dts[1].Id,
			mockQuery:         "SELECT id, value FROM document_type WHERE id = ?",
			mockRows:          sqlmock.NewRows([]string{"id", "value"}),
			expectedError:     true,
			expectedErrorType: sharedDomain.ResourceNotFound(dts[1].Id),
			expectedEmpty:     true,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			db, mock := database.NewMockDatabase(t)
			repo := postgresDocumentTypeRepo{
				Connection: db,
			}
			defer func() {
				_ = repo.Connection.Close()
			}()

			mock.ExpectQuery(test.mockQuery).WithArgs(test.idArg).WillReturnRows(test.mockRows)

			result, err := repo.FetchByID(context.Background(), test.idArg)
			assert.Equalf(t, test.expectedError, err != nil, test.description)

			if test.expectedError {
				assert.Equalf(t, test.expectedErrorType, err, test.description)
			}
			if test.expectedEmpty {
				assert.Emptyf(t, result, test.description)
			} else {
				assert.NotEmpty(t, result, test.description)
			}
		})
	}
}
