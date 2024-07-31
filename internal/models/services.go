package models

import "context"

type TaskManager interface {
	Start()
	Stop()
	CreateTask(task *Task)
}

type TaskRepository interface {
	Create(ctx context.Context, task Task) (int64, error)
	GetByID(ctx context.Context, id int64) (Task, error)
	GetAll(ctx context.Context) ([]Task, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, task TaskUpdateInput) error
}
