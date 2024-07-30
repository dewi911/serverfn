package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"serverfn/internal/domain"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) CreateTask(ctx context.Context, task domain.Task) (domain.TaskResponse, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(domain.TaskResponse), args.Error(1)
}

func (m *MockTaskService) GetTask(ctx context.Context, id int64) (domain.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskService) GetAllTask(ctx context.Context) ([]domain.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskService) RemoveTask(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTaskService) UpdateTask(ctx context.Context, id int64, task domain.TaskUpdateInput) error {
	args := m.Called(ctx, id, task)
	return args.Error(0)
}

func TestCreateTask(t *testing.T) {
	mockService := new(MockTaskService)
	handler := &Handler{taskService: mockService}

	tests := []struct {
		name           string
		input          domain.Task
		expectedStatus int
		expectedID     int64
		mockError      error
	}{
		{
			name: "Successful creation",
			input: domain.Task{
				Method: "GET",
				URL:    "http://example.com",
			},
			expectedStatus: http.StatusOK,
			expectedID:     1,
			mockError:      nil,
		},
		{
			name: "Failed creation",
			input: domain.Task{
				Method: "POST",
				URL:    "http://example.com/post",
			},
			expectedStatus: http.StatusInternalServerError,
			expectedID:     0,
			mockError:      errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.On("CreateTask", mock.Anything, tt.input).Return(domain.TaskResponse{ID: tt.expectedID}, tt.mockError).Once()

			body, _ := json.Marshal(tt.input)
			req, _ := http.NewRequest("POST", "/task", bytes.NewBuffer(body))
			rr := httptest.NewRecorder()

			handler.CreateTask(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.mockError == nil {
				var response domain.TaskResponse
				json.Unmarshal(rr.Body.Bytes(), &response)
				assert.Equal(t, tt.expectedID, response.ID)
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestGetTask(t *testing.T) {
	mockService := new(MockTaskService)
	handler := &Handler{taskService: mockService}

	tests := []struct {
		name           string
		taskID         string
		expectedStatus int
		mockTask       domain.Task
		mockError      error
	}{
		{
			name:           "Successful retrieval",
			taskID:         "1",
			expectedStatus: http.StatusOK,
			mockTask: domain.Task{
				ID:     1,
				Method: "GET",
				URL:    "http://example.com",
			},
			mockError: nil,
		},
		{
			name:           "Task not found",
			taskID:         "2",
			expectedStatus: http.StatusInternalServerError,
			mockTask:       domain.Task{},
			mockError:      errors.New("task not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskID, _ := strconv.ParseInt(tt.taskID, 10, 64)

			mockService.On("GetTask", mock.Anything, taskID).Return(tt.mockTask, tt.mockError).Once()

			req, _ := http.NewRequest("GET", "/task/"+tt.taskID, nil)
			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/task/{id}", handler.GetTask)
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.mockError == nil {
				var response domain.Task
				json.Unmarshal(rr.Body.Bytes(), &response)
				assert.Equal(t, tt.mockTask, response)
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestGetAllTasks(t *testing.T) {
	mockService := new(MockTaskService)
	handler := &Handler{taskService: mockService}

	tests := []struct {
		name           string
		expectedStatus int
		mockTasks      []domain.Task
		mockError      error
	}{
		{
			name:           "Successful retrieval",
			expectedStatus: http.StatusOK,
			mockTasks: []domain.Task{
				{ID: 1, Method: "GET", URL: "http://example.com"},
				{ID: 2, Method: "POST", URL: "http://example.com/post"},
			},
			mockError: nil,
		},
		{
			name:           "Failed retrieval",
			expectedStatus: http.StatusInternalServerError,
			mockTasks:      nil,
			mockError:      errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.On("GetAllTask", mock.Anything).Return(tt.mockTasks, tt.mockError).Once()

			req, _ := http.NewRequest("GET", "/task", nil)
			rr := httptest.NewRecorder()

			handler.GetAllTasks(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.mockError == nil {
				var response []domain.Task
				json.Unmarshal(rr.Body.Bytes(), &response)
				assert.Equal(t, tt.mockTasks, response)
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestDeleteTask(t *testing.T) {
	mockService := new(MockTaskService)
	handler := &Handler{taskService: mockService}

	tests := []struct {
		name           string
		taskID         string
		expectedStatus int
		mockError      error
	}{
		{
			name:           "Successful deletion",
			taskID:         "1",
			expectedStatus: http.StatusNoContent,
			mockError:      nil,
		},
		{
			name:           "Failed deletion",
			taskID:         "2",
			expectedStatus: http.StatusInternalServerError,
			mockError:      errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskID, _ := strconv.ParseInt(tt.taskID, 10, 64)

			mockService.On("RemoveTask", mock.Anything, taskID).Return(tt.mockError).Once()

			req, _ := http.NewRequest("DELETE", "/task/"+tt.taskID, nil)
			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/task/{id}", handler.DeleteTask)
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			mockService.AssertExpectations(t)
		})
	}
}

func TestUpdateTask(t *testing.T) {
	mockService := new(MockTaskService)
	handler := &Handler{taskService: mockService}

	tests := []struct {
		name           string
		taskID         string
		input          domain.TaskUpdateInput
		expectedStatus int
		mockError      error
	}{
		{
			name:   "Successful update",
			taskID: "1",
			input: domain.TaskUpdateInput{
				Status: domain.TaskStatusDone,
			},
			expectedStatus: http.StatusOK,
			mockError:      nil,
		},
		{
			name:   "Failed update",
			taskID: "2",
			input: domain.TaskUpdateInput{
				Status: domain.TaskStatusError,
			},
			expectedStatus: http.StatusInternalServerError,
			mockError:      errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskID, _ := strconv.ParseInt(tt.taskID, 10, 64)

			mockService.On("UpdateTask", mock.Anything, taskID, tt.input).Return(tt.mockError).Once()

			body, _ := json.Marshal(tt.input)
			req, _ := http.NewRequest("PUT", "/task/"+tt.taskID, bytes.NewBuffer(body))
			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/task/{id}", handler.UpdateTask)
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			mockService.AssertExpectations(t)
		})
	}
}
