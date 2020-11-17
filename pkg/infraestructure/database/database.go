package database

import (
	"database/sql"
	"ephelsa/my-career/internal/env"
	"log"
	"sync"
)

type Information struct {
	Instance *sql.DB
}

func New(db env.Database) *Information {
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
			Instance: db,
		}
	})

	return data
}

func (d *Information) Close() {
	if err := d.Instance.Close(); err != nil {
		log.Panic(err)
	}
}
