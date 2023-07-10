package momentumcontrollers

import (
	"errors"
	"fmt"
	gitclient "momentum/git-client"
	kustomizeclient "momentum/kustomize-client"
	config "momentum/momentum-core/momentum-config"
	model "momentum/momentum-core/momentum-model"
	services "momentum/momentum-core/momentum-services"
	utils "momentum/momentum-core/momentum-utils"

	"github.com/pocketbase/pocketbase/models"
)

type DeploymentController struct {
	deploymentService  *services.DeploymentService
	stageService       *services.StageService
	applicationService *services.ApplicationService
	repositoryService  *services.RepositoryService
	keyValueService    *services.KeyValueService
	kustomizeValidator *kustomizeclient.KustomizationValidationService
}

func NewDeploymentController(
	deploymentService *services.DeploymentService,
	stageService *services.StageService,
	applicationService *services.ApplicationService,
	repositoryService *services.RepositoryService,
	keyValueService *services.KeyValueService,
	kustomizeValidator *kustomizeclient.KustomizationValidationService) *DeploymentController {

	deploymentController := new(DeploymentController)
	deploymentController.deploymentService = deploymentService
	deploymentController.repositoryService = repositoryService
	deploymentController.applicationService = applicationService
	deploymentController.stageService = stageService
	deploymentController.keyValueService = keyValueService
	deploymentController.kustomizeValidator = kustomizeValidator

	return deploymentController
}

func (dc *DeploymentController) AddDeployment(deploymentRecord *models.Record, conf *config.MomentumConfig) error {

	fmt.Println("Adding deployment...")

	deploymentWithoutId, err := model.ToDeployment(deploymentRecord)
	if err != nil {
		fmt.Println("error mapping record to model:", err.Error())
		return err
	}

	stagesSorted, isStagelessDeployment, err := dc.stageService.GetStagesSortedTopDownById(deploymentWithoutId.ParentStageId())
	if err != nil {
		fmt.Println("Loading stages failed:", err.Error())
		return err
	}

	app, err := dc.applicationService.GetById(stagesSorted[0].ParentApplicationId())
	if err != nil {
		fmt.Println("Loading app failed:", err.Error())
		return err
	}

	repo, err := dc.repositoryService.GetById(app.ParentRepositoryId())
	if err != nil {
		fmt.Println("Loading repo failed:", err.Error())
		return err
	}

	gitPath := utils.BuildPath(conf.DataDir(), repo.Name())
	err = gitclient.PullRepo(gitPath)
	if err != nil {
		fmt.Println("updating repository failed:", err.Error())
		return err
	}

	err = dc.deploymentService.CreateDeployment(deploymentWithoutId, stagesSorted, app, repo, isStagelessDeployment)
	if err != nil {
		fmt.Println("creating deployment failed:", err.Error())
		return err
	}

	// TODO sync files (key values etc)
	// use keyValueService -> Use pattern with channel...

	valid, err := dc.kustomizeValidator.Validate(repo.Name())
	if err != nil {
		fmt.Println("validation has errors. manual actions required", err.Error())
		return err
	}

	if !valid {
		fmt.Println("validation failed. manual actions required")
		return errors.New("kustomize validation failed and manual actions are required now")
	}

	commitMsg := "added deployment " + deploymentWithoutId.Name()
	err = gitclient.CommitAllChangesAndPush(gitPath, commitMsg)
	if err != nil {
		fmt.Println("failed committing and pushing changes, manual actions required:", err.Error())
	}

	return nil
}

func (dc *DeploymentController) UpdateDeployment(record *models.Record, conf *config.MomentumConfig) error {

	return nil
}

func (dc *DeploymentController) DeleteDeployment(record *models.Record, conf *config.MomentumConfig) error {

	return nil
}

func (dc *DeploymentController) AddRepositoryToDeployments(repositoryAddedEvent *RepositoryAddedEvent) error {

	repositoryRecord, err := dc.repositoryService.FindByName(repositoryAddedEvent.RepositoryName)
	if err != nil {
		return err
	}

	return dc.deploymentService.AddRepository(repositoryRecord, repositoryAddedEvent.Deployments)
}
