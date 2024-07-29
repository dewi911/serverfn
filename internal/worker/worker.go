package worker

import (
	"context"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"serverfn/internal/domain"
	"serverfn/internal/queue"
	"time"
)

type Worker struct {
	id        int
	taskQueue queue.TaskQueue
	taskRepo  domain.TaskRepository
	quit      chan bool
	logger    *logrus.Logger
}

func NewWorker(id int, taskQueue queue.TaskQueue, taskRepo domain.TaskRepository, logger *logrus.Logger) *Worker {
	return &Worker{
		id:        id,
		taskQueue: taskQueue,
		taskRepo:  taskRepo,
		quit:      make(chan bool),
		logger:    logger,
	}
}

func (w *Worker) Start() {
	for {
		select {
		case <-w.quit:
			return
		default:
			task := w.taskQueue.Dequeue()
			w.processTask(task)
		}
	}
}

func (w *Worker) Stop() {
	w.quit <- true
}

func (w *Worker) processTask(task *domain.Task) {
	task.TaskStatus = domain.TaskStatusInProcess

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest(task.Method, task.URL, nil)
	if err != nil {
		w.logger.WithError(err).Error("Error creating request")
		task.TaskStatus = domain.TaskStatusError
		task.Headers.Error = err.Error()
		w.updateTaskStatus(task)
		return
	}

	for key, value := range task.Headers.Headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		w.logger.WithError(err).Error("Error executing request")
		task.TaskStatus = domain.TaskStatusError
		task.Headers.Error = err.Error()
		w.updateTaskStatus(task)
		return
	}
	defer resp.Body.Close()

	task.Headers.HTTPStatusCode = resp.StatusCode
	task.Headers.ResponseHeaders = resp.Header
	task.Headers.ResponseLength = resp.ContentLength

	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		w.logger.WithError(err).Warn("Error reading response body")
	}

	task.TaskStatus = domain.TaskStatusDone
	w.updateTaskStatus(task)
}

func (w *Worker) updateTaskStatus(task *domain.Task) {
	err := w.taskRepo.Update(context.Background(), task.ID, domain.TaskUpdateInput{Status: task.TaskStatus})
	if err != nil {
		w.logger.WithError(err).Error("Failed to update task status in database")
	}
}
