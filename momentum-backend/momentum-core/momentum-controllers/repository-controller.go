package momentumcontrollers

import (
	"errors"
	"fmt"
	"strings"

	gitclient "momentum/git-client"
	kustomizeclient "momentum/kustomize-client"
	config "momentum/momentum-core/momentum-config"
	services "momentum/momentum-core/momentum-services"
	tree "momentum/momentum-core/momentum-tree"
	utils "momentum/momentum-core/momentum-utils"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

type RepositoryAddedEvent struct {
	RepositoryName string
	Applications   []*models.Record
	Deployments    []*models.Record
}

type RepositoryController struct {
	repositorySyncService       *services.RepositorySyncService
	repositoryService           *services.RepositoryService
	repositoryAddedEventChannel chan *RepositoryAddedEvent
	kustomizeValidation         *kustomizeclient.KustomizationValidationService
}

func NewRepositoryController(
	repoSyncService *services.RepositorySyncService,
	repoService *services.RepositoryService,
	repositoryAddedEventChannel chan *RepositoryAddedEvent,
	kustomizeValidator *kustomizeclient.KustomizationValidationService) *RepositoryController {

	repoController := new(RepositoryController)

	repoController.repositorySyncService = repoSyncService
	repoController.repositoryService = repoService
	repoController.repositoryAddedEventChannel = repositoryAddedEventChannel
	repoController.kustomizeValidation = kustomizeValidator

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

	validationSuccessful, err := rc.kustomizeValidation.Validate(repoName)
	if !validationSuccessful {
		if err != nil {
			return err
		}
		return errors.New("the repository could not be validated with kustomize")
	}

	repo, err := tree.Parse(path, []string{".git"})
	if err != nil {
		return err
	}

	_, apps, deployments, err := rc.repositorySyncService.SyncRepositoryFromDisk(repo, record)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		utils.DirDelete(path)
		return err
	}

	repoAddedEvent := new(RepositoryAddedEvent)
	repoAddedEvent.RepositoryName = repoName
	repoAddedEvent.Applications = apps
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
