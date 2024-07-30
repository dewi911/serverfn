package queue

import (
	"serverfn/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTaskQueue(t *testing.T) {
	q := NewTaskQueue(10)
	assert.NotNil(t, q)
	assert.IsType(t, &taskQueue{}, q)
}

func TestEnqueueDequeue(t *testing.T) {
	q := NewTaskQueue(1)

	task := &domain.Task{
		ID:     1,
		Method: "GET",
		URL:    "http://example.com",
	}

	q.Enqueue(task)

	dequeuedTask := q.Dequeue()

	assert.Equal(t, task, dequeuedTask)
}

func TestClose(t *testing.T) {
	q := NewTaskQueue(1)

	task := &domain.Task{
		ID:     1,
		Method: "GET",
		URL:    "http://example.com",
	}

	q.Enqueue(task)

	go func() {
		time.Sleep(50 * time.Millisecond)
		q.Close()
	}()

	_ = q.Dequeue()

	nilTask := q.Dequeue()
	assert.Nil(t, nilTask)
}

func TestEnqueueAfterClose(t *testing.T) {
	q := NewTaskQueue(1)

	q.Close()

	task := &domain.Task{
		ID:     1,
		Method: "GET",
		URL:    "http://example.com",
	}
	
	q.Enqueue(task)
}
