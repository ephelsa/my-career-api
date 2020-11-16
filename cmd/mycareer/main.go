package main

import (
	"ephelsa/my-career/internal/database"
	"ephelsa/my-career/internal/env"
	"ephelsa/my-career/internal/server"
)

func main() {
	envConfig := env.Setup()

	api := server.NewServer(envConfig.Server.Port)
	db := database.NewDatabase(envConfig.Database)

	api.Start()
	db.Connect()
}
