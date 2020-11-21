package database

import (
	"context"
	"ephelsa/my-career/pkg/location/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"ephelsa/my-career/test/database"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	cs = []domain.Country{
		{
			ISOCode: "CO",
			Name:    "Colombia",
		},
		{
			ISOCode: "MX",
			Name:    "Mexico",
		},
	}

	ds = []domain.Department{
		{
			CountryCode:    cs[0].ISOCode,
			DepartmentCode: "05",
			Name:           "Antioquia",
		},
	}

	ms = []domain.Municipality{
		{
			CountryCode:      "CO",
			DepartmentCode:   "05",
			MunicipalityCode: "01",
			Name:             "Medellin",
		},
		{
			CountryCode:      "CO",
			DepartmentCode:   "05",
			MunicipalityCode: "266",
			Name:             "Envigado",
		},
	}
)

func TestPostgresLocationRepo_FetchAllCountries(t *testing.T) {
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
			mockQuery:          "SELECT iso_code, name FROM country",
			mockRows:           sqlmock.NewRows([]string{"iso_code", "name"}).AddRow(cs[0].ISOCode, cs[0].Name).AddRow(cs[1].ISOCode, cs[1].Name),
			expectedError:      false,
			expectedLenResults: 2,
		},
		{
			description:        "fetch all with resource empty error",
			mockQuery:          "SELECT iso_code, name FROM country",
			mockRows:           sqlmock.NewRows([]string{"iso_code", "name"}),
			expectedError:      true,
			expectedErrorType:  sharedDomain.ResourcesEmpty,
			expectedLenResults: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			db, mock := database.NewMockDatabase(t)
			repo := postgresLocationRepo{
				Connection: db,
			}
			defer func() {
				_ = repo.Connection.Close()
			}()

			mock.ExpectQuery(test.mockQuery).WillReturnRows(test.mockRows)
			result, err := repo.FetchAllCountries(context.Background())

			assert.Equalf(t, test.expectedLenResults, len(result), test.description)
			assert.Equalf(t, test.expectedError, err != nil, test.description)
			if test.expectedError {
				assert.Equalf(t, test.expectedErrorType, err, test.description)
			}
		})
	}
}

func TestPostgresLocationRepo_FetchDepartmentsByCountry(t *testing.T) {
	tests := []struct {
		description string

		countryCodeArg string

		mockQuery string
		mockRows  *sqlmock.Rows

		expectedError      bool
		expectedErrorType  error
		expectedLenResults int
	}{
		{
			description:        "retrieve all departments by country",
			countryCodeArg:     cs[0].ISOCode,
			mockQuery:          "SELECT country_code, code, name FROM department WHERE country_code = \\$1",
			mockRows:           sqlmock.NewRows([]string{"country_code", "code", "name"}).AddRow(ds[0].CountryCode, ds[0].DepartmentCode, ds[0].Name),
			expectedError:      false,
			expectedErrorType:  nil,
			expectedLenResults: 1,
		},
		{
			description:        "empty resource error",
			countryCodeArg:     cs[0].ISOCode,
			mockQuery:          "SELECT country_code, code, name FROM department WHERE country_code = \\$1",
			mockRows:           sqlmock.NewRows([]string{"country_code", "code", "name"}),
			expectedError:      true,
			expectedErrorType:  sharedDomain.ResourcesEmpty,
			expectedLenResults: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			db, mock := database.NewMockDatabase(t)
			repo := postgresLocationRepo{
				Connection: db,
			}
			defer func() {
				_ = repo.Connection.Close()
			}()

			mock.ExpectQuery(tt.mockQuery).WithArgs(tt.countryCodeArg).WillReturnRows(tt.mockRows)
			result, err := repo.FetchDepartmentsByCountry(context.Background(), tt.countryCodeArg)

			assert.Equalf(t, tt.expectedLenResults, len(result), tt.description)
			assert.Equalf(t, tt.expectedError, err != nil, tt.description)
			if tt.expectedError {
				assert.Equalf(t, tt.expectedErrorType, err, tt.description)
			}
		})
	}
}

func TestPostgresLocationRepo_FetchMunicipalitiesByDepartmentAndCountry(t *testing.T) {
	type arg struct {
		CountryCode    string
		DepartmentCode string
	}

	tests := []struct {
		description string

		args arg

		mockQuery string
		mockRows  *sqlmock.Rows

		expectedError      bool
		expectedErrorType  error
		expectedLenResults int
	}{
		{
			description: "all municipalities by department and country",
			args: arg{
				CountryCode:    "CO",
				DepartmentCode: "05",
			},
			mockQuery:          "SELECT country_code, department_code, code, name FROM municipality WHERE country_code = \\$1 AND department_code = \\$2",
			mockRows:           sqlmock.NewRows([]string{"country_code", "department_code", "code", "name"}).AddRow(ms[0].CountryCode, ms[0].DepartmentCode, ms[0].MunicipalityCode, ms[0].Name).AddRow(ms[1].CountryCode, ms[1].DepartmentCode, ms[1].MunicipalityCode, ms[1].Name),
			expectedError:      false,
			expectedErrorType:  nil,
			expectedLenResults: 2,
		},
		{
			description: "resource empty",
			args: arg{
				CountryCode:    "CO",
				DepartmentCode: "05",
			},
			mockQuery:          "SELECT country_code, department_code, code, name FROM municipality WHERE country_code = \\$1 AND department_code = \\$2",
			mockRows:           sqlmock.NewRows([]string{"country_code", "department_code", "code", "name"}),
			expectedError:      true,
			expectedErrorType:  sharedDomain.ResourcesEmpty,
			expectedLenResults: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			db, mock := database.NewMockDatabase(t)
			repo := postgresLocationRepo{
				Connection: db,
			}
			defer func() {
				_ = repo.Connection.Close()
			}()

			mock.ExpectQuery(tt.mockQuery).WithArgs(tt.args.CountryCode, tt.args.DepartmentCode).WillReturnRows(tt.mockRows)
			result, err := repo.FetchMunicipalitiesByDepartmentAndCountry(context.Background(), tt.args.CountryCode, tt.args.DepartmentCode)

			assert.Equalf(t, tt.expectedLenResults, len(result), tt.description)
			assert.Equalf(t, tt.expectedError, err != nil, tt.description)
			if tt.expectedError {
				assert.Equalf(t, tt.expectedErrorType, err, tt.description)
			}
		})
	}
}
