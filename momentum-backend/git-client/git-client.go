package gitclient

import (
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

func CloneRepoTo(url string, username string, password string, location string) error {
	_, err := git.PlainClone(location, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		// Auth: ..., TODO in case not public dir
	})

	return err
}

func PullRepo(location string) error {

	repo, err := git.PlainOpen(location)
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil && err.Error() != "already up-to-date" {
		// TODO is there a more elegant solution?
		return err
	}
	return nil
}

func CommitAllChangesAndPush(location string, commitMsg string) error {

	repo, err := git.PlainOpen(location)
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	_, err = worktree.Add(location)
	if err != nil {
		return err
	}

	_, err = worktree.Commit(commitMsg, &git.CommitOptions{})
	if err != nil {
		return err
	}

	return nil
}
