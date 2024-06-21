package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"serverfn/models"
	"testing"

	"github.com/sourcegraph/conc/pool"
)

func TestCreateTaskHandler(t *testing.T) {
	manager := models.NewTaskManager()
	taskQueue := make(chan *models.Task, 10)
	handler := NewTaskHandler(manager, taskQueue)

	p := pool.New().WithMaxGoroutines(10)

	server := httptest.NewServer(http.HandlerFunc(handler.CreateTaskHandler(p)))
	defer server.Close()

	task := models.Task{
		Method:  "GET",
		URL:     "http://google.com",
		Headers: map[string]string{"Content-Type": "application/json"},
	}

	taskBytes, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Failed to marshal task: %v", err)
	}

	resp, err := http.Post(server.URL, "application/json", bytes.NewReader(taskBytes))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var result map[string]string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if _, exists := result["id"]; !exists {
		t.Fatalf("Expected task ID in response, but got none")
	}
}

func TestCreateTaskHandler_InvalidMethod(t *testing.T) {
	manager := models.NewTaskManager()
	taskQueue := make(chan *models.Task, 10)
	handler := NewTaskHandler(manager, taskQueue)

	p := pool.New().WithMaxGoroutines(10)

	server := httptest.NewServer(http.HandlerFunc(handler.CreateTaskHandler(p)))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("Expected status %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}

func TestGetTaskHandler(t *testing.T) {
	manager := models.NewTaskManager()
	taskQueue := make(chan *models.Task, 10)
	handler := NewTaskHandler(manager, taskQueue)

	task := models.Task{
		Method:  "GET",
		URL:     "http://google.com",
		Headers: map[string]string{"Content-Type": "application/json"},
		Status:  models.StatusNew,
	}
	taskID := handler.manager.AddTask(&task)

	server := httptest.NewServer(http.HandlerFunc(handler.GetTaskHandler))
	defer server.Close()

	resp, err := http.Get(server.URL + "/task/" + taskID)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var result models.Task
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if result.Method != task.Method || result.URL != task.URL || result.Status != task.Status {
		t.Fatalf("Expected task %+v, got %+v", task, result)
	}
}

func TestGetTaskHandler_NotFound(t *testing.T) {
	manager := models.NewTaskManager()
	taskQueue := make(chan *models.Task, 10)
	handler := NewTaskHandler(manager, taskQueue)

	server := httptest.NewServer(http.HandlerFunc(handler.GetTaskHandler))
	defer server.Close()

	resp, err := http.Get(server.URL + "/task/nonexistent")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
	}
}
