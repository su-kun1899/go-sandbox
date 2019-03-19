package git

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	// メモリ上に Clone する
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/su-kun1899/go-sandbox.git",
	})
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// 特定のコミットを取り出す
	h, err := r.ResolveRevision(plumbing.Revision("2e9f025cddcb47dcc54d768f46837d36828fa4cb"))
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	cIter, err := r.Log(&git.LogOptions{From: *h})
	defer cIter.Close()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	commit, err := cIter.Next()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// assert
	got := commit.Committer.Name
	want := "su-kun1899"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("commit.Committer.Name = %v, want %v", got, want)
	}

	got = commit.Message
	want = "selectでエラーハンドリングする\n"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("commit.Message = %v, want %v", got, want)
	}
}
