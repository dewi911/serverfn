package psql

import (
	"database/sql"
	"serverfn/internal/domain"
)

type Repositories struct {
	tasksRepository *TasksRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		tasksRepository: NewTasksRepository(db),
	}
}

func (rs *Repositories) GetTaskRepository() domain.TaskRepository {
	return rs.tasksRepository
}
