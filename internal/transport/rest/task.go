package rest

import (
	"encoding/json"
	"github.com/dewi911/serverfn/internal/models"
	"io"
	"net/http"
)

// @Summary Create a new task
// @Description Create a new task with the input payload
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.TaskCreateInput true "Create task"
// @Success 200 {object} models.TaskResponse
// @Router /task [post]
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logError("CreateTask", "read body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp models.Task
	if err := json.Unmarshal(reqBytes, &inp); err != nil {
		logError("CreateTask", "unmarshal json", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := inp.Validate(); err != nil {
		logError("CreateTask", "validate", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := h.taskService.CreateTask(r.Context(), inp)
	if err != nil {
		logError("CreateTask", "create task", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		logError("CreateTask", "marshal json", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// @Summary Get a task by ID
// @Description Get details of a task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Router /task/{id} [get]
func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		logError("GetTask", "get id from request", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task, err := h.taskService.GetTask(r.Context(), id)
	if err != nil {
		logError("GetTask", "get task", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(task)
	if err != nil {
		logError("GetTask", "marshal json", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// @Summary Get all tasks
// @Description Get details of all tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} models.Task
// @Router /task [get]
func (h *Handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.taskService.GetAllTask(r.Context())
	if err != nil {
		logError("GetAllTasks", "get all tasks", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(tasks)
	if err != nil {
		logError("GetAllTasks", "marshal json", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// @Summary Delete a task
// @Description Delete a task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 204 "No Content"
// @Router /task/{id} [delete]
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		logError("DeleteTask", "get id from request", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.taskService.RemoveTask(r.Context(), id)
	if err != nil {
		logError("DeleteTask", "remove task", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary Update a task
// @Description Update a task's status by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body models.TaskUpdateInput true "Update task"
// @Success 200 "OK"
// @Router /task/{id} [put]
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		logError("UpdateTask", "get id from request", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp models.TaskUpdateInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		logError("UpdateTask", "decode request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.taskService.UpdateTask(r.Context(), id, inp)
	if err != nil {
		logError("UpdateTask", "update task", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
