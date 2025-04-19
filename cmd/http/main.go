package main

import (
	"net/http"

	// "github.com/bmstu-itstech/apollo/internal/common/server"
	// "github.com/bmstu-itstech/apollo/internal/ports/httpport"
	// "github.com/bmstu-itstech/apollo/internal/service"
	"git.a71.su/Andrew71/gisopvk-bot/internal/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	app, cleanup := service.NewApplication()
	defer cleanup()

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return httpport.HandlerFromMux(httpport.NewHTTPServer(app), router)
	})
}
