package worker

import (
	"context"
	"github.com/dewi911/serverfn/internal/domain"
	"github.com/dewi911/serverfn/internal/queue"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Worker struct {
	id        int
	taskQueue queue.TaskQueue
	taskRepo  domain.TaskRepository
	quit      chan struct{}
	logger    *logrus.Logger
}

func NewWorker(id int, taskQueue queue.TaskQueue, taskRepo domain.TaskRepository, logger *logrus.Logger) *Worker {
	return &Worker{
		id:        id,
		taskQueue: taskQueue,
		taskRepo:  taskRepo,
		quit:      make(chan struct{}),
		logger:    logger,
	}
}

func (w *Worker) Start() {
	w.logger.WithField("workerID", w.id).Info("Worker started")
	for {
		select {
		case <-w.quit:
			w.logger.WithField("workerID", w.id).Info("Worker stopping")
			return
		default:
			task := w.taskQueue.Dequeue()
			if task != nil {
				w.processTask(task)
			}
		}
	}
}

func (w *Worker) Stop() {
	close(w.quit)
}

func (w *Worker) processTask(task *domain.Task) {
	w.logger.WithFields(logrus.Fields{
		"workerID": w.id,
		"taskID":   task.ID,
	}).Info("Processing task")

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest(task.Method, task.URL, nil)
	if err != nil {
		w.logger.WithError(err).Error("Error creating request")
		w.updateTaskStatus(task, domain.TaskStatusError)
		return
	}

	for key, value := range task.Headers.Headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		w.logger.WithError(err).Error("Error executing request")
		w.updateTaskStatus(task, domain.TaskStatusError)
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

	w.updateTaskStatus(task, domain.TaskStatusDone)
}

func (w *Worker) updateTaskStatus(task *domain.Task, status domain.TaskStatus) {
	err := w.taskRepo.Update(context.Background(), task.ID, domain.TaskUpdateInput{Status: status})
	if err != nil {
		w.logger.WithError(err).Error("Failed to update task status in database")
	} else {
		w.logger.WithFields(logrus.Fields{
			"taskID": task.ID,
			"status": status,
		}).Info("Task status updated")
	}
}
