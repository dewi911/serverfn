package service

import (
	"context"
	"serverfn/internal/domain"
	"serverfn/internal/taskmanager"
)

type TasksService struct {
	repo        domain.TaskRepository
	taskManager taskmanager.TaskManager
}

func NewTasksService(repo domain.TaskRepository, tm taskmanager.TaskManager) *TasksService {
	return &TasksService{
		repo:        repo,
		taskManager: tm,
	}
}

func (s *TasksService) CreateTask(ctx context.Context, inp domain.Task) (domain.TaskResponse, error) {
	task := &domain.Task{
		Method:     inp.Method,
		TaskStatus: domain.TaskStatusNew,
		URL:        inp.URL,
		Headers:    inp.Headers,
	}

	id, err := s.repo.Create(ctx, *task)
	if err != nil {
		return domain.TaskResponse{}, err
	}

	task.ID = id
	s.taskManager.CreateTask(task)

	return domain.TaskResponse{ID: id}, nil
}

func (s *TasksService) GetTask(ctx context.Context, id int64) (domain.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TasksService) GetAllTask(ctx context.Context) ([]domain.Task, error) {
	return s.repo.GetAll(ctx)
}

func (s *TasksService) RemoveTask(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *TasksService) UpdateTask(ctx context.Context, id int64, task domain.TaskUpdateInput) error {
	return s.repo.Update(ctx, id, task)
}
