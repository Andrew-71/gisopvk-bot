package main

import (
	"net/http"

	"github.com/Andrew-71/gisopvk-bot/internal/ports/httpport"
	"github.com/Andrew-71/gisopvk-bot/internal/server"
	"github.com/Andrew-71/gisopvk-bot/internal/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	app, cleanup := service.NewApplication()
	defer cleanup()

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return httpport.HandlerFromMux(httpport.NewHTTPServer(app), router)
	})
}
