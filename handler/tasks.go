package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/letung3105/todos-svr/storage"
)

type TaskService struct {
	handler *storage.TaskHandler
	router  chi.Router
}

func CreateTaskRouter(db *sql.DB) TaskService {
	r := chi.NewRouter()
	service := TaskService{storage.NewHandler(db), r}
	r.Get("/", service.getTaskFromUser())
	r.Post("/", service.postTask())
	return service
}

func (service TaskService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	service.router.ServeHTTP(w, r)
}

func (service *TaskService) getTaskFromUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		newID := int64(id)
		fmt.Printf("%T\n", newID)
		if err != nil {
			log.Fatal(err)
		}
		err, task := service.handler.GetTask(id)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(task)
		w.WriteHeader(200)
	}
}

func (service *TaskService) postTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task := storage.Task{}
		json.NewDecoder(r.Body).Decode(&task)
		err, response := service.handler.CreateTask(task)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(200)
	}
}
