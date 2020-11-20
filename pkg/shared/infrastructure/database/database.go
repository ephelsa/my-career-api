package database

import (
	"database/sql"
	"ephelsa/my-career/internal/env"
	"log"
	"sync"
)

type Information struct {
	Postgres *sql.DB
}

// NewPostgresDatabase create a new instance of Information setting Postgres
func NewPostgresDatabase(db env.Database) *Information {
	var (
		once sync.Once
		data *Information
	)

	once.Do(func() {
		db, err := postgresConnection(db)
		if err != nil {
			log.Panic(err)
		}
		if err := db.Ping(); err != nil {
			log.Panic(err)
		}

		data = &Information{
			Postgres: db,
		}
	})

	return data
}

// ClosePostgres ends the database connection
func (d *Information) ClosePostgres() {
	if err := d.Postgres.Close(); err != nil {
		log.Panic(err)
	}
}
