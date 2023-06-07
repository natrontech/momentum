package gitclient

import (
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

func PullRepoTo(url string, username string, password string, location string) error {
	_, err := git.PlainClone(location, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		// Auth: ..., TODO in case not public dir
	})

	return err
}
