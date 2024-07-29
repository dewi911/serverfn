package psql

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"serverfn/internal/domain"
	"testing"
)

func TestNewTasksRepository(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *TasksRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTasksRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTasksRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTasksRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTasksRepository(db)

	task := domain.Task{
		Method: "GET",
		URL:    "http://example.com",
	}

	mock.ExpectQuery("INSERT INTO tasks").
		WithArgs(task.Method, task.TaskStatus, task.URL, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	id, err := repo.Create(context.Background(), task)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTasksRepository_Delete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &TasksRepository{
				db: tt.fields.db,
			}
			if err := r.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTasksRepository_GetAll(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &TasksRepository{
				db: tt.fields.db,
			}
			got, err := r.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTasksRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTasksRepository(db)

	expectedTask := domain.Task{
		ID:     1,
		Method: "GET",
		URL:    "http://example.com",
	}

	rows := sqlmock.NewRows([]string{"id", "method", "task_status", "url", "headers"}).
		AddRow(expectedTask.ID, expectedTask.Method, expectedTask.TaskStatus, expectedTask.URL, "{}")

	mock.ExpectQuery("SELECT (.+) FROM tasks WHERE id=?").
		WithArgs(1).
		WillReturnRows(rows)

	task, err := repo.GetByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedTask.ID, task.ID)
	assert.Equal(t, expectedTask.Method, task.Method)
	assert.Equal(t, expectedTask.URL, task.URL)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTasksRepository_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		id   int64
		task domain.TaskUpdateInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &TasksRepository{
				db: tt.fields.db,
			}
			if err := r.Update(tt.args.ctx, tt.args.id, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
