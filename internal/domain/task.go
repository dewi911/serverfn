package domain

import (
	"context"
	"github.com/go-playground/validator/v10"
)

func init() {
	validate = validator.New()
}

type keyType string
type TaskStatus string

const (
	TaskStatusNew       TaskStatus = "new"
	TaskStatusInProcess TaskStatus = "in_process"
	TaskStatusDone      TaskStatus = "done"
	TaskStatusError     TaskStatus = "error"
	TaskId              keyType    = "task_id"
)

var validate *validator.Validate

type TaskResponse struct {
	ID int64 `json:"id"`
}

func (t Task) Validate() error {
	return validate.Struct(t)
}

type Task struct {
	ID         int64      `json:"id" example:"1"`
	Method     string     `json:"method" example:"GET"`
	TaskStatus TaskStatus `json:"task_status" example:"new"`
	URL        string     `json:"url" example:"http://google.com"`
	Headers    Headers    `json:"headers" `
}

type TaskCreateInput struct {
	Method     string     `json:"method" example:"GET"`
	TaskStatus TaskStatus `json:"task_status" example:"new"`
	URL        string     `json:"url" example:"http://google.com"`
	Headers    Headers    `json:"headers"`
}

type TaskUpdateInput struct {
	Status TaskStatus `json:"status" validate:"required"`
}

type Headers struct {
	Authentication  string              `json:"authentication"`
	Headers         map[string]string   `json:"headers"`
	HTTPStatusCode  int                 `json:"HTTPStatusCode"`
	ResponseHeaders map[string][]string `json:"responseHeaders"`
	ResponseLength  int64               `json:"responseLength"`
	Error           string              `json:"error"`
}

type TaskService interface {
	CreateTask(ctx context.Context, task Task) (TaskResponse, error)
	GetTask(ctx context.Context, id int64) (Task, error)
	GetAllTask(ctx context.Context) ([]Task, error)
	RemoveTask(ctx context.Context, id int64) error
	UpdateTask(ctx context.Context, id int64, task TaskUpdateInput) error
}

type TaskRepository interface {
	Create(ctx context.Context, task Task) (int64, error)
	GetByID(ctx context.Context, id int64) (Task, error)
	GetAll(ctx context.Context) ([]Task, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, task TaskUpdateInput) error
}
