package service

import "serverfn/internal/domain"

type Repositories interface {
	GetTaskRepository() domain.TaskRepository
}

type Services struct {
	taskService *TasksService
}

func (ss *Services) GetTaskService() domain.TaskService {
	return ss.taskService
}

func NewServices(repo Repositories) *Services {
	return &Services{
		taskService: NewTasksService(repo),
	}
}
