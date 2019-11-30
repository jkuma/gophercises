package secret

import (
	"testing"
)

var key = []byte("6368616e676520746869732070617373")

func TestFileVault_Get(t *testing.T) {
	type fields struct {
		Key      []byte
		Filename string
	}

	type args struct {
		key string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "discogs", fields: fields{Key: key, Filename: "./fixtures/test.secret"},
			args: args{key: unique("discogs_api")}, want: "supermotdepasse", wantErr: false,
		},
		{
			name: "facebook", fields: fields{Key: key, Filename: "./fixtures/test.secret"},
			args: args{key: unique("facebook_api")}, want: "facebookpassword", wantErr: false,
		},
		{
			name: "erreur", fields: fields{Key: key, Filename: "./fixtures/test.secret"},
			args: args{key: unique("error")}, want: "", wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FileVault{
				Key:      tt.fields.Key,
				Filename: tt.fields.Filename,
			}
			got, err := f.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileVault_Set(t *testing.T) {
	type fields struct {
		Key      []byte
		Filename string
	}
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "discogs",
			fields:  fields{Key: key, Filename: "./fixtures/test.secret"},
			args:    args{key: "discogs_api", val: "supermotdepasse"},
			wantErr: false,
		},
		{
			name:    "erreur",
			fields:  fields{Key: key, Filename: "./fixtures/err.secret"},
			args:    args{key: "discogs_api", val: "supermotdepasse"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FileVault{
				Key:      tt.fields.Key,
				Filename: tt.fields.Filename,
			}
			if err := f.Set(tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
