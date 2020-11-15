package main

import (
	"ephelsa/my-career/config"
	"fmt"
)

func main() {
	conf := config.SetupEnvironment()

	fmt.Println("Configuration")
	fmt.Println(conf)
}
