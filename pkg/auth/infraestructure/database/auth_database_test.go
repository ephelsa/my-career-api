package database

import (
	"context"
	"ephelsa/my-career/pkg/auth/domain"
	testDatabase "ephelsa/my-career/test/database"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type checkUser struct {
	description string

	mockQuery string
	mockRows  *sqlmock.Rows

	emailArg string
	passArg  string

	expectedResult bool
}

func TestPostgresAuthRepo_IsUserRegistered(t *testing.T) {
	tests := []checkUser{
		{
			description:    "user exists",
			mockQuery:      "SELECT check_user_existence(?)",
			mockRows:       sqlmock.NewRows([]string{"check_user_existence"}).AddRow(true),
			emailArg:       "xephelsax@gmail.com",
			expectedResult: true,
		},
		{
			description:    "user doesn't exists",
			mockQuery:      "SELECT check_user_existence(?)",
			mockRows:       sqlmock.NewRows([]string{"check_user_existence"}).AddRow(false),
			emailArg:       "xephelsax@gmail.com",
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			db, mock := testDatabase.NewMockDatabase(t)
			repo := postgresAuthRepo{
				Connection: db,
			}
			defer func() {
				_ = repo.Connection.Close()
			}()

			mock.ExpectQuery(tt.mockQuery).WithArgs(tt.emailArg).WillReturnRows(tt.mockRows)

			result, err := repo.IsUserRegistered(context.Background(), tt.emailArg)
			assert.NoErrorf(t, err, tt.description)
			assert.Equalf(t, tt.expectedResult, result, tt.description)
		})
	}
}

func TestPostgresAuthRepo_IsUserRegistryConfirmed(t *testing.T) {
	tests := []checkUser{
		{
			description:    "confirmed",
			mockQuery:      "SELECT check_user_registry_confirmed(?)",
			mockRows:       sqlmock.NewRows([]string{"check_user_registry_confirmed"}).AddRow(true),
			emailArg:       "xephelsax@gmail.com",
			expectedResult: true,
		},
		{
			description:    "pending for confirm",
			mockQuery:      "SELECT check_user_registry_confirmed(?)",
			mockRows:       sqlmock.NewRows([]string{"check_user_registry_confirmed"}).AddRow(false),
			emailArg:       "xephelsax@gmail.com",
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			db, mock := testDatabase.NewMockDatabase(t)
			repo := postgresAuthRepo{
				Connection: db,
			}
			defer func() {
				_ = repo.Connection.Close()
			}()

			mock.ExpectQuery(tt.mockQuery).WithArgs(tt.emailArg).WillReturnRows(tt.mockRows)

			result, err := repo.IsUserRegistryConfirmed(context.Background(), tt.emailArg)
			assert.NoErrorf(t, err, tt.description)
			assert.Equalf(t, tt.expectedResult, result, tt.description)
		})
	}
}

func TestPostgresAuthRepo_Register(t *testing.T) {
	db, mock := testDatabase.NewMockDatabase(t)
	repo := postgresAuthRepo{
		Connection: db,
	}
	defer func() {
		_ = repo.Connection.Close()
	}()

	query := "INSERT INTO \"user\" \\(first_name, second_name, first_surname, second_surname, email, password, document_type, " +
		"institution_name, study_level, institution_type, registry_confirmed, department_code, municipality_code, " +
		"country_code, document\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6, \\$7, \\$8, \\$9, \\$10, \\$11, \\$12, \\$13, \\$14, \\$15\\) " +
		"RETURNING email "
	r := domain.Register{
		Email:             "xephelsax@gmail.com",
		DocumentType:      "1",
		Document:          "123123",
		FirstName:         "Leonardo",
		SecondName:        "Andres",
		FirstSurname:      "Perez",
		SecondSurname:     "Castilla",
		Password:          "SuperSecretPass",
		InstitutionType:   4,
		InstitutionName:   "UdeA",
		StudyLevel:        4,
		RegistryConfirmed: false,
		CountryCode:       "CO",
		DepartmentCode:    "70",
		MunicipalityCode:  "001",
	}
	expectedResult := domain.RegisterSuccess{
		Email: r.Email,
	}
	mr := sqlmock.NewRows([]string{"email"}).AddRow(r.Email)

	mock.ExpectQuery(query).WithArgs(
		r.FirstName,
		r.SecondName,
		r.FirstSurname,
		r.SecondSurname,
		r.Email,
		r.Password,
		r.DocumentType,
		r.InstitutionName,
		r.StudyLevel,
		r.InstitutionType,
		r.RegistryConfirmed,
		r.DepartmentCode,
		r.MunicipalityCode,
		r.CountryCode,
		r.Document,
	).WillReturnRows(mr)

	result, err := repo.Register(context.Background(), r)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestPostgresAuthRepo_IsAuthSuccess(t *testing.T) {
	tests := []checkUser{
		{
			description:    "success",
			mockQuery:      "SELECT authenticate_user\\(\\$1, \\$2\\)",
			mockRows:       sqlmock.NewRows([]string{"authenticate_user"}).AddRow(true),
			emailArg:       "xephelsax@gmail.com",
			passArg:        "123123",
			expectedResult: true,
		},
		{
			description:    "fail",
			mockQuery:      "SELECT authenticate_user\\(\\$1, \\$2\\)",
			mockRows:       sqlmock.NewRows([]string{"authenticate_user"}).AddRow(false),
			emailArg:       "xephelsax@gmail.com",
			passArg:        "123123",
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			db, mock := testDatabase.NewMockDatabase(t)
			repo := postgresAuthRepo{
				Connection: db,
			}
			defer func() {
				_ = repo.Connection.Close()
			}()

			mock.ExpectQuery(tt.mockQuery).WithArgs(tt.emailArg, tt.passArg).WillReturnRows(tt.mockRows)

			result, err := repo.IsAuthSuccess(context.Background(), domain.AuthCredentials{Email: tt.emailArg, Password: tt.passArg})
			assert.NoErrorf(t, err, tt.description)
			assert.Equalf(t, tt.expectedResult, result, tt.description)
		})
	}
}
