package database

import (
	"context"
	"database/sql"
	"ephelsa/my-career/pkg/institutiontype/data"
	"ephelsa/my-career/pkg/institutiontype/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	sharedDatabase "ephelsa/my-career/pkg/shared/infrastructure/database"
	"github.com/sirupsen/logrus"
)

type postgresInstitutionTypeRepo struct {
	Connection *sql.DB
}

func NewPostgresInstitutionTypeRepository(db *sql.DB) data.InstitutionTypeRepository {
	return &postgresInstitutionTypeRepo{
		Connection: db,
	}
}

func (p *postgresInstitutionTypeRepo) fetch(c context.Context, query string, args ...interface{}) (result []domain.InstitutionType, err error) {
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

	result = make([]domain.InstitutionType, 0)
	for rows.Next() {
		r := domain.InstitutionType{}
		if err := rows.Scan(&r.Id, &r.Name); err != nil {
			logrus.Error(err)
			return result, err
		}
		result = append(result, r)
	}

	return result, nil
}

func (p *postgresInstitutionTypeRepo) FetchAll(c context.Context) (result []domain.InstitutionType, err error) {
	query := `SELECT id, value FROM institution_type`
	result, err = p.fetch(c, query)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return result, sharedDomain.ResourcesEmpty
	}

	return
}
