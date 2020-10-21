package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"todos-svr/storage"

	"github.com/go-chi/chi"
)

//UserService is a wrapper around user router,decoupling service from database low-level implementation
type UserService struct {
	Handler storage.UserHandler
	route   chi.Router
}

// CreateUserRouter create router instance?
func CreateUserRouter(db *sql.DB) chi.Router {
	r := chi.NewRouter()
	service := UserService{storage.UserHandler{db}, r}
	r.Get("/", service.getInfoUser())
	r.Post("/register", service.registerUser())
	r.Delete("/", service.deleteUser())
	return r
}

func (service *UserService) getInfoUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := storage.User{}
		rows, err := service.Handler.Db.Query("SELECT * FROM users")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}
		for rows.Next() == true {
			rows.Scan(&user.ID, &user.Name, &user.Password)
			json.NewEncoder(w).Encode(user)
		}
		w.WriteHeader(200)
	}
}
func (service *UserService) registerUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := storage.User{}
		json.NewDecoder(r.Body).Decode(&user)
		err, newUser := service.Handler.CreateUser(&user)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}
		token, err := newUser.GenerateJwtKey()
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}
		w.Header().Add("jwt-auth-key", token)
		json.NewEncoder(w).Encode(newUser)
	}
}

func (service *UserService) deleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		err = service.Handler.DeleteUser(id)
		if err != nil {
			log.Print(err)
			w.WriteHeader(500)
			return
		}
		w.Write([]byte("deleted successfully"))
	}
}
