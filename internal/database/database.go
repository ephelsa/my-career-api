package database

import (
	"context"
	"ephelsa/my-career/internal/env"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	URI  string
	Pool *pgxpool.Pool
}

func NewDatabase(db env.Database) *Database {
	return &Database{
		URI: fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db.Username, db.Password, db.Host, db.Port, db.Name),
	}
}

func (d *Database) Connect() {
	pool, err := pgxpool.Connect(context.Background(), d.URI)
	d.Pool = pool
	if err != nil {
		panic(fmt.Errorf("Error connecting with database %s \n", err))
	}
	defer pool.Close()
}

func (d *Database) CloseConnection() {
	d.Pool.Close()
}
