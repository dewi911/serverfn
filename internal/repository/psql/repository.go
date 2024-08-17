package psql

import (
	"database/sql"
)

type Repositories struct {
	tasksRepository *TasksRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		tasksRepository: NewTasksRepository(db),
	}
}
