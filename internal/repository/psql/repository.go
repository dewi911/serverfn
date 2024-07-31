package psql

import (
	"database/sql"
	"github.com/dewi911/serverfn/internal/models"
)

type Repositories struct {
	tasksRepository *TasksRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		tasksRepository: NewTasksRepository(db),
	}
}

func (rs *Repositories) GetTaskRepository() models.TaskRepository {
	return rs.tasksRepository
}
