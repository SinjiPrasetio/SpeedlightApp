package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// Middleware must placed before routes
	a.middleware(a.Middleware.CheckRemember)

	// Add routes here!
	a.get("/", a.Handlers.Home)

	// Static routes. DO NOT CHANGE THIS!
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
