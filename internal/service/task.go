package service

import (
	"context"
	"github.com/dewi911/serverfn/internal/models"
)

type TaskRepository interface {
	Create(ctx context.Context, task models.Task) (int64, error)
	GetByID(ctx context.Context, id int64) (models.Task, error)
	GetAll(ctx context.Context) ([]models.Task, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, task models.TaskUpdateInput) error
}

type TasksService struct {
	repo        TaskRepository
	taskManager models.TaskManager
}

func NewTasksService(repo TaskRepository, tm models.TaskManager) *TasksService {
	return &TasksService{
		repo:        repo,
		taskManager: tm,
	}
}

func (s *TasksService) CreateTask(ctx context.Context, inp models.Task) (models.TaskResponse, error) {
	task := &models.Task{
		Method:     inp.Method,
		TaskStatus: models.TaskStatusNew,
		URL:        inp.URL,
		Headers:    inp.Headers,
	}

	id, err := s.repo.Create(ctx, *task)
	if err != nil {
		return models.TaskResponse{}, err
	}

	task.ID = id
	s.taskManager.CreateTask(task)

	return models.TaskResponse{ID: id}, nil
}

func (s *TasksService) GetTask(ctx context.Context, id int64) (models.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TasksService) GetAllTask(ctx context.Context) ([]models.Task, error) {
	return s.repo.GetAll(ctx)
}

func (s *TasksService) RemoveTask(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *TasksService) UpdateTask(ctx context.Context, id int64, task models.TaskUpdateInput) error {
	return s.repo.Update(ctx, id, task)
}
