package handler

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
)

func CreateUserRouter(db *sql.DB) chi.Router {
	r := chi.NewRouter()
	r.Get("/", getInfoUser(db))
	r.Post("/", registerUser(db))
	return r
}

func getInfoUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func registerUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
