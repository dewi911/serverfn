package mocks

import (
	"serverfn/internal/task"

	"github.com/stretchr/testify/mock"
)

type MockTaskManager struct {
	mock.Mock
}

func (m *MockTaskManager) Start() {
	m.Called()
}

func (m *MockTaskManager) Stop() {
	m.Called()
}

func (m *MockTaskManager) CreateTask(method, url string, headers map[string]string) *task.Task {
	args := m.Called(method, url, headers)
	return args.Get(0).(*task.Task)
}

func (m *MockTaskManager) GetTask(id string) (*task.Task, bool) {
	args := m.Called(id)
	return args.Get(0).(*task.Task), args.Bool(1)
}
