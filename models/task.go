package models

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type TaskStatus string

const (
	StatusNew       TaskStatus = "new"
	StatusInProcess TaskStatus = "in_process"
	StatusDone      TaskStatus = "done"
	StatusError     TaskStatus = "error"
)

type Task struct {
	ID              string              `json:"id"`
	Method          string              `json:"method"`
	URL             string              `json:"url"`
	Headers         map[string]string   `json:"headers"`
	Status          TaskStatus          `json:"status"`
	HTTPStatusCode  int                 `json:"httpStatusCode"`
	ResponseHeaders map[string][]string `json:"headers"`
	Length          int                 `json:"length"`
}

type TaskManager struct {
	tasks map[string]*Task
	mu    sync.Mutex
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[string]*Task),
	}
}

func (tm *TaskManager) AddTask(task *Task) string {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	task.ID = generateTaskID()
	tm.tasks[task.ID] = task
	return task.ID
}

func (tm *TaskManager) GetTask(id string) (*Task, bool) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	task, exists := tm.tasks[id]
	return task, exists
}

func (tm *TaskManager) UpdateTask(task *Task) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.tasks[task.ID] = task
}

func generateTaskID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Int())
}
