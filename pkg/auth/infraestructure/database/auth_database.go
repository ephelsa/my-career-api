package database

import (
	"context"
	"database/sql"
	"ephelsa/my-career/pkg/auth/data"
	"ephelsa/my-career/pkg/auth/domain"
	"ephelsa/my-career/pkg/shared/infrastructure/database"
	"github.com/sirupsen/logrus"
)

type postgresAuthRepo struct {
	Connection *sql.DB
}

func NewPostgresAuthDatabase(db *sql.DB) data.AuthRepository {
	return &postgresAuthRepo{
		Connection: db,
	}
}

func (p *postgresAuthRepo) IsUserRegistered(c context.Context, email string) (res bool, err error) {
	query := `SELECT check_user_existence($1)`
	rows, err := database.NewRowsByQueryContext(p.Connection, c, query, email)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	if rows.Next() {
		if err = rows.Scan(&res); err != nil {
			logrus.Error(err)
		}
	}

	return
}

func (p *postgresAuthRepo) IsUserRegistryConfirmed(c context.Context, email string) (res bool, err error) {
	query := `SELECT check_user_registry_confirmed($1)`
	rows, err := database.NewRowsByQueryContext(p.Connection, c, query, email)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	if rows.Next() {
		if err = rows.Scan(&res); err != nil {
			logrus.Error(err)
		}
	}

	return
}

func (p *postgresAuthRepo) Register(c context.Context, r domain.Register) (res domain.RegisterSuccess, err error) {
	query := `INSERT INTO "user" (first_name, second_name, first_surname, second_surname, email, password, document_type, 
		institution_name, study_level, institution_type, registry_confirmed, department_code, municipality_code,
		country_code, document) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) 
		RETURNING email`
	row, err := database.NewRowsByQueryContext(p.Connection, c, query,
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
	)
	if err != nil {
		logrus.Error(err)
		return domain.RegisterSuccess{}, err
	}
	defer func() {
		if err = row.Close(); err != nil {
			logrus.Error(err)
		}
	}()

	if row.Next() {
		if err = row.Scan(&res.Email); err != nil {
			logrus.Error(err)
		}
	}

	return
}

// TODO: Implement
func (p *postgresAuthRepo) IsAuthSuccess(c context.Context, auth domain.AuthCredentials) (res bool, err error) {
	query := `SELECT authenticate_user($1, $2)`
	row, err := database.NewRowsByQueryContext(p.Connection, c, query, auth.Email, auth.Password)
	if err != nil {
		logrus.Error(err)
		return res, err
	}
	defer func() {
		if err = row.Close(); err != nil {
			logrus.Error(err)
		}
	}()

	if row.Next() {
		if err = row.Scan(&res); err != nil {
			logrus.Error(err)
			return res, err
		}
	}

	return
}
