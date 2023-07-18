package services

import (
	"errors"
	"fmt"
	"momentum-core/clients"
	"momentum-core/config"
	"momentum-core/models"
	"momentum-core/tree"
	"momentum-core/utils"
	"os"
)

type RepositoryService struct {
	config          *config.MomentumConfig
	treeService     *TreeService
	gitClient       *clients.GitClient
	kustomizeClient *clients.KustomizationValidationClient
}

func NewRepositoryService(config *config.MomentumConfig, treeService *TreeService, gitClient *clients.GitClient, kustomizeClient *clients.KustomizationValidationClient) *RepositoryService {

	repositoryService := new(RepositoryService)

	repositoryService.config = config
	repositoryService.treeService = treeService
	repositoryService.gitClient = gitClient
	repositoryService.kustomizeClient = kustomizeClient

	return repositoryService
}

func (r *RepositoryService) AddRepository(createRequest *models.RepositoryCreateRequest, traceId string) (*models.Repository, error) {

	repoPath := utils.BuildPath(r.config.DataDir(), createRequest.Name)

	if utils.FileExists(repoPath) {
		return nil, errors.New("repository with name " + createRequest.Name + " already exists")
	}

	err := clients.CloneRepoTo(createRequest.Url, "", "", repoPath)
	if err != nil {
		config.LOGGER.LogWarning("failed cloning repository", err, traceId)
		return nil, err
	}

	repo, err := tree.Parse(repoPath)
	if err != nil {
		config.LOGGER.LogWarning("failed parsing momentum tree", err, traceId)
		return nil, err
	}

	if len(repo.Directories()) != 1 {
		return nil, errors.New("only one directory allowed in repository root")
	}

	err = r.kustomizeClient.Validate(createRequest.Name)
	if err != nil {
		config.LOGGER.LogWarning("failed validating repository with kustomize", err, traceId)
		return nil, err
	}

	// TODO: GIT PUSH

	repository := new(models.Repository)
	repository.Name = createRequest.Name
	repository.Id = repo.Id
	return repository, nil
}

func (r *RepositoryService) GetRepositories(traceId string) []string {

	dir, err := utils.FileOpen(r.config.DataDir(), os.O_RDONLY)
	if err != nil {
		config.LOGGER.LogError("failed reading data directory", err, traceId)
		return make([]string, 0)
	}

	entries, err := dir.ReadDir(-1) // -1 reads all
	if err != nil {
		config.LOGGER.LogError("failed reading data directory", err, traceId)
		return make([]string, 0)
	}

	repos := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			repos = append(repos, entry.Name())
		}
	}

	return repos
}

func (r *RepositoryService) GetRepository(name string, traceId string) (*models.Repository, error) {

	fmt.Println("Logger instance:", config.LOGGER)
	config.LOGGER.LogInfo("getting repository with name "+name, "")

	repoPath := utils.BuildPath(r.config.DataDir(), name)
	fmt.Println("Repo path:", repoPath)
	if !utils.FileExists(repoPath) {
		return nil, errors.New("repository does not exist")
	}

	n, err := r.treeService.repository(name, traceId)
	if err != nil {
		config.LOGGER.LogWarning("unable to find repositories", err, traceId)
		return nil, err
	}

	result, err := models.ToRepositoryFromNode(n)
	fmt.Println(result)
	if err != nil {
		config.LOGGER.LogWarning("failed mapping repository from node", err, traceId)
		return nil, err
	}

	return result, nil
}
