package database

import (
	"github.com/dgraph-io/badger"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "get",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got interface{} = Get()

			if _, ok := got.(*badger.DB); !ok {
				t.Errorf("StatsMux() = %t, want %v", got, "badger.DB")
			}
		})
	}
}

func TestInit(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "init",
			args:    args{filename: "./test/test.db"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}