package file_test

import (
	"reflect"
	"testing"
	"github.com/su-kun1899/go-sandbox/file"
)

func TestReadFromLast(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "Read file from last line",
			args:    args{fileName: "testdata/foo.txt"},
			want:    []string{"100"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := file.ReadFromLast(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFromLast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFromLast() = %v, want %v", got, tt.want)
			}
		})
	}
}
