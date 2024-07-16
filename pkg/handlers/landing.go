package handlers

import (
	"net/http"

	"github.com/TreyBastian/shed-stack/pkg/views"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

type LandingPages struct{}

func (l *LandingPages) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", templ.Handler(views.Hello("World")).ServeHTTP)
	return r
}
