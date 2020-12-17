package database

import (
	"context"
	"database/sql"
	authDomain "ephelsa/my-career/pkg/auth/domain"
	"ephelsa/my-career/pkg/shared/infrastructure/database"
	"ephelsa/my-career/pkg/user/data"
	"ephelsa/my-career/pkg/user/domain"
	"github.com/sirupsen/logrus"
	"time"
)

type postgresUserRepo struct {
	Connection *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) data.UserLocalRepository {
	return &postgresUserRepo{
		Connection: db,
	}
}

func (p *postgresUserRepo) InformationByEmail(ctx context.Context, email string) (domain.User, error) {
	query := "SELECT first_name, second_name, first_surname, second_surname, email FROM user_information u WHERE email = $1"
	row, err := database.NewRowsByQueryContext(p.Connection, ctx, query, email)
	result := domain.User{}
	defer func() {
		err = row.Close()
		logrus.Error(err)
	}()
	if err != nil {
		logrus.Error(err)
		return result, err
	}

	if row.Next() {
		if err = row.Scan(&result.FirstName, &result.SecondName, &result.FirstSurname, &result.SecondSurname, &result.Email); err != nil {
			logrus.Error(err)
			return result, err
		}
	}

	return result, err
}

func (p *postgresUserRepo) StoreUserInformation(c context.Context, u domain.User) (authDomain.RegisterSuccess, error) {
	result := authDomain.RegisterSuccess{}
	query := `INSERT INTO user_information (first_name, second_name, first_surname, second_surname, email, document_type,
		institution_name, study_level, institution_type, department_code, municipality_code,
		country_code, document, birthdate) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING email`

	// TODO: Replace this after front implementation
	u.Birthdate = "01/01/1900"
	t, err := time.Parse("01/02/2006", u.Birthdate)
	if err != nil {
		logrus.Error(err.Error())
		return result, err
	}
	birthdate := t.Format("01-02-2006")

	row, err := database.NewRowsByQueryContext(p.Connection, c, query,
		u.FirstName,
		u.SecondName,
		u.FirstSurname,
		u.SecondSurname,
		u.Email,
		u.DocumentTypeCode,
		u.InstitutionName,
		u.StudyLevelCode,
		u.InstitutionTypeCode,
		u.DepartmentCode,
		u.MunicipalityCode,
		u.CountryCode,
		u.Document,
		birthdate,
	)
	if err != nil {
		logrus.Error(err)
		return result, err
	}
	defer func() {
		if err = row.Close(); err != nil {
			logrus.Error(err)
		}
	}()

	if row.Next() {
		if err = row.Scan(&result.Email); err != nil {
			logrus.Error(err)
		}
	}

	return result, nil
}
