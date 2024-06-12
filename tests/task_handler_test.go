package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"serverfn/handlers"
	"serverfn/models"
)

func TestCreateTaskHandler(t *testing.T) {
	manager := models.NewTaskManager()
	handler := handlers.NewTaskHandler(manager)

	server := httptest.NewServer(http.HandlerFunc(handler.CreateTaskHandler))
	defer server.Close()

	task := models.Task{
		Method:  "GET",
		URL:     "http://google.com",
		Headers: map[string]string{"Authentication": "Basic bG9naW46cGFzc3dvcmQ="},
	}

	data, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("failed task: %v", err)
	}

	resp, err := http.Post(server.URL+"/task", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("failed create task: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var respData map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		t.Fatalf("failed decode: %v", err)
	}

	if _, ok := respData["id"]; !ok {
		t.Fatalf("response does not have task ID")
	}
}

func TestGetTaskHandler(t *testing.T) {
	manager := models.NewTaskManager()
	handler := handlers.NewTaskHandler(manager)

	server := httptest.NewServer(http.HandlerFunc(handler.CreateTaskHandler))
	defer server.Close()

	task := models.Task{
		Method:  "GET",
		URL:     "http://google.com",
		Headers: map[string]string{"Authentication": "Basic bG9naW46cGFzc3dvcmQ="},
	}
	taskID := manager.AddTask(&task)

	server = httptest.NewServer(http.HandlerFunc(handler.GetTaskHandler))
	defer server.Close()

	time.Sleep(1 * time.Second)

	resp, err := http.Get(server.URL + "/task/" + taskID)
	if err != nil {
		t.Fatalf("failed task: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var respTask models.Task
	if err := json.NewDecoder(resp.Body).Decode(&respTask); err != nil {
		t.Fatalf("failed decode: %v", err)
	}

	if respTask.ID != taskID {
		t.Fatalf("expected task ID %s, got %s", taskID, respTask.ID)
	}
}
