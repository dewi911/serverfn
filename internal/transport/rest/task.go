package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"serverfn/internal/domain"
)

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logError("CreateTask", "read body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp domain.Task
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

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		logError("UpdateTask", "get id from request", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp domain.TaskUpdateInput
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
