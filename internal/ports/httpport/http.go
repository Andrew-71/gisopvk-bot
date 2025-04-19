package httpport

import (
	"net/http"

	"github.com/Andrew-71/gisopvk-bot/internal/app"
	"github.com/go-chi/render"
)

type Server struct {
	app *app.Application
}

func NewHTTPServer(app *app.Application) *Server {
	return &Server{app: app}
}

func (s Server) GetReply(w http.ResponseWriter, r *http.Request) {
	var q Query
	if err := render.Decode(r, &q); err != nil {
		httpError(w, r, err, http.StatusBadRequest)
		return
	}
	reply, err := s.app.Queries.GetReply.Handle(r.Context(), q.FromApi())
	if err != nil {
		httpError(w, r, err, http.StatusBadRequest) // TODO: Possibly a 500
		return
	}
	render.JSON(w, r, mapReplyToApi(reply))
}

func httpError(w http.ResponseWriter, r *http.Request, err error, code int) {
	w.WriteHeader(code)
	render.JSON(w, r, Error{Message: err.Error()})
}
