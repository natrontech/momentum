package momentumcontrollers

import (
	"fmt"
	config "momentum/momentum-core/momentum-config"
	model "momentum/momentum-core/momentum-model"
	services "momentum/momentum-core/momentum-services"

	"github.com/pocketbase/pocketbase/models"
)

type DeploymentController struct {
	deploymentService  *services.DeploymentService
	stageService       *services.StageService
	applicationService *services.ApplicationService
	repositoryService  *services.RepositoryService
}

func NewDeploymentController(
	deploymentService *services.DeploymentService,
	stageService *services.StageService,
	applicationService *services.ApplicationService,
	repositoryService *services.RepositoryService) *DeploymentController {

	deploymentController := new(DeploymentController)
	deploymentController.deploymentService = deploymentService
	deploymentController.repositoryService = repositoryService
	deploymentController.applicationService = applicationService
	deploymentController.stageService = stageService

	return deploymentController
}

func (dc *DeploymentController) AddDeployment(deploymentRecord *models.Record, conf *config.MomentumConfig) error {

	fmt.Println("Adding deployment...")

	// TODO GIT PULL / SYNC REPO

	deploymentWithoutId, err := model.ToDeployment(deploymentRecord)
	if err != nil {
		fmt.Println("error mapping record to model:", err.Error())
		return err
	}

	fmt.Println("Creating deployment:", deploymentWithoutId.Name(), deploymentWithoutId.ParentStageId())

	stagesSorted, isStagelessDeployment, err := dc.stageService.GetStagesSortedTopDownById(deploymentWithoutId.ParentStageId())
	if err != nil {
		fmt.Println("Loading stages failed:", err.Error())
		return err
	}

	fmt.Println("loaded stages", stagesSorted)

	app, err := dc.applicationService.GetById(stagesSorted[0].ParentApplicationId())
	if err != nil {
		fmt.Println("Loading app failed:", err.Error())
		return err
	}

	fmt.Println("loaded app")

	repo, err := dc.repositoryService.GetById(app.ParentRepositoryId())
	if err != nil {
		fmt.Println("Loading repo failed:", err.Error())
		return err
	}

	fmt.Println("loaded repo")

	err = dc.deploymentService.CreateDeployment(deploymentWithoutId, stagesSorted, app, repo, isStagelessDeployment)
	if err != nil {
		fmt.Println("creating deployment failed:", err.Error())
		return err
	}

	// TODO sync files (key values etc)

	// TODO GIT PUSH

	fmt.Println("created deployment")

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
