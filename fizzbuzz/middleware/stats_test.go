package middleware

import (
	"net/http"
	"testing"
)

func TestStatsMux(t *testing.T) {
	type args struct {
		mux  *http.ServeMux
		hook string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Stats", args: args{mux: http.DefaultServeMux, hook: "/api"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got interface{} = StatsMux(tt.args.mux, tt.args.hook)

			if _, ok := got.(http.HandlerFunc); !ok {
				t.Errorf("StatsMux() = %t, want %v", got, "http.HandlerFunc")
			}
		})
	}
}
