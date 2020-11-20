package database

import (
	"context"
	"database/sql"
	"ephelsa/my-career/pkg/documenttype/data"
	"ephelsa/my-career/pkg/documenttype/domain"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	"github.com/sirupsen/logrus"
)

type postgresDocumentTypeRepo struct {
	Connection *sql.DB
}

func NewPostgresDocumentTypeRepository(db *sql.DB) data.DocumentTypeRepository {
	return &postgresDocumentTypeRepo{
		Connection: db,
	}
}

func (p *postgresDocumentTypeRepo) fetch(c context.Context, query string, args ...interface{}) (result []domain.DocumentType, err error) {
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

	result = make([]domain.DocumentType, 0)
	for rows.Next() {
		r := domain.DocumentType{}
		if err = rows.Scan(&r.Id, &r.Name); err != nil {
			logrus.Error(err)
			return result, err
		}
		result = append(result, r)
	}

	return result, nil
}

func (p *postgresDocumentTypeRepo) FetchAll(c context.Context) (result []domain.DocumentType, err error) {
	query := `SELECT id, value FROM document_type`
	result, err = p.fetch(c, query)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return result, sharedDomain.ResourcesEmpty
	}

	return
}

func (p *postgresDocumentTypeRepo) FetchByID(c context.Context, id string) (result domain.DocumentType, err error) {
	query := `SELECT id, value FROM document_type WHERE id = $1`
	list, err := p.fetch(c, query, id)
	if err != nil {
		return domain.DocumentType{}, err
	}

	if len(list) > 0 {
		result = list[0]
	} else {
		return result, sharedDomain.ResourceNotFound(id)
	}

	return
}
