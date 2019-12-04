package repository

import (
	"github.com/jkuma/gophercises/fizzbuzz/database"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"
)

const path string = "./test/test.db"

func init() {
	err := os.RemoveAll(path)

	if err != nil {
		log.Fatal(err)
	}
}

func setUp(t *testing.T) {
	err := database.Init(path)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {
	setUp(t)
	defer database.Get().Close()

	type args struct {
		r *http.Request
	}

	u, _ := url.Parse("http://localhost:3000/api/toto?int1=1")

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "update",
			args:    args{r: &http.Request{URL: u}},
			wantErr: false,
		},
		{
			name:    "update1",
			args:    args{r: &http.Request{URL: u}},
			wantErr: false,
		},
		{
			name:    "update2",
			args:    args{r: &http.Request{URL: u}},
			wantErr: false,
		},
		{
			name:    "update3",
			args:    args{r: &http.Request{URL: u}},
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

func TestGet(t *testing.T) {
	setUp(t)
	defer database.Get().Close()
	
	type args struct {
		key []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "get",
			args:    args{key: []byte("http://localhost:3000/api/toto?int1=1")},
			want:    uint64ToBytes(4),
			wantErr: false,
		},
		{
			name:    "get",
			args:    args{key: []byte("http://localhost:3000/api/tata?int1=1")},
			want:    []byte(nil),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fetch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHighScore(t *testing.T) {
	setUp(t)
	defer database.Get().Close()

	tests := []struct {
		name    string
		wantKey []byte
		wantErr bool
	}{
		{
			name:    "highscore",
			wantKey: []byte("http://localhost:3000/api/toto?int1=1"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := HighScore()
			if (err != nil) != tt.wantErr {
				t.Errorf("HighScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKey, tt.wantKey) {
				t.Errorf("HighScore() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}
