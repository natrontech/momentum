package clients

import (
	"errors"
	"fmt"
	"momentum-core/config"
	"momentum-core/utils"
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

func InitGitTransaction(path string, name string) (string, error) {

	repo, err := repository(path)
	if err != nil {
		return "", err
	}

	id, err := utils.GenerateId(path + name)
	if err != nil {
		return "", err
	}

	transaction := new(GitTransaction)

	transaction.id = id
	transaction.path = path
	transaction.name = name

	headRef, err := repo.Head()
	if err != nil {
		return "", err
	}
	transaction.headRef = headRef

	transactions = append(transactions, transaction)

	return transaction.id, nil
}

func GitTransactionWrite(path string, transactionId string) error {

	transaction, err := transaction(transactionId)
	if err != nil {
		return err
	}

	worktree, err := worktree(transaction.path)
	if err != nil {
		return err
	}

	err = add(path, worktree)
	if err != nil {
		return err
	}

	err = commit(transaction.path, "feat: added "+utils.LastPartOfPath(path)+" (ID:"+transactionId+")")
	if err != nil {
		return err
	}

	return nil
}

func GitTransactionCommit(transactionId string) error {

	t, err := transaction(transactionId)
	if err != nil {
		return err
	}

	err = push(t.path)
	if err != nil {
		return err
	}

	return nil
}

func GitTransactionRollback(transactionId string) error {

	t, err := transaction(transactionId)
	if err != nil {
		return err
	}

	repo, err := worktree(t.path)
	if err != nil {
		return err
	}

	repo.Reset(&git.ResetOptions{
		Commit: t.headRef.Hash(),
	})

	return nil
}

func transaction(transactionId string) (*GitTransaction, error) {

	for _, t := range transactions {
		if t.id == transactionId {
			return t, nil
		}
	}

	return nil, errors.New("invalid transactionId=" + transactionId)
}

func CloneRepoTo(url string, username string, password string, location string) error {
	_, err := git.PlainClone(location, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		// Auth: ..., TODO in case not public dir
	})

	return err
}

func pullRepo(location string) error {

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

func repository(location string) (*git.Repository, error) {

	repo, err := git.PlainOpen(location)
	if err != nil {
		fmt.Println("opening git repo at", location, "failed:", err.Error())
		return nil, err
	}

	return repo, nil
}

func worktree(location string) (*git.Worktree, error) {

	repo, err := repository(location)
	if err != nil {
		return nil, err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Println("reading worktree failed git repo failed:", err.Error())
		return nil, err
	}

	return worktree, nil
}

func add(location string, worktree *git.Worktree) error {

	_, err := worktree.Add(location)
	if err != nil {
		fmt.Println("adding changes in worktree failed failed:", err.Error())
		return err
	}
	return nil
}

func commit(location string, commitMsg string) error {

	worktree, err := worktree(location)
	if err != nil {
		return err
	}

	_, err = worktree.Commit(commitMsg, &git.CommitOptions{})
	if err != nil {
		fmt.Println("committing git repo failed:", err.Error())
		return err
	}

	return nil
}

func push(location string) error {

	repo, err := repository(location)
	if err != nil {
		return err
	}

	err = repo.Push(&git.PushOptions{})
	if err != nil {
		return err
	}

	return nil
}
