package handler

import (
	"database/sql"

	"github.com/go-chi/chi"
)

// MainHandler takes a db connection and pass them down to sub routers.
func MainHandler(db *sql.DB) chi.Router {
	r := chi.NewRouter()
	r.Mount("/task", CreateTaskRouter(db))
	r.Mount("/user", CreateUserRouter(db))
	
	return r
}
