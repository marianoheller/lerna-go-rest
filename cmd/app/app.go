package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	_ "github.com/mattn/go-sqlite3"

	server "learn/internal/api/http/server"
	config "learn/internal/config"
	domain "learn/internal/domain"
	infra "learn/internal/infrastructure"
)

var PORT = getPort()

func main() {
	port := getPort()

	config.InitialiseDatabase()

	r := setupServer()

	http.ListenAndServe(":"+port, r)
}

func setupServer() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	repository := infra.NewSQLiteRepository(config.GetDBConnection())
	service := domain.NewDomainService(repository)
	server.Init(service, r)

	return r
}

func getPort() string {
	port, present := os.LookupEnv("PORT")
	if !present {
		panic("PORT env var not set")
	}
	return port
}
