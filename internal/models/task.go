package models

import (
	"github.com/go-playground/validator/v10"
)

func init() {
	validate = validator.New()
}

type TaskStatus string

const (
	TaskStatusNew   TaskStatus = "new"
	TaskStatusDone  TaskStatus = "done"
	TaskStatusError TaskStatus = "error"
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
