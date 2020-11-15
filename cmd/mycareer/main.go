package main

import (
	"ephelsa/my-career/pkg/app"
	"ephelsa/my-career/pkg/env"
)

func main() {
	envConfig := env.Setup()
	app.StartApi(envConfig.Server)
}
