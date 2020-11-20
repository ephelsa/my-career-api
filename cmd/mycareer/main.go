package main

import (
	"ephelsa/my-career/internal/env"
	sharedData "ephelsa/my-career/pkg/shared/data"
	sharedDatabase "ephelsa/my-career/pkg/shared/infrastructure/database"
	sharedServer "ephelsa/my-career/pkg/shared/infrastructure/server"
)

func main() {
	envConfig := env.Setup()
	api := sharedServer.NewServer()
	db := sharedDatabase.NewPostgresDatabase(envConfig.Database) //nolint:staticcheck

	api.Middleware()
	sharedData.AddServerRouter(api, db.Postgres)

	api.Start(envConfig.Server.Port)
	api.Close()
	db.ClosePostgres()
}
