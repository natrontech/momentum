package momentumcontrollers

import (
	"fmt"
	"strings"

	gitclient "momentum/git-client"
	momentumconfig "momentum/momentum-core/momentum-config"
	momentumservices "momentum/momentum-core/momentum-services"
	tree "momentum/momentum-core/momentum-tree"
	utils "momentum/momentum-core/momentum-utils"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

type RepositoryController struct {
	applicationService *momentumservices.ApplicationService
}

func NewRepositoryController(appService *momentumservices.ApplicationService) *RepositoryController {

	repoController := new(RepositoryController)
	repoController.applicationService = appService

	return repoController
}

func (rc *RepositoryController) AddRepository(record *models.Record, conf *momentumconfig.MomentumConfig) error {

	repoName := record.GetString(momentumconfig.TABLE_REPOSITORIES_FIELD_NAME)
	repoUrl := record.GetString(momentumconfig.TABLE_REPOSITORIES_FIELD_URL)
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

	tree.Print(repo)

	err = rc.applicationService.AddApplications(repo, record)
	if err != nil {
		return apis.NewApiError(500, err.Error(), nil)
	}

	return nil
}

func (rc *RepositoryController) UpdateRepository(record *models.Record, conf *momentumconfig.MomentumConfig) error {

	return nil
}

func (rc *RepositoryController) DeleteRepository(record *models.Record, conf *momentumconfig.MomentumConfig) error {

	return nil
}
