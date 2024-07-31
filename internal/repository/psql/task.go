package psql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dewi911/serverfn/internal/models"
)

type TasksRepository struct {
	db *sql.DB
}

func NewTasksRepository(db *sql.DB) *TasksRepository {
	return &TasksRepository{
		db: db,
	}
}

func (r *TasksRepository) Create(ctx context.Context, task models.Task) (int64, error) {
	headerJSON, err := json.Marshal(task.Headers)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal headers: %w", err)
	}

	var taskId int64

	query := "INSERT INTO tasks (method, task_status, url, headers) VALUES ($1, $2, $3, $4) RETURNING id"
	err = r.db.QueryRowContext(ctx, query, task.Method, task.TaskStatus, task.URL, headerJSON).
		Scan(&taskId)

	if err != nil {
		return 0, err
	}

	return taskId, err
}

func (r *TasksRepository) GetByID(ctx context.Context, id int64) (models.Task, error) {
	var task models.Task
	var headerJSON []byte
	err := r.db.QueryRowContext(ctx, "SELECT id, method, task_status, url, headers FROM tasks WHERE id=$1", id).
		Scan(&task.ID, &task.Method, &task.TaskStatus, &task.URL, &headerJSON)
	if err != nil {
		return models.Task{}, err
	}

	err = json.Unmarshal(headerJSON, &task.Headers)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (r *TasksRepository) GetAll(ctx context.Context) ([]models.Task, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, method, task_status, url, headers FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		var headerJSON []byte
		err := rows.Scan(&task.ID, &task.Method, &task.TaskStatus, &task.URL, &headerJSON)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(headerJSON, &task.Headers)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TasksRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM tasks WHERE id=$1", id)
	return err
}

func (r *TasksRepository) Update(ctx context.Context, id int64, task models.TaskUpdateInput) error {
	_, err := r.db.ExecContext(ctx, "UPDATE tasks SET task_status=$1 WHERE id=$2", task.Status, id)
	return err
}
