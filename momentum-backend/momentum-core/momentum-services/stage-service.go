package momentumservices

import (
	"errors"
	model "momentum/momentum-core/momentum-model"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type StageService struct {
	dao               *daos.Dao
	deploymentService *DeploymentService
	keyValueService   *KeyValueService
}

func NewStageService(dao *daos.Dao, deploymentService *DeploymentService, keyValueService *KeyValueService) *StageService {

	if dao == nil {
		panic("cannot initialize service with nil dao")
	}

	stageService := new(StageService)
	stageService.deploymentService = deploymentService
	stageService.dao = dao
	stageService.keyValueService = keyValueService

	return stageService
}

func (ss *StageService) AddParentApplication(stageIds []string, app *models.Record) error {

	if app.Collection().Name != model.TABLE_APPLICATIONS_NAME {
		return errors.New("can only process records of applications collection")
	}

	for _, stageId := range stageIds {

		stage, err := ss.dao.FindRecordById(model.TABLE_STAGES_NAME, stageId)
		if err != nil {
			return err
		}

		stage.Set(model.TABLE_STAGES_FIELD_PARENTAPPLICATION, app.Id)
		err = ss.saveWithoutEvent(stage)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ss *StageService) addParentStage(parent *models.Record, child *models.Record) error {

	if parent.Collection().Name != model.TABLE_STAGES_NAME || child.Collection().Name != model.TABLE_STAGES_NAME {
		return errors.New("can only process records of stages collection")
	}

	child.Set(model.TABLE_STAGES_FIELD_PARENTSTAGE, parent.Id)
	return ss.saveWithoutEvent(child)
}

func (ss *StageService) GetStagesCollection() (*models.Collection, error) {

	return ss.dao.FindCollectionByNameOrId(model.TABLE_STAGES_NAME)
}

func (ss *StageService) CreateWithoutEvent(name string, deploymentIds []*models.Record) (*models.Record, error) {

	stageCollection, err := ss.GetStagesCollection()
	if err != nil {
		return nil, err
	}

	stageRecord := models.NewRecord(stageCollection)
	stageRecord.Set(model.TABLE_STAGES_FIELD_NAME, name)
	stageRecord.Set(model.TABLE_STAGES_FIELD_DEPLOYMENTS, deploymentIds)

	err = ss.saveWithoutEvent(stageRecord)

	return stageRecord, err
}

func (ss *StageService) saveWithoutEvent(stage *models.Record) error {

	return ss.dao.Clone().SaveRecord(stage)
}
