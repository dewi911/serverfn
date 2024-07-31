package psql

import (
	"database/sql"
	"github.com/dewi911/serverfn/internal/service"
)

type Repositories struct {
	tasksRepository *TasksRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		tasksRepository: NewTasksRepository(db),
	}
}

func (rs *Repositories) GetTaskRepository() service.TaskRepository {
	return rs.tasksRepository
}
