package main

import (
	"log"
	"net/http"

	"github.com/TreyBastian/shed-stack/pkg/app"
	"github.com/TreyBastian/shed-stack/pkg/handlers"
	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	lp := &handlers.LandingPages{}

	app.WithRoute("/", lp.Routes())

	app.Boot()
}

func routes() http.Handler {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello World"))
	})

	return r
}
