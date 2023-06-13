package momentumcontrollers

import (
	config "momentum/momentum-core/momentum-config"
	services "momentum/momentum-core/momentum-services"

	"github.com/pocketbase/pocketbase/models"
)

type ApplicationController struct {
	appService *services.ApplicationService
}

func NewApplicationController(appService *services.ApplicationService) *ApplicationController {

	appController := new(ApplicationController)
	appController.appService = appService

	return appController
}

func (ac *ApplicationController) AddApplication(record *models.Record, conf *config.MomentumConfig) error {

	return nil
}

func (ac *ApplicationController) UpdateApplication(record *models.Record, conf *config.MomentumConfig) error {

	return nil
}

func (ac *ApplicationController) DeleteApplication(record *models.Record, conf *config.MomentumConfig) error {

	return nil
}
