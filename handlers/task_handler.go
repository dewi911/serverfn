package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"serverfn/models"
)

type TaskHandler struct {
	manager *models.TaskManager
}

func NewTaskHandler(manager *models.TaskManager) *TaskHandler {
	return &TaskHandler{manager: manager}
}

func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	task.Status = models.StatusNew
	taskID := h.manager.AddTask(&task)

	go h.executeTask(&task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"id": taskID})
}

func (h *TaskHandler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskID := r.URL.Path[len("/task/"):]
	task, exists := h.manager.GetTask(taskID)
	if !exists {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) executeTask(task *models.Task) {
	task.Status = models.StatusInProcess
	h.manager.UpdateTask(task)

	client := &http.Client{}
	req, err := http.NewRequest(task.Method, task.URL, nil)
	if err != nil {
		task.Status = models.StatusError
		h.manager.UpdateTask(task)
		return
	}

	for key, value := range task.Headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		task.Status = models.StatusError
		h.manager.UpdateTask(task)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		task.Status = models.StatusError
		h.manager.UpdateTask(task)
		return
	}

	task.HTTPStatusCode = resp.StatusCode
	task.ResponseHeaders = resp.Header
	task.Length = len(body)
	task.Status = models.StatusDone
	h.manager.UpdateTask(task)
}
