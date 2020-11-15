package main

import (
	"ephelsa/my-career/internal/env"
	"ephelsa/my-career/pkg/app"
)

func main() {
	envConfig := env.Setup()
	app.StartApi(envConfig.Server)
}
