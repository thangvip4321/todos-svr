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

//TaskService is a wrapper around task router, which decouples web service from underlying database implementation
type TaskService struct {
	handler *storage.TaskHandler
	router  chi.Router
}

//CreateTaskRouter instantiate a TaskService.
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

// doesnt support multi-tasks query. ( yet)
// TODO: make json Encoder support multiple tasks, make handler.GetTask(id) return multiple tasks
func (service *TaskService) getTaskFromUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("id") == "" {
			w.Write([]byte("ID not found"))
			return
		}
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Fatal(err)
		}
		err, tasks := service.handler.GetTask(id)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}
		json.NewEncoder(w).Encode(tasks)
		w.WriteHeader(200)
	}
}

func (service *TaskService) postTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task := storage.Task{}
		json.NewDecoder(r.Body).Decode(&task)
		err, response := service.handler.CreateTask(task)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(200)
	}
}
