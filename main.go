package main

import (
	"ephelsa/my-career/config"
	"fmt"
)

func main() {
	conf := config.EnvironmentConfig()

	fmt.Println("Configuration")
	fmt.Println(conf)
}
