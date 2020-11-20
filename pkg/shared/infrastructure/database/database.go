package database

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
)

// NewRowsByQueryContext provide sql.Rows and handle errors
func NewRowsByQueryContext(db *sql.DB, c context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	rows, err = db.QueryContext(c, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return rows, err
}
