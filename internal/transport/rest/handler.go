package rest

import (
	"errors"
	_ "github.com/dewi911/serverfn/docs"
	"github.com/dewi911/serverfn/internal/service"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"strconv"
)

// @title Server API
// @version 1.0
// @description This is a server API.
// @host localhost:8080
// @BasePath /
type Services interface {
	GetTaskService() service.TaskService
}

type Handler struct {
	taskService service.TaskService
}

func NewHandler(s Services) *Handler {
	return &Handler{
		taskService: s.GetTaskService(),
	}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
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
