package service

import (
	"context"
	"reflect"
	"serverfn/internal/domain"
	"serverfn/internal/taskmanager"
	"testing"
)

func TestNewTasksService(t *testing.T) {

	type args struct {
		repo domain.TaskRepository
		tm   taskmanager.TaskManager
	}
	tests := []struct {
		name string
		args args
		want *TasksService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTasksService(tt.args.repo, tt.args.tm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTasksService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTasksService_CreateTask(t *testing.T) {
	type fields struct {
		repo        domain.TaskRepository
		taskManager taskmanager.TaskManager
	}
	type args struct {
		ctx context.Context
		inp domain.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.TaskResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TasksService{
				repo:        tt.fields.repo,
				taskManager: tt.fields.taskManager,
			}
			got, err := s.CreateTask(tt.args.ctx, tt.args.inp)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTasksService_GetAllTask(t *testing.T) {
	type fields struct {
		repo        domain.TaskRepository
		taskManager taskmanager.TaskManager
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
			s := &TasksService{
				repo:        tt.fields.repo,
				taskManager: tt.fields.taskManager,
			}
			got, err := s.GetAllTask(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTasksService_GetTask(t *testing.T) {
	type fields struct {
		repo        domain.TaskRepository
		taskManager taskmanager.TaskManager
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TasksService{
				repo:        tt.fields.repo,
				taskManager: tt.fields.taskManager,
			}
			got, err := s.GetTask(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTasksService_RemoveTask(t *testing.T) {
	type fields struct {
		repo        domain.TaskRepository
		taskManager taskmanager.TaskManager
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
			s := &TasksService{
				repo:        tt.fields.repo,
				taskManager: tt.fields.taskManager,
			}
			if err := s.RemoveTask(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("RemoveTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTasksService_UpdateTask(t *testing.T) {
	type fields struct {
		repo        domain.TaskRepository
		taskManager taskmanager.TaskManager
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
			s := &TasksService{
				repo:        tt.fields.repo,
				taskManager: tt.fields.taskManager,
			}
			if err := s.UpdateTask(tt.args.ctx, tt.args.id, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
