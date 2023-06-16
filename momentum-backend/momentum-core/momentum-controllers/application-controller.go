package momentumcontrollers

import (
	config "momentum/momentum-core/momentum-config"
	services "momentum/momentum-core/momentum-services"

	"github.com/pocketbase/pocketbase/models"
)

type ApplicationController struct {
	appService        *services.ApplicationService
	repositoryService *services.RepositoryService
}

func NewApplicationController(appService *services.ApplicationService, repoService *services.RepositoryService) *ApplicationController {

	appController := new(ApplicationController)
	appController.appService = appService
	appController.repositoryService = repoService

	return appController
}

func (ac *ApplicationController) AddRepositoryToApplications(repoAddedEvent *RepositoryAddedEvent) error {

	repositoryRecord, err := ac.repositoryService.FindForName(repoAddedEvent.RepositoryName)
	if err != nil {
		return err
	}

	return ac.appService.AddRepository(repositoryRecord, repoAddedEvent.Applications)
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
