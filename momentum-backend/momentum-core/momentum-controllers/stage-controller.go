package momentumcontrollers

import (
	config "momentum/momentum-core/momentum-config"
	services "momentum/momentum-core/momentum-services"

	"github.com/pocketbase/pocketbase/models"
)

type StageController struct {
	stageService *services.StageService
}

func NewStageController(stageService *services.StageService) *StageController {

	stageController := new(StageController)
	stageController.stageService = stageService

	return stageController
}

func (sc *StageController) AddStage(record *models.Record, conf *config.MomentumConfig) error {

	return nil
}

func (sc *StageController) UpdateStage(record *models.Record, conf *config.MomentumConfig) error {

	return nil
}

func (sc *StageController) DeleteStage(record *models.Record, conf *config.MomentumConfig) error {

	return nil
}
