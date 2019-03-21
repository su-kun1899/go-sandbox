package git_test

import (
	"github.com/su-kun1899/go-sandbox/git"
	"reflect"
	"testing"
)

func TestResolveCommit(t *testing.T) {
	type args struct {
		hash string
	}
	type want struct {
		committerName string
		commitMessage string
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name:    "resolve commit from hash",
			args:    args{hash: "2e9f025cddcb47dcc54d768f46837d36828fa4cb"},
			want:    want{committerName: "su-kun1899", commitMessage: "selectでエラーハンドリングする\n"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := git.ResolveCommit(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveCommit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Committer.Name, tt.want.committerName) {
				t.Errorf("ResolveCommit().Committer.Name = %v, want %v", got.Committer.Name, tt.want.committerName)
			}
			if !reflect.DeepEqual(got.Message, tt.want.commitMessage) {
				t.Errorf("ResolveCommit().Message = %v, want %v", got.Message, tt.want.commitMessage)
			}
		})
	}
}
