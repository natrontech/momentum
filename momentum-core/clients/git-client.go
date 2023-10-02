package clients

import (
	"momentum-core/config"
	"os"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type GitClient struct {
	config *config.MomentumConfig
}

func NewGitClient(config *config.MomentumConfig) *GitClient {

	gitClient := new(GitClient)
	gitClient.config = config
	return gitClient
}

type GitTransaction struct {
	id      string
	path    string
	name    string
	headRef *plumbing.Reference
}

var transactions []*GitTransaction = make([]*GitTransaction, 0)

func CloneRepoTo(url string, username string, password string, location string) error {
	_, err := git.PlainClone(location, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		// Auth: ..., TODO in case not public dir
	})

	return err
}
