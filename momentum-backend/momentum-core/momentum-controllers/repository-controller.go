package momentumcontrollers

import (
	"fmt"
	"strings"

	gitclient "momentum/git-client"
	config "momentum/momentum-core/momentum-config"
	services "momentum/momentum-core/momentum-services"
	tree "momentum/momentum-core/momentum-tree"
	utils "momentum/momentum-core/momentum-utils"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

type RepositoryAddedEvent struct {
	RepositoryName string
	Deployments    []*models.Record
}

type RepositoryController struct {
	repositoryService           *services.RepositoryService
	deploymentService           *services.DeploymentService
	repositoryAddedEventChannel chan *RepositoryAddedEvent
}

func NewRepositoryController(repoService *services.RepositoryService, deploymentService *services.DeploymentService, repositoryAddedEventChannel chan *RepositoryAddedEvent) *RepositoryController {

	repoController := new(RepositoryController)

	repoController.repositoryService = repoService
	repoController.deploymentService = deploymentService
	repoController.repositoryAddedEventChannel = repositoryAddedEventChannel

	return repoController
}

func (rc *RepositoryController) AddRepository(record *models.Record, conf *config.MomentumConfig) error {

	repoName := record.GetString(config.TABLE_REPOSITORIES_FIELD_NAME)
	repoUrl := record.GetString(config.TABLE_REPOSITORIES_FIELD_URL)
	path := utils.BuildPath(conf.DataDir(), strings.ReplaceAll(repoName, " ", ""))

	fmt.Println("adding repo", repoName, ", located at", repoUrl, "and to be written to", path)

	if utils.FileExists(path) {
		return apis.NewBadRequestError("repository with this name already exists", nil)
	}

	fmt.Println("Cloning repo to:", path)

	err := gitclient.PullRepoTo(repoUrl, "", "", path)
	if err != nil {
		return err
	}

	repo, err := tree.Parse(path, []string{".git"})
	if err != nil {
		return err
	}

	_, deployments, err := rc.repositoryService.SyncRepositoryFromDisk(repo, record)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		utils.DirDelete(path)
		return err
	}

	repoAddedEvent := new(RepositoryAddedEvent)
	repoAddedEvent.RepositoryName = repoName
	repoAddedEvent.Deployments = deployments

	rc.repositoryAddedEventChannel <- repoAddedEvent

	return nil
}

func (rc *RepositoryController) UpdateRepository(record *models.Record, conf *config.MomentumConfig) error {

	return nil
}

func (rc *RepositoryController) DeleteRepository(record *models.Record, conf *config.MomentumConfig) error {

	return nil
}
