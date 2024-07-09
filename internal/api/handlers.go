package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CreateTaskRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type CreateTaskResponse struct {
	ID string `json:"id"`
}

func (s *Server) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.logger.WithError(err).Error("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task := s.taskManager.CreateTask(req.Method, req.URL, req.Headers)

	resp := CreateTaskResponse{ID: task.ID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) GetTaskStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["taskId"]

	task, exists := s.taskManager.GetTask(taskID)
	if !exists {
		s.logger.WithField("taskId", taskID).Warn("Task not found")
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
