package service

import (
	"context"
	"serverfn/internal/domain"
)

type TasksService struct {
	repo domain.TaskRepository
}

func NewTasksService(repo Repositories) *TasksService {
	return &TasksService{repo: repo.GetTaskRepository()}
}

func (t *TasksService) CreateTask(ctx context.Context, inp domain.Task) (domain.TaskResponse, error) {
	task := domain.Task{
		Method:     inp.Method,
		TaskStatus: inp.TaskStatus,
		URL:        inp.URL,
		Headers:    inp.Headers,
	}

	id, err := t.repo.Create(ctx, task)
	if err != nil {
		return domain.TaskResponse{}, err
	}

	response := domain.TaskResponse{
		ID: id,
	}

	return response, err
}

func (t *TasksService) GetTask(ctx context.Context, id int64) (domain.Task, error) {
	return t.repo.GetByID(ctx, id)
}

func (t *TasksService) GetAllTask(ctx context.Context) ([]domain.Task, error) {
	return t.repo.GetAll(ctx)
}

func (t *TasksService) RemoveTask(ctx context.Context, id int64) error {
	return t.repo.Delete(ctx, id)
}

func (t *TasksService) UpdateTask(ctx context.Context, id int64, task domain.TaskUpdateInput) error {
	return t.repo.Update(ctx, id, task)
}
