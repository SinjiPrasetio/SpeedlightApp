package main

import (
	"App/data"
	"App/handlers"
	"App/middleware"

	"github.com/SinjiPrasetio/speedlight"
)

type application struct {
	App        *speedlight.Speedlight
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
