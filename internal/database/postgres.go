package database

import (
	"database/sql"
	"ephelsa/my-career/internal/env"
	"fmt"

	_ "github.com/lib/pq"
)

func postgresConnection(db env.Database) (*sql.DB, error) {
	uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.Username, db.Password, db.Host, db.Port, db.Name)
	return sql.Open("postgres", uri)
}
