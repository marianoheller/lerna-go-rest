package server

import (
	d "learn/internal/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Init(service *d.DomainService, r *chi.Mux) {
	NewArticlesApi(service, r)
	NewAdminApi(r)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin: index"))
	})
}
