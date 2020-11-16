package main

import (
	"ephelsa/my-career/internal/database"
	"ephelsa/my-career/internal/env"
	"ephelsa/my-career/internal/server"
)

func main() {
	envConfig := env.Setup()

	db := database.New(envConfig.Database)
	api := server.New(envConfig.Server.Port)


	db.Close()
	api.Close()
}
