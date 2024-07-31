package worker

import (
	"github.com/dewi911/serverfn/internal/models"
	"github.com/dewi911/serverfn/internal/queue"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestNewWorker(t *testing.T) {
	type args struct {
		id        int
		taskQueue queue.TaskQueue
		taskRepo  models.TaskRepository
		logger    *logrus.Logger
	}
	tests := []struct {
		name string
		args args
		want *Worker
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWorker(tt.args.id, tt.args.taskQueue, tt.args.taskRepo, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWorker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorker_Start(t *testing.T) {
	type fields struct {
		id        int
		taskQueue queue.TaskQueue
		taskRepo  models.TaskRepository
		quit      chan struct{}
		logger    *logrus.Logger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Worker{
				id:        tt.fields.id,
				taskQueue: tt.fields.taskQueue,
				taskRepo:  tt.fields.taskRepo,
				quit:      tt.fields.quit,
				logger:    tt.fields.logger,
			}
			w.Start()
		})
	}
}

func TestWorker_Stop(t *testing.T) {
	type fields struct {
		id        int
		taskQueue queue.TaskQueue
		taskRepo  models.TaskRepository
		quit      chan struct{}
		logger    *logrus.Logger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Worker{
				id:        tt.fields.id,
				taskQueue: tt.fields.taskQueue,
				taskRepo:  tt.fields.taskRepo,
				quit:      tt.fields.quit,
				logger:    tt.fields.logger,
			}
			w.Stop()
		})
	}
}

func TestWorker_processTask(t *testing.T) {
	type fields struct {
		id        int
		taskQueue queue.TaskQueue
		taskRepo  models.TaskRepository
		quit      chan struct{}
		logger    *logrus.Logger
	}
	type args struct {
		task *models.Task
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Worker{
				id:        tt.fields.id,
				taskQueue: tt.fields.taskQueue,
				taskRepo:  tt.fields.taskRepo,
				quit:      tt.fields.quit,
				logger:    tt.fields.logger,
			}
			w.processTask(tt.args.task)
		})
	}
}

func TestWorker_updateTaskStatus(t *testing.T) {
	type fields struct {
		id        int
		taskQueue queue.TaskQueue
		taskRepo  models.TaskRepository
		quit      chan struct{}
		logger    *logrus.Logger
	}
	type args struct {
		task   *models.Task
		status models.TaskStatus
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Worker{
				id:        tt.fields.id,
				taskQueue: tt.fields.taskQueue,
				taskRepo:  tt.fields.taskRepo,
				quit:      tt.fields.quit,
				logger:    tt.fields.logger,
			}
			w.updateTaskStatus(tt.args.task, tt.args.status)
		})
	}
}
