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
			want:    []string{"100\n"},
			wantErr: false,
		},
		{
			name:    "Read file from last line",
			args:    args{fileName: "testdata/bar.txt"},
			want:    []string{"50\n"},
			wantErr: false,
		},
		{
			name:    "Read file with CR line feed",
			args:    args{fileName: "testdata/crlf.txt"},
			want:    []string{"is CRLF\r\n"},
			wantErr: false,
		},
		{
			name:    "Empty file",
			args:    args{fileName: "testdata/empty.txt"},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Read single line file",
			args:    args{fileName: "testdata/oneline.txt"},
			want:    []string{"first line is last line\n"},
			wantErr: false,
		},
		{
			name:    "File open error",
			args:    args{fileName: "dummy.txt"},
			want:    nil,
			wantErr: true,
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
