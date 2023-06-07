package momentumcontrollers

import (
	momentumconfig "momentum/momentum-core/momentum-config"
	momentumservices "momentum/momentum-core/momentum-services"

	"github.com/pocketbase/pocketbase/models"
)

type StageController struct {
	stageService *momentumservices.StageService
}

func NewStageController(stageService *momentumservices.StageService) *StageController {

	stageController := new(StageController)
	stageController.stageService = stageService

	return stageController
}

func (sc *StageController) AddStage(record *models.Record, conf *momentumconfig.MomentumConfig) error {

	return nil
}

func (sc *StageController) UpdateStage(record *models.Record, conf *momentumconfig.MomentumConfig) error {

	return nil
}

func (sc *StageController) DeleteStage(record *models.Record, conf *momentumconfig.MomentumConfig) error {

	return nil
}
