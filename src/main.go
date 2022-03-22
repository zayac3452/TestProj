package main

import (
	app "TestProj"
	"TestProj/internal/actions"
	config "TestProj/pkg/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	actions.InsertBooksDB(app.DB)
	app.Run(":3000")
}
