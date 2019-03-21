package git

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

var repo *git.Repository

func init() {
	// メモリ上に Clone する
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		// TODO 環境変数とかから取り出す
		URL: "https://github.com/su-kun1899/go-sandbox.git",
	})
	if err != nil {
		panic(err)
	}

	repo = r
}

// TODO go-git の型は隠蔽してインターフェイスで扱うのがよさそう
func ResolveCommit(hash string) (*object.Commit, error) {
	// 特定のコミットを取り出す
	h, err := repo.ResolveRevision(plumbing.Revision(hash))
	if err != nil {
		return nil, err
	}
	cIter, err := repo.Log(&git.LogOptions{From: *h})
	defer cIter.Close()
	if err != nil {
		return nil, err
	}
	return cIter.Next()
}
