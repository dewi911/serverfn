package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"serverfn/internal/task"
	"serverfn/tests/mocks"
)

func TestCreateTask(t *testing.T) {
	mockManager := new(mocks.MockTaskManager)
	logger := logrus.New()
	server := NewServer(mockManager, logger)

	mockTask := &task.Task{
		ID:     "test-id",
		Status: task.TaskStatusNew,
		Method: "GET",
		URL:    "https://example.com",
	}

	mockManager.On("CreateTask", "GET", "https://example.com", mock.Anything).Return(mockTask)

	reqBody := CreateTaskRequest{
		Method: "GET",
		URL:    "https://example.com",
	}
	body, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/task", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.CreateTask)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response CreateTaskResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "test-id", response.ID)

	mockManager.AssertExpectations(t)
}

func TestGetTaskStatus(t *testing.T) {
	mockManager := new(mocks.MockTaskManager)
	logger := logrus.New()
	server := NewServer(mockManager, logger)

	mockTask := &task.Task{
		ID:     "test-id",
		Status: task.TaskStatusDone,
		Method: "GET",
		URL:    "https://example.com",
	}

	mockManager.On("GetTask", "test-id").Return(mockTask, true)

	req, err := http.NewRequest("GET", "/task/test-id", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/task/{taskId}", server.GetTaskStatus)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response task.Task
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "test-id", response.ID)
	assert.Equal(t, task.TaskStatusDone, response.Status)

	mockManager.AssertExpectations(t)
}
