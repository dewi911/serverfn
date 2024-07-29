package taskmanager

import (
	"github.com/sirupsen/logrus"
	"serverfn/internal/domain"
	"serverfn/internal/queue"
	"serverfn/internal/worker"
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
}

func NewTaskManager(queueCapacity, workerCount int, taskRepo domain.TaskRepository, logger *logrus.Logger) TaskManager {
	taskQueue := queue.NewTaskQueue(queueCapacity)
	tm := &taskManager{
		taskQueue: taskQueue,
		logger:    logger,
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
}

func (tm *taskManager) Stop() {
	for _, w := range tm.workers {
		w.Stop()
	}
	tm.taskQueue.Close()
	tm.wg.Wait()
	tm.logger.Info("All workers have stopped")
}

func (tm *taskManager) CreateTask(task *domain.Task) {
	tm.taskQueue.Enqueue(task)
}
