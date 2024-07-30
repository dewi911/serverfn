package taskmanager

import (
	"bytes"
	"context"
	"github.com/dewi911/serverfn/internal/domain"
	"github.com/dewi911/serverfn/internal/worker"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sync"
	"testing"
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

type MockTaskQueue struct {
	mock.Mock
}

func (m *MockTaskQueue) Enqueue(task *domain.Task) {
	m.Called(task)
}

func (m *MockTaskQueue) Dequeue() *domain.Task {
	args := m.Called()
	return args.Get(0).(*domain.Task)
}

func (m *MockTaskQueue) Close() {
	m.Called()
}

func TestNewTaskManager(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	logger := logrus.New()

	tm := NewTaskManager(10, 2, mockRepo, logger)

	assert.NotNil(t, tm)
	assert.IsType(t, &taskManager{}, tm)
}

func TestCreateTask(t *testing.T) {
	mockQueue := new(MockTaskQueue)

	var buf bytes.Buffer
	logger := logrus.New()
	logger.SetOutput(&buf)

	tm := &taskManager{
		taskQueue: mockQueue,
		logger:    logger,
	}

	task := &domain.Task{
		ID:     1,
		Method: "GET",
		URL:    "http://example.com",
	}

	mockQueue.On("Enqueue", task).Once()

	tm.CreateTask(task)

	mockQueue.AssertExpectations(t)
	
	assert.Contains(t, buf.String(), "Task added to queue")
	assert.Contains(t, buf.String(), "taskID=1")
}

func TestStop(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	mockQueue := new(MockTaskQueue)
	logger := logrus.New()

	tm := &taskManager{
		taskQueue: mockQueue,
		workers:   make([]*worker.Worker, 2),
		logger:    logger,
		wg:        sync.WaitGroup{},
		quit:      make(chan struct{}),
	}

	for i := range tm.workers {
		tm.workers[i] = worker.NewWorker(i, mockQueue, mockRepo, logger)
	}

	mockQueue.On("Close").Once()

	tm.Stop()

	mockQueue.AssertExpectations(t)

	select {
	case <-tm.quit:
		// Канал закрыт, как и ожидалось
	default:
		t.Error("quit channel was not closed")
	}
}
