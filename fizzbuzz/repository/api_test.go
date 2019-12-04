package repository

import (
	"github.com/jkuma/gophercises/fizzbuzz/database"
	"net/http"
	url2 "net/url"
	"testing"
)

func setUp(t *testing.T) {
	err := database.Init("./test/test.db")

	if err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {
	setUp(t)

	type args struct {
		r *http.Request
	}

	url, _ := url2.Parse("http://localhost:3000/api/toto?int1=1")

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "update",
			args:    args{r: &http.Request{URL: url}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Update(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
