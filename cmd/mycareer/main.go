package main

import (
	"fmt"

	"ephelsa/my-career/pkg/config"
)

func main() {
	conf := config.SetupEnvironment()

	fmt.Println("Configuration")
	fmt.Println(conf)
}
