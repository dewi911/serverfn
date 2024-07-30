package database

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestNewPostgresConnection(t *testing.T) {
	type args struct {
		info ConnectionInfo
	}
	tests := []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPostgresConnection(tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPostgresConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostgresConnection() got = %v, want %v", got, tt.want)
			}
		})
	}
}
