package momentumservices

import (
	"errors"
	"fmt"
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
	stageService.dao = dao
	stageService.deploymentService = deploymentService
	stageService.keyValueService = keyValueService

	return stageService
}

func (ss *StageService) GetById(stageId string) (model.IStage, error) {

	record, err := ss.dao.FindRecordById(model.TABLE_STAGES_NAME, stageId)
	if err != nil {
		fmt.Println("find record by id failed:", err.Error())
		return nil, err
	}

	m, err := model.ToStage(record)
	if err != nil {
		fmt.Println("failed converting model:", err.Error())
		return nil, err
	}
	return m, nil
}

// Stages are recursive. This method returns the hierarchy starting at the highest stage until the stage with the stageId of the argument.
func (ss *StageService) GetStagesSortedTopDownById(stageId string) ([]model.IStage, bool, error) {

	isStagelessDeployment := false
	parentStageRecordId := stageId
	if parentStageRecordId == "" {
		isStagelessDeployment = true
	}

	currentStage, err := ss.GetById(parentStageRecordId)
	if err != nil {
		fmt.Println("loading stage failed:", err.Error())
		return nil, false, err
	}
	stagesSorted := []model.IStage{currentStage}

	parentStageRecordId = currentStage.ParentStageId()
	for parentStageRecordId != "" {
		currentStage, err = ss.GetById(parentStageRecordId)
		if err != nil {
			fmt.Println("loading stage failed:", err.Error())
			return nil, false, err
		}
		stagesSorted = append([]model.IStage{currentStage}, stagesSorted...)
		parentStageRecordId = currentStage.ParentStageId()
	}

	return stagesSorted, isStagelessDeployment, nil
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

func (ss *StageService) GetStagesCollection() (*models.Collection, error) {

	return ss.dao.FindCollectionByNameOrId(model.TABLE_STAGES_NAME)
}

func (ss *StageService) addParentStage(parent *models.Record, child *models.Record) error {

	if parent.Collection().Name != model.TABLE_STAGES_NAME || child.Collection().Name != model.TABLE_STAGES_NAME {
		return errors.New("can only process records of stages collection")
	}

	child.Set(model.TABLE_STAGES_FIELD_PARENTSTAGE, parent.Id)
	return ss.saveWithoutEvent(child)
}

func (ss *StageService) createWithoutEvent(name string, deploymentIds []*models.Record) (*models.Record, error) {

	stageCollection, err := ss.GetStagesCollection()
	if err != nil {
		return nil, err
	}

	stageRecord := models.NewRecord(stageCollection)
	stageRecord.Set(model.TABLE_STAGES_FIELD_NAME, name)
	stageRecord.Set(model.TABLE_STAGES_FIELD_DEPLOYMENTS, deploymentIds)

	err = ss.saveWithoutEvent(stageRecord)
	if err != nil {
		return nil, err
	}

	return stageRecord, err
}

func (ss *StageService) saveWithoutEvent(stage *models.Record) error {

	return ss.dao.Clone().SaveRecord(stage)
}
