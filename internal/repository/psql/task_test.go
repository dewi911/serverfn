package psql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dewi911/serverfn/internal/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTasksRepository(db)

	tests := []struct {
		name          string
		task          domain.Task
		expectedID    int64
		expectedError error
	}{
		{
			name: "Successful creation",
			task: domain.Task{
				Method: "GET",
				URL:    "http://example.com",
				Headers: domain.Headers{
					Authentication: "Bearer token",
				},
			},
			expectedID:    1,
			expectedError: nil,
		},
		{
			name: "Failed creation",
			task: domain.Task{
				Method: "POST",
				URL:    "http://example.com/post",
			},
			expectedID:    0,
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedError == nil {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(tt.expectedID)
				mock.ExpectQuery("INSERT INTO tasks").WithArgs(tt.task.Method, tt.task.TaskStatus, tt.task.URL, sqlmock.AnyArg()).WillReturnRows(rows)
			} else {
				mock.ExpectQuery("INSERT INTO tasks").WithArgs(tt.task.Method, tt.task.TaskStatus, tt.task.URL, sqlmock.AnyArg()).WillReturnError(tt.expectedError)
			}

			id, err := repo.Create(context.Background(), tt.task)

			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedID, id)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTasksRepository(db)

	tests := []struct {
		name          string
		taskID        int64
		expectedTask  domain.Task
		expectedError error
	}{
		{
			name:   "Successful retrieval",
			taskID: 1,
			expectedTask: domain.Task{
				ID:     1,
				Method: "GET",
				URL:    "http://example.com",
			},
			expectedError: nil,
		},
		{
			name:          "Failed retrieval",
			taskID:        2,
			expectedTask:  domain.Task{},
			expectedError: sql.ErrNoRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedError == nil {
				rows := sqlmock.NewRows([]string{"id", "method", "task_status", "url", "headers"}).
					AddRow(tt.expectedTask.ID, tt.expectedTask.Method, tt.expectedTask.TaskStatus, tt.expectedTask.URL, "{}")
				mock.ExpectQuery("SELECT (.+) FROM tasks WHERE id=?").WithArgs(tt.taskID).WillReturnRows(rows)
			} else {
				mock.ExpectQuery("SELECT (.+) FROM tasks WHERE id=?").WithArgs(tt.taskID).WillReturnError(tt.expectedError)
			}

			task, err := repo.GetByID(context.Background(), tt.taskID)

			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedTask, task)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTasksRepository(db)

	tests := []struct {
		name          string
		expectedTasks []domain.Task
		expectedError error
	}{
		{
			name: "Successful retrieval",
			expectedTasks: []domain.Task{
				{ID: 1, Method: "GET", URL: "http://example.com"},
				{ID: 2, Method: "POST", URL: "http://example.com/post"},
			},
			expectedError: nil,
		},
		{
			name:          "Failed retrieval",
			expectedTasks: nil,
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedError == nil {
				rows := sqlmock.NewRows([]string{"id", "method", "task_status", "url", "headers"})
				for _, task := range tt.expectedTasks {
					rows.AddRow(task.ID, task.Method, task.TaskStatus, task.URL, "{}")
				}
				mock.ExpectQuery("SELECT (.+) FROM tasks").WillReturnRows(rows)
			} else {
				mock.ExpectQuery("SELECT (.+) FROM tasks").WillReturnError(tt.expectedError)
			}

			tasks, err := repo.GetAll(context.Background())

			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedTasks, tasks)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTasksRepository(db)

	tests := []struct {
		name          string
		taskID        int64
		expectedError error
	}{
		{
			name:          "Successful deletion",
			taskID:        1,
			expectedError: nil,
		},
		{
			name:          "Failed deletion",
			taskID:        2,
			expectedError: errors.New("task not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedError == nil {
				mock.ExpectExec("DELETE FROM tasks WHERE id=?").WithArgs(tt.taskID).WillReturnResult(sqlmock.NewResult(0, 1))
			} else {
				mock.ExpectExec("DELETE FROM tasks WHERE id=?").WithArgs(tt.taskID).WillReturnError(tt.expectedError)
			}

			err := repo.Delete(context.Background(), tt.taskID)

			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
