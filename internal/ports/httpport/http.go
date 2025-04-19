package httpport

import (
	"net/http"

	// "github.com/bmstu-itstech/apollo/internal/app/query"
	// "github.com/bmstu-itstech/apollo/internal/domain/material"
	"git.a71.su/Andrew71/gisopvk-bot/internal/app"
	"git.a71.su/Andrew71/gisopvk-bot/internal/app/query"
	"github.com/go-chi/render"
	// openapi_types "github.com/oapi-codegen/runtime/types"
)

type Server struct {
	app *app.Application
}

func NewHTTPServer(app *app.Application) *Server {
	return &Server{app: app}
}

func (s Server) GetReply(w http.ResponseWriter, r *http.Request) {
	q, err := s.app.Queries.GetReply.Handle(r.Context(), query.Query{})
}

func (s Server) GetDepartments(w http.ResponseWriter, r *http.Request) {
	q, err := s.app.Queries.GetDepartments.Handle(r.Context(), query.GetDepartments{})
	if err != nil {
		httpError(w, r, err, http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, mapDepartmentsToApi(q))
}