package database

import (
	"database/sql"
	"ephelsa/my-career/internal/env"
	"log"
	"sync"
)

type Data struct {
	Database *sql.DB
}

func New(db env.Database) *Data {
	var (
		once sync.Once
		data *Data
	)

	once.Do(func() {
		db, err := postgresConnection(db)
		if err != nil {
			log.Panic(err)
		}
		if err := db.Ping(); err != nil {
			log.Panic(err)
		}

		data = &Data{
			Database: db,
		}
	})

	return data
}

func (d *Data) Close() {
	if err := d.Database.Close(); err != nil {
		log.Panic(err)
	}
}
