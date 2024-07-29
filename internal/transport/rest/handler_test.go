package rest

import (
	"reflect"
	"serverfn/internal/domain"
	"testing"
)

func TestHandler_InitRouter(t *testing.T) {
	type fields struct {
		taskService domain.TaskService
	}
	tests := []struct {
		name   string
		fields fields
		want   *mux.Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				taskService: tt.fields.taskService,
			}
			if got := h.InitRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHandler(t *testing.T) {
	type args struct {
		s Services
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIdFromRequest(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIdFromRequest(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIdFromRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getIdFromRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
