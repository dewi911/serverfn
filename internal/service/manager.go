package service

import (
	"github.com/dewi911/serverfn/internal/models"
	"github.com/dewi911/serverfn/utils/queue"
	"github.com/dewi911/serverfn/utils/worker"
	"github.com/sirupsen/logrus"
	"sync"
)

type taskManager struct {
	taskQueue queue.TaskQueue
	workers   []*worker.Worker
	logger    *logrus.Logger
	wg        sync.WaitGroup
	quit      chan struct{}
}

func NewTaskManager(queueCapacity, workerCount int, taskRepo models.TaskRepository, logger *logrus.Logger) models.TaskManager {
	taskQueue := queue.NewTaskQueue(queueCapacity)
	tm := &taskManager{
		taskQueue: taskQueue,
		logger:    logger,
		quit:      make(chan struct{}),
	}

	tm.workers = make([]*worker.Worker, workerCount)
	for i := 0; i < workerCount; i++ {
		tm.workers[i] = worker.NewWorker(i, taskQueue, taskRepo, logger)
	}

	return tm
}

func (tm *taskManager) Start() {
	for _, w := range tm.workers {
		tm.wg.Add(1)
		go func(worker *worker.Worker) {
			defer tm.wg.Done()
			worker.Start()
		}(w)
	}
	tm.logger.Info("Task manager started")
}

func (tm *taskManager) Stop() {
	tm.logger.Info("Stopping task manager...")
	close(tm.quit)
	for _, w := range tm.workers {
		w.Stop()
	}
	tm.taskQueue.Close()
	tm.wg.Wait()
	tm.logger.Info("Task manager stopped")
}

func (tm *taskManager) CreateTask(task *models.Task) {
	tm.taskQueue.Enqueue(task)
	tm.logger.WithField("taskID", task.ID).Info("Task added to queue")
}
