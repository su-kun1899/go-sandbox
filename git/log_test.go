package git_test

import (
	"gopkg.in/src-d/go-git.v4/utils/merkletrie"
	"reflect"
	"testing"

	"github.com/su-kun1899/go-sandbox/git"
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

func TestResolveChanges(t *testing.T) {
	type args struct {
		fromHash string
		toHash   string
	}
	type change struct {
		action   merkletrie.Action
		fileName string
	}
	type want struct {
		len     int
		changes []change
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "resolve changes",
			args: args{
				fromHash: "544ca7476346cdb2da65d28d1b75f8dabba10dbb",
				toHash:   "d8b777bab15ebf1981b8b25cbffb721835141d82",
			},
			want: want{
				len: 3,
				changes: []change{
					{action: merkletrie.Modify, fileName: "file.go"},
					{action: merkletrie.Modify, fileName: "file_test.go"},
					{action: merkletrie.Insert, fileName: "empty.txt"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from, err := git.ResolveCommit(tt.args.fromHash)
			if err != nil {
				t.Fatalf("ResolveCommit() error = %v", err)
			}
			to, err := git.ResolveCommit(tt.args.toHash)
			if err != nil {
				t.Fatalf("ResolveCommit() error = %v", err)
			}

			got, err := git.ResolveChanges(from, to)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveChanges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got.Len(), tt.want.len) {
				t.Errorf("got.Len() = %v, want %v", got.Len(), tt.want.len)
			}

			for i, c := range got {
				action, err := c.Action()
				if err != nil {
					t.Fatalf("got[%d].Action() error = %v", i, err)
				}
				if !reflect.DeepEqual(action, tt.want.changes[i].action) {
					t.Errorf("got[%d].Action() = %v, want %v", i, action, tt.want.changes[i].action)
				}

				fromFile, toFile, err := c.Files()
				if err != nil {
					t.Fatalf("got[%d].Files() error = %v", i, err)
				}

				switch action {
				case merkletrie.Insert:

					if fromFile != nil {
						t.Errorf("got[%d].Files() fromFile = %v, want %v", i, fromFile, nil)
					}
					if !reflect.DeepEqual(toFile.Name, tt.want.changes[i].fileName) {
						t.Errorf("got[%d].Files() toFile.Name = %v, want %v",
							i, toFile.Name, tt.want.changes[i].fileName)
					}
				case merkletrie.Modify:

					if fromFile == nil || toFile == nil ||
						!reflect.DeepEqual(fromFile.Name, tt.want.changes[i].fileName) ||
						!reflect.DeepEqual(toFile.Name, tt.want.changes[i].fileName) {

						t.Errorf("got[%d].Files() fromFile = %v, toFile = %v, want %v",
							i, fromFile, toFile, tt.want.changes[i].fileName)
					}
				default:
					t.Fatalf("unhandled action: %v", action)
				}
			}
		})
	}
}
