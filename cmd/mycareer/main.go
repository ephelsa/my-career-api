package main

import (
	"ephelsa/my-career/internal/env"
	"ephelsa/my-career/pkg/infraestructure/database"
	"ephelsa/my-career/pkg/infraestructure/server"
)

func main() {
	envConfig := env.Setup()
	db := database.New(envConfig.Database) //nolint:staticcheck
	api := server.New(db)

	api.Middleware()
	api.Router()
	api.Start(envConfig.Server.Port)
	api.Close()
	db.Close()
}
