package momentumcontrollers

import (
	momentumconfig "momentum/momentum-core/momentum-config"
	momentumservices "momentum/momentum-core/momentum-services"

	"github.com/pocketbase/pocketbase/models"
)

type ApplicationController struct {
	appService *momentumservices.ApplicationService
}

func NewApplicationController(appService *momentumservices.ApplicationService) *ApplicationController {

	appController := new(ApplicationController)
	appController.appService = appService

	return appController
}

func (ac *ApplicationController) AddApplication(record *models.Record, conf *momentumconfig.MomentumConfig) error {

	return nil
}

func (ac *ApplicationController) UpdateApplication(record *models.Record, conf *momentumconfig.MomentumConfig) error {

	return nil
}

func (ac *ApplicationController) DeleteApplication(record *models.Record, conf *momentumconfig.MomentumConfig) error {

	return nil
}
