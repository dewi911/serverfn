package models

import "testing"

func TestTask_Validate(t1 *testing.T) {
	type fields struct {
		ID         int64
		Method     string
		TaskStatus TaskStatus
		URL        string
		Headers    Headers
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Task{
				ID:         tt.fields.ID,
				Method:     tt.fields.Method,
				TaskStatus: tt.fields.TaskStatus,
				URL:        tt.fields.URL,
				Headers:    tt.fields.Headers,
			}
			if err := t.Validate(); (err != nil) != tt.wantErr {
				t1.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
