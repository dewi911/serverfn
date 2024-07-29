package rest

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"serverfn/internal/domain"
	"strconv"
)

type Services interface {
	GetTaskService() domain.TaskService
}

type Handler struct {
	taskService domain.TaskService
}

func NewHandler(s Services) *Handler {
	return &Handler{
		taskService: s.GetTaskService(),
	}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()

	task := r.PathPrefix("/task").Subrouter()
	{
		task.HandleFunc("", h.CreateTask).Methods(http.MethodPost)
		task.HandleFunc("", h.GetAllTasks).Methods(http.MethodGet)
		task.HandleFunc("/{id:[0-9]+}", h.GetTask).Methods(http.MethodGet)
		task.HandleFunc("/{id:[0-9]+}", h.DeleteTask).Methods(http.MethodDelete)
		task.HandleFunc("/{id:[0-9]+}", h.UpdateTask).Methods(http.MethodPut)
	}

	return r
}

func getIdFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("id can't be 0")
	}

	return id, nil
}
