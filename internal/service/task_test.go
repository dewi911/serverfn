package service

import (
	"context"
	"errors"
	"github.com/dewi911/serverfn/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(ctx context.Context, task models.Task) (int64, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTaskRepository) GetByID(ctx context.Context, id int64) (models.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockTaskRepository) GetAll(ctx context.Context) ([]models.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockTaskRepository) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTaskRepository) Update(ctx context.Context, id int64, task models.TaskUpdateInput) error {
	args := m.Called(ctx, id, task)
	return args.Error(0)
}

type MockTaskManager struct {
	mock.Mock
}

func (m *MockTaskManager) CreateTask(task *models.Task) {
	m.Called(task)
}

func (m *MockTaskManager) Start() {
	m.Called()
}
func (m *MockTaskManager) Stop() {
	m.Called()
}

func TestCreateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	mockManager := new(MockTaskManager)
	service := NewTasksService(mockRepo, mockManager)

	tests := []struct {
		name          string
		input         models.Task
		expectedID    int64
		expectedError error
	}{
		{
			name: "Successful creation",
			input: models.Task{
				Method: "GET",
				URL:    "http://example.com",
			},
			expectedID:    1,
			expectedError: nil,
		},
		{
			name: "Failed creation",
			input: models.Task{
				Method: "POST",
				URL:    "http://example.com/post",
			},
			expectedID:    0,
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Create", mock.Anything, mock.AnythingOfType("models.Task")).Return(tt.expectedID, tt.expectedError).Once()
			if tt.expectedError == nil {
				mockManager.On("CreateTask", mock.AnythingOfType("*models.Task")).Once()
			}

			result, err := service.CreateTask(context.Background(), tt.input)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, models.TaskResponse{ID: tt.expectedID}, result)
			}

			mockRepo.AssertExpectations(t)
			mockManager.AssertExpectations(t)
		})
	}
}

func TestGetTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	mockManager := new(MockTaskManager)
	service := NewTasksService(mockRepo, mockManager)

	tests := []struct {
		name          string
		taskID        int64
		expectedTask  models.Task
		expectedError error
	}{
		{
			name:   "Successful retrieval",
			taskID: 1,
			expectedTask: models.Task{
				ID:     1,
				Method: "GET",
				URL:    "http://example.com",
			},
			expectedError: nil,
		},
		{
			name:          "Failed retrieval",
			taskID:        2,
			expectedTask:  models.Task{},
			expectedError: errors.New("task not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("GetByID", mock.Anything, tt.taskID).Return(tt.expectedTask, tt.expectedError).Once()

			result, err := service.GetTask(context.Background(), tt.taskID)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedTask, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetAllTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	mockManager := new(MockTaskManager)
	service := NewTasksService(mockRepo, mockManager)

	tests := []struct {
		name          string
		expectedTasks []models.Task
		expectedError error
	}{
		{
			name: "Successful retrieval",
			expectedTasks: []models.Task{
				{ID: 1, Method: "GET", URL: "http://example.com"},
				{ID: 2, Method: "POST", URL: "http://example.com/post"},
			},
			expectedError: nil,
		},
		{
			name:          "Failed retrieval",
			expectedTasks: nil,
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("GetAll", mock.Anything).Return(tt.expectedTasks, tt.expectedError).Once()

			result, err := service.GetAllTask(context.Background())

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedTasks, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestRemoveTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	mockManager := new(MockTaskManager)
	service := NewTasksService(mockRepo, mockManager)

	tests := []struct {
		name          string
		taskID        int64
		expectedError error
	}{
		{
			name:          "Successful removal",
			taskID:        1,
			expectedError: nil,
		},
		{
			name:          "Failed removal",
			taskID:        2,
			expectedError: errors.New("task not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Delete", mock.Anything, tt.taskID).Return(tt.expectedError).Once()

			err := service.RemoveTask(context.Background(), tt.taskID)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	mockManager := new(MockTaskManager)
	service := NewTasksService(mockRepo, mockManager)

	tests := []struct {
		name          string
		taskID        int64
		updateInput   models.TaskUpdateInput
		expectedError error
	}{
		{
			name:   "Successful update",
			taskID: 1,
			updateInput: models.TaskUpdateInput{
				Status: models.TaskStatusDone,
			},
			expectedError: nil,
		},
		{
			name:   "Failed update",
			taskID: 2,
			updateInput: models.TaskUpdateInput{
				Status: models.TaskStatusError,
			},
			expectedError: errors.New("task not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Update", mock.Anything, tt.taskID, tt.updateInput).Return(tt.expectedError).Once()

			err := service.UpdateTask(context.Background(), tt.taskID, tt.updateInput)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
