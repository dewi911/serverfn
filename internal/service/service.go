package service

import (
	"serverfn/internal/domain"
	"serverfn/internal/taskmanager"
)

type Repositories interface {
	GetTaskRepository() domain.TaskRepository
}

type Services struct {
	taskService *TasksService
}

func (ss *Services) GetTaskService() domain.TaskService {
	return ss.taskService
}

func NewServices(repo Repositories, tm taskmanager.TaskManager) *Services {
	return &Services{
		taskService: NewTasksService(repo.GetTaskRepository(), tm),
	}
}
