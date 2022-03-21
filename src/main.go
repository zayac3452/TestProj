package main

import (
	app "TestProj"
	config "TestProj/pkg/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
