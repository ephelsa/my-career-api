package database

import (
	"context"
	"database/sql"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"ephelsa/my-career/pkg/studylevel/data"
	"ephelsa/my-career/pkg/studylevel/domain"
	"github.com/sirupsen/logrus"
)

type postgresStudyLevelRepo struct {
	Connection *sql.DB
}

func NewPostgresStudyLevelRepository(db *sql.DB) data.StudyLevelRepository {
	return &postgresStudyLevelRepo{
		Connection: db,
	}
}

func (p *postgresStudyLevelRepo) fetch(c context.Context, query string, args ...interface{}) (result []domain.StudyLevel, err error) {
	rows, err := p.Connection.QueryContext(c, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result = make([]domain.StudyLevel, 0)
	for rows.Next() {
		r := domain.StudyLevel{}
		if err = rows.Scan(&r.Id, &r.Name); err != nil {
			logrus.Error(err)
			return result, err
		}
		result = append(result, r)
	}

	return result, nil
}

func (p *postgresStudyLevelRepo) FetchAll(c context.Context) (result []domain.StudyLevel, err error) {
	query := `SELECT id, value FROM study_level`
	result, err = p.fetch(c, query)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return result, sharedDomain.ResourcesEmpty
	}

	return
}
