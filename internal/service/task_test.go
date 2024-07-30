package service

import (
	"context"
	"errors"
	"github.com/dewi911/serverfn/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(ctx context.Context, task domain.Task) (int64, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTaskRepository) GetByID(ctx context.Context, id int64) (domain.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetAll(ctx context.Context) ([]domain.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskRepository) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTaskRepository) Update(ctx context.Context, id int64, task domain.TaskUpdateInput) error {
	args := m.Called(ctx, id, task)
	return args.Error(0)
}

type MockTaskManager struct {
	mock.Mock
}

func (m *MockTaskManager) CreateTask(task *domain.Task) {
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
		input         domain.Task
		expectedID    int64
		expectedError error
	}{
		{
			name: "Successful creation",
			input: domain.Task{
				Method: "GET",
				URL:    "http://example.com",
			},
			expectedID:    1,
			expectedError: nil,
		},
		{
			name: "Failed creation",
			input: domain.Task{
				Method: "POST",
				URL:    "http://example.com/post",
			},
			expectedID:    0,
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.Task")).Return(tt.expectedID, tt.expectedError).Once()
			if tt.expectedError == nil {
				mockManager.On("CreateTask", mock.AnythingOfType("*domain.Task")).Once()
			}

			result, err := service.CreateTask(context.Background(), tt.input)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, domain.TaskResponse{ID: tt.expectedID}, result)
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
		expectedTask  domain.Task
		expectedError error
	}{
		{
			name:   "Successful retrieval",
			taskID: 1,
			expectedTask: domain.Task{
				ID:     1,
				Method: "GET",
				URL:    "http://example.com",
			},
			expectedError: nil,
		},
		{
			name:          "Failed retrieval",
			taskID:        2,
			expectedTask:  domain.Task{},
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
		expectedTasks []domain.Task
		expectedError error
	}{
		{
			name: "Successful retrieval",
			expectedTasks: []domain.Task{
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
		updateInput   domain.TaskUpdateInput
		expectedError error
	}{
		{
			name:   "Successful update",
			taskID: 1,
			updateInput: domain.TaskUpdateInput{
				Status: domain.TaskStatusDone,
			},
			expectedError: nil,
		},
		{
			name:   "Failed update",
			taskID: 2,
			updateInput: domain.TaskUpdateInput{
				Status: domain.TaskStatusError,
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
