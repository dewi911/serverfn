package service

import (
	"github.com/dewi911/serverfn/internal/models"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Repositories interface {
	GetTaskRepository() models.TaskRepository
}

type Services struct {
	taskService *TasksService
}

func (ss *Services) GetTaskService() TaskService {
	return ss.taskService
}

func NewServices(repo Repositories, tm models.TaskManager) *Services {
	return &Services{
		taskService: NewTasksService(repo.GetTaskRepository(), tm),
	}
}
