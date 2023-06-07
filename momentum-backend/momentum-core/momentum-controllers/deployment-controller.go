package momentumcontrollers

import (
	momentumconfig "momentum/momentum-core/momentum-config"
	momentumservices "momentum/momentum-core/momentum-services"

	"github.com/pocketbase/pocketbase/models"
)

type DeploymentController struct {
	deploymentService *momentumservices.DeplyomentService
}

func NewDeploymentController(deploymentService *momentumservices.DeplyomentService) *DeploymentController {

	deploymentController := new(DeploymentController)
	deploymentController.deploymentService = deploymentService

	return deploymentController
}

func (dc *DeploymentController) AddDeplyoment(record *models.Record, conf *momentumconfig.MomentumConfig) error {

	return nil
}

func (dc *DeploymentController) UpdateDeplyoment(record *models.Record, conf *momentumconfig.MomentumConfig) error {

	return nil
}

func (dc *DeploymentController) DeleteDeplyoment(record *models.Record, conf *momentumconfig.MomentumConfig) error {

	return nil
}
