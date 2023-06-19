package momentumcontrollers

import (
	"errors"
	config "momentum/momentum-core/momentum-config"
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

	if deploymentRecord.Collection().Name != config.TABLE_DEPLOYMENTS_NAME {
		return errors.New("can only add deployment records")
	}

	stagesSorted, isStagelessDeployment, err := dc.stageService.GetStagesSortedById(deploymentRecord.GetString(config.TABLE_DEPLOYMENTS_FIELD_PARENTSTAGE))
	if err != nil {
		return err
	}
	stageNamesSorted := make([]string, 0)
	for _, stage := range stagesSorted {
		stageNamesSorted = append(stageNamesSorted, stage.GetString(config.TABLE_STAGES_FIELD_NAME))
	}

	stageApplicationId := stagesSorted[0].GetString(config.TABLE_STAGES_FIELD_PARENTAPPLICATION)
	appRecord, err := dc.applicationService.GetById(stageApplicationId)
	if err != nil {
		return err
	}
	appName := appRecord.GetString(config.TABLE_APPLICATIONS_FIELD_NAME)
	repoId := appRecord.GetString(config.TABLE_APPLICATIONS_FIELD_PARENTREPOSITORY)

	repoRecord, err := dc.repositoryService.GetById(repoId)
	if err != nil {
		return err
	}
	repoName := repoRecord.GetString(config.TABLE_REPOSITORIES_FIELD_NAME)

	dc.deploymentService.CreateDeployment(deploymentRecord, stageNamesSorted, appName, repoName, isStagelessDeployment)
	return nil
}

func (dc *DeploymentController) UpdateDeployment(record *models.Record, conf *config.MomentumConfig) error {

	return nil
}

func (dc *DeploymentController) DeleteDeployment(record *models.Record, conf *config.MomentumConfig) error {

	return nil
}

func (dc *DeploymentController) AddRepositoryToDeployments(repositoryAddedEvent *RepositoryAddedEvent) error {

	repositoryRecord, err := dc.repositoryService.FindForName(repositoryAddedEvent.RepositoryName)
	if err != nil {
		return err
	}

	return dc.deploymentService.AddRepository(repositoryRecord, repositoryAddedEvent.Deployments)
}
