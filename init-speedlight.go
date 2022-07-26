package main

import (
	"App/data"
	"App/handlers"
	"App/middleware"
	"log"
	"os"

	"github.com/SinjiPrasetio/speedlight"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Init Speedlight
	speed := &speedlight.Speedlight{}
	err = speed.New(path)
	if err != nil {
		log.Fatal(err)
	}

	speed.AppName = "myapp"

	myMiddleware := &middleware.Middleware{
		App: speed,
	}

	myHandlers := &handlers.Handlers{
		App: speed,
	}

	app := &application{
		App:        speed,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)

	myHandlers.Models = app.Models

	app.Middleware.Models = app.Models

	return app
}
