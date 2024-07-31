// internal/queue/queue.go

package queue

import (
	"github.com/dewi911/serverfn/internal/models"
	"sync"
)

type TaskQueue interface {
	Enqueue(*models.Task)
	Dequeue() *models.Task
	Close()
}

type taskQueue struct {
	tasks chan *models.Task
	wg    sync.WaitGroup
}

func NewTaskQueue(capacity int) TaskQueue {
	return &taskQueue{
		tasks: make(chan *models.Task, capacity),
	}
}

func (q *taskQueue) Enqueue(task *models.Task) {
	q.wg.Add(1)
	go func() {
		defer q.wg.Done()
		q.tasks <- task
	}()
}

func (q *taskQueue) Dequeue() *models.Task {
	return <-q.tasks
}

func (q *taskQueue) Close() {
	close(q.tasks)
	q.wg.Wait()
}
