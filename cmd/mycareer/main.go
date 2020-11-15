package main

import (
	app2 "ephelsa/my-career/pkg/app"
	"ephelsa/my-career/pkg/env"
)

func main() {
	envConfig := env.Setup()
	app2.StartApi(envConfig.Server)
}
