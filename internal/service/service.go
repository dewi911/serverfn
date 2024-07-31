package service

import (
	"github.com/dewi911/serverfn/internal/models"
	"github.com/dewi911/serverfn/internal/transport/rest"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Repositories interface {
	GetTaskRepository() TaskRepository
}

type Services struct {
	taskService *TasksService
}

func (ss *Services) GetTaskService() rest.TaskService {
	return ss.taskService
}

func NewServices(repo Repositories, tm models.TaskManager) *Services {
	return &Services{
		taskService: NewTasksService(repo.GetTaskRepository(), tm),
	}
}
