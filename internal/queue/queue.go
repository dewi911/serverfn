// internal/queue/queue.go

package queue

import (
	"github.com/dewi911/serverfn/internal/domain"
	"sync"
)

type TaskQueue interface {
	Enqueue(*domain.Task)
	Dequeue() *domain.Task
	Close()
}

type taskQueue struct {
	tasks chan *domain.Task
	wg    sync.WaitGroup
}

func NewTaskQueue(capacity int) TaskQueue {
	return &taskQueue{
		tasks: make(chan *domain.Task, capacity),
	}
}

func (q *taskQueue) Enqueue(task *domain.Task) {
	q.wg.Add(1)
	go func() {
		defer q.wg.Done()
		q.tasks <- task
	}()
}

func (q *taskQueue) Dequeue() *domain.Task {
	return <-q.tasks
}

func (q *taskQueue) Close() {
	close(q.tasks)
	q.wg.Wait()
}
