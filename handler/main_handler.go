package handler

import (
	"database/sql"
	"net/http"
	customware "todos-svr/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// MainHandler takes a db connection and pass them down to sub routers.
func MainHandler(db *sql.DB) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Mount("/static", http.FileServer(http.Dir(".")))
	r.Mount("/login", createLoginRouter(db))
	r.Mount("/users", CreateUserRouter(db))
	r.Group(func(route chi.Router) {
		route.Use(customware.Authenticate())
		route.Mount("/tasks", CreateTaskRouter(db))
	})

	return r
}
