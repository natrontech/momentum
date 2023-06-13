package momentumcontrollers

import (
	config "momentum/momentum-core/momentum-config"
	services "momentum/momentum-core/momentum-services"

	"github.com/pocketbase/pocketbase/models"
)

type DeploymentController struct {
	deploymentService *services.DeploymentService
	repositoryService *services.RepositoryService
}

func NewDeploymentController(deploymentService *services.DeploymentService, repositoryService *services.RepositoryService) *DeploymentController {

	deploymentController := new(DeploymentController)
	deploymentController.deploymentService = deploymentService
	deploymentController.repositoryService = repositoryService

	return deploymentController
}

func (dc *DeploymentController) AddDeployment(record *models.Record, conf *config.MomentumConfig) error {

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
