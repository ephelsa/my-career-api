package database

import (
	"context"
	"database/sql"
	"ephelsa/my-career/pkg/location/data"
	"ephelsa/my-career/pkg/location/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	sharedDatabase "ephelsa/my-career/pkg/shared/infrastructure/database"
	"github.com/sirupsen/logrus"
)

type postgresLocationRepo struct {
	Connection *sql.DB
}

func NewPostgresLocationRepository(db *sql.DB) data.LocationRepository {
	return &postgresLocationRepo{
		Connection: db,
	}
}

func (p *postgresLocationRepo) fetchCountries(c context.Context, query string, args ...interface{}) (result []domain.Country, err error) {
	rows, err := sharedDatabase.NewRowsByQueryContext(p.Connection, c, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result = make([]domain.Country, 0)
	for rows.Next() {
		r := domain.Country{}
		if err := rows.Scan(&r.ISOCode, &r.Name); err != nil {
			logrus.Error(err)
			return result, err
		}
		result = append(result, r)
	}

	return result, nil
}

func (p *postgresLocationRepo) FetchAllCountries(c context.Context) (result []domain.Country, err error) {
	query := `SELECT iso_code, name FROM country`
	result, err = p.fetchCountries(c, query)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		err = sharedDomain.ResourcesEmpty
	}

	return
}

func (p *postgresLocationRepo) fetchDepartments(c context.Context, query string, args ...interface{}) (result []domain.Department, err error) {
	rows, err := sharedDatabase.NewRowsByQueryContext(p.Connection, c, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result = make([]domain.Department, 0)
	for rows.Next() {
		r := domain.Department{}
		if err := rows.Scan(&r.CountryCode, &r.DepartmentCode, &r.Name); err != nil {
			logrus.Error(err)
			return result, err
		}
		result = append(result, r)
	}

	return result, nil
}

func (p *postgresLocationRepo) FetchDepartmentsByCountry(c context.Context, countryCode string) (result []domain.Department, err error) {
	query := `SELECT country_code, code, name FROM department WHERE country_code = $1`
	result, err = p.fetchDepartments(c, query, countryCode)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		err = sharedDomain.ResourcesEmpty
	}

	return
}

func (p *postgresLocationRepo) fetchMunicipalities(c context.Context, query string, args ...interface{}) (result []domain.Municipality, err error) {
	rows, err := sharedDatabase.NewRowsByQueryContext(p.Connection, c, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result = make([]domain.Municipality, 0)
	for rows.Next() {
		r := domain.Municipality{}
		if err := rows.Scan(&r.CountryCode, &r.DepartmentCode, &r.MunicipalityCode, &r.Name); err != nil {
			logrus.Error(err)
			return result, err
		}
		result = append(result, r)
	}

	return result, nil
}

func (p *postgresLocationRepo) FetchMunicipalitiesByDepartmentAndCountry(c context.Context, countryCode string, departmentCode string) (result []domain.Municipality, err error) {
	query := `SELECT country_code, department_code, code, name FROM municipality WHERE country_code = $1 AND department_code = $2`
	result, err = p.fetchMunicipalities(c, query, countryCode, departmentCode)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		err = sharedDomain.ResourcesEmpty
	}

	return
}
