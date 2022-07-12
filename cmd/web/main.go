package main

import (
	"github.com/hmdrzaa11/example-go-api/pkg/kernel"
	"github.com/hmdrzaa11/example-go-api/pkg/routes"
	"github.com/joho/godotenv"
)

func main() {
	//load our environemt variuable
	if err := godotenv.Load(); err != nil {
		panic("no .env found!")
	}
	app := kernel.Boot()

	//setup routes
	routes.LoadRoutes(app)

	//handles listening
	go func() {
		app.Run()
	}()

	app.GracefullShutdown()
}
