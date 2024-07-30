package taskmanager

import (
	"github.com/dewi911/serverfn/internal/domain"
	"github.com/dewi911/serverfn/internal/queue"
	"github.com/dewi911/serverfn/internal/worker"
	"github.com/sirupsen/logrus"
	"sync"
)

type TaskManager interface {
	Start()
	Stop()
	CreateTask(task *domain.Task)
}

type taskManager struct {
	taskQueue queue.TaskQueue
	workers   []*worker.Worker
	logger    *logrus.Logger
	wg        sync.WaitGroup
	quit      chan struct{}
}

func NewTaskManager(queueCapacity, workerCount int, taskRepo domain.TaskRepository, logger *logrus.Logger) TaskManager {
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

func (tm *taskManager) CreateTask(task *domain.Task) {
	tm.taskQueue.Enqueue(task)
	tm.logger.WithField("taskID", task.ID).Info("Task added to queue")
}
