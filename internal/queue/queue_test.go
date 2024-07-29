package queue

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewTaskQueue(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want TaskQueue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskQueue(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskQueue_Close(t *testing.T) {
	type fields struct {
		tasks chan *domain.Task
		wg    sync.WaitGroup
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &taskQueue{
				tasks: tt.fields.tasks,
				wg:    tt.fields.wg,
			}
			q.Close()
		})
	}
}

func Test_taskQueue_Dequeue(t *testing.T) {
	type fields struct {
		tasks chan *domain.Task
		wg    sync.WaitGroup
	}
	tests := []struct {
		name   string
		fields fields
		want   *domain.Task
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &taskQueue{
				tasks: tt.fields.tasks,
				wg:    tt.fields.wg,
			}
			if got := q.Dequeue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskQueue_Enqueue(t *testing.T) {
	type fields struct {
		tasks chan *domain.Task
		wg    sync.WaitGroup
	}
	type args struct {
		task *domain.Task
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &taskQueue{
				tasks: tt.fields.tasks,
				wg:    tt.fields.wg,
			}
			q.Enqueue(tt.args.task)
		})
	}
}
