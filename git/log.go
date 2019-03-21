package git

import (
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

var repo *git.Repository

func init() {
	// メモリ上に Clone する
	r, err := git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{
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
	h, err := repo.ResolveRevision(plumbing.Revision(hash))
	if err != nil {
		return nil, err
	}
	return repo.CommitObject(*h)
}

func ResolveChanges(from, to *object.Commit) (object.Changes, error) {
	fromTree, err := from.Tree()
	if err != nil {
		return nil, err
	}
	toTree, err := to.Tree()
	if err != nil {
		return nil, err
	}

	return fromTree.Diff(toTree)
}

func ResolveBranch(name string) (*plumbing.Reference, error) {
	refs, err := repo.Storer.IterReferences()
	if err != nil {
		return nil, err
	}
	bIter := storer.NewReferenceFilteredIter(func(r *plumbing.Reference) bool {
		return (r.Name().IsBranch() || r.Name().IsRemote()) && r.Name().Short() == name
	}, refs)
	return bIter.Next()
}

func CheckoutBranch(name string) error {
	branch, err := ResolveBranch(name)
	if err != nil {
		return err
	}

	w, err := repo.Worktree()
	if err != nil {
		return err
	}

	return w.Checkout(&git.CheckoutOptions{
		Branch: branch.Name(),
	})
}

func HeadCommit() (*object.Commit, error) {
	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}
	return repo.CommitObject(ref.Hash())
}
