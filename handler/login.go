package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todos-svr/storage"

	"github.com/go-chi/chi"
	"golang.org/x/crypto/bcrypt"
)

func createLoginRouter(db *sql.DB) chi.Router {
	r := chi.NewRouter()
	r.Get("/", createLoginPage())
	r.Post("/", login(db))
	return r
}
func createLoginPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("very nice")
		http.ServeFile(w, r, "./static/")
	}
}

func login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("path in login is", r.URL.Path)
		p := storage.User{}
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(p)
		rows, err := db.Query("SELECT * FROM users WHERE Username = ?", p.Name)
		if !rows.Next() {
			w.Write([]byte("username did not exist"))
			w.WriteHeader(400)
			return
		}
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusForbidden)
			return
		}
		var encPass string
		rows.Scan(&p.ID, &p.Name, &encPass)
		fmt.Println(encPass)
		err = bcrypt.CompareHashAndPassword([]byte(encPass), []byte(p.Password))
		if err != nil {
			log.Print(err)
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusForbidden)
			return
		}
		jwt, err := p.GenerateJwtKey()
		w.Header().Add("jwt-auth-key", jwt)
		// route, _ := r.Context().Value("after_login").(string)
		// fmt.Println(route)
		// http.Redirect(w, r, route, 304)
		w.WriteHeader(200)
		return
	}
}
