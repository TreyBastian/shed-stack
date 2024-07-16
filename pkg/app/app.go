package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	host string
	port uint16

	routes chi.Router
}

const (
	APP_HOST = "APP_HOST"
	APP_PORT = "APP_PORT"
)

func New() (*App, error) {
	p, err := strconv.ParseUint(os.Getenv(APP_PORT), 10, 16)
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.CleanPath)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/ping"))

	return &App{host: os.Getenv(APP_HOST), port: uint16(p), routes: router}, nil
}

func (a *App) WithRoute(prefix string, route http.Handler) {
	a.routes.Mount(prefix, route)
}

func (a *App) Boot() {
	server := &http.Server{Addr: a.address(), Handler: a.routes}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	go func() {
		<-sig
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	log.Printf("Server started at http://%s", a.address())
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-serverCtx.Done()
}

func (a *App) address() string {
	return fmt.Sprintf("%s:%d", a.host, a.port)
}
