package main

import (
	"ephelsa/my-career/internal/database"
	"ephelsa/my-career/internal/env"
	"ephelsa/my-career/internal/server"
)

func main() {
	envConfig := env.Setup()

	db := database.NewDatabase(envConfig.Database)
	db.Connect()
	api := server.NewServer(envConfig.Server.Port)
	api.Start()
}
