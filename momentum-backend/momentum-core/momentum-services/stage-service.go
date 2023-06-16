package momentumservices

import (
	"errors"
	consts "momentum/momentum-core/momentum-config"
	tree "momentum/momentum-core/momentum-tree"

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

func (ss *StageService) SyncStagesFromDisk(n *tree.Node) ([]*models.Record, error) {

	stages := n.AllStages()
	stageRecords := make([]*models.Record, 0)
	var lastStageNode *tree.Node = nil
	var lastStage *models.Record = nil
	for _, stage := range stages {

		deployments, err := ss.deploymentService.SyncDeploymentsFromDisk(n)
		if err != nil {
			return nil, err
		}

		stageRecord, err := ss.createWithoutEvent(stage.NormalizedPath(), deployments)
		if err != nil {
			return nil, err
		}

		if stage.Kind == tree.Directory {
			stageFiles := stage.Files()
			for _, f := range stageFiles {

				err := ss.keyValueService.SyncFile(f, stageRecord)
				if err != nil {
					return nil, err
				}
			}
		}

		err = ss.deploymentService.AddParentStage(stageRecord, deployments)
		if err != nil {
			return nil, err
		}

		if lastStage != nil && lastStageNode != nil && stage.Parent != nil && stage.Parent.IsStage() && lastStageNode.FullPath() == stage.Parent.FullPath() {
			err = ss.addParentStage(lastStage, stageRecord)
			if err != nil {
				return nil, err
			}
		}

		stageRecords = append(stageRecords, stageRecord)
		lastStage = stageRecord
		lastStageNode = stage
	}

	return stageRecords, nil
}

func (ss *StageService) AddParentApplication(stageIds []string, app *models.Record) error {

	if app.Collection().Name != consts.TABLE_APPLICATIONS_NAME {
		return errors.New("can only process records of applications collection")
	}

	for _, stageId := range stageIds {

		stage, err := ss.dao.FindRecordById(consts.TABLE_STAGES_NAME, stageId)
		if err != nil {
			return err
		}

		stage.Set(consts.TABLE_STAGES_FIELD_PARENTAPPLICATION, app.Id)
		err = ss.saveWithoutEvent(stage)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ss *StageService) addParentStage(parent *models.Record, child *models.Record) error {

	if parent.Collection().Name != consts.TABLE_STAGES_NAME || child.Collection().Name != consts.TABLE_STAGES_NAME {
		return errors.New("can only process records of stages collection")
	}

	child.Set(consts.TABLE_STAGES_FIELD_PARENTSTAGE, parent.Id)
	return ss.saveWithoutEvent(child)
}

func (ss *StageService) GetStagesCollection() (*models.Collection, error) {

	return ss.dao.FindCollectionByNameOrId(consts.TABLE_STAGES_NAME)
}

func (ss *StageService) createWithoutEvent(name string, deploymentIds []*models.Record) (*models.Record, error) {

	stageCollection, err := ss.GetStagesCollection()
	if err != nil {
		return nil, err
	}

	stageRecord := models.NewRecord(stageCollection)
	stageRecord.Set(consts.TABLE_STAGES_FIELD_NAME, name)
	stageRecord.Set(consts.TABLE_STAGES_FIELD_DEPLOYMENTS, deploymentIds)

	err = ss.saveWithoutEvent(stageRecord)

	return stageRecord, err
}

func (ss *StageService) saveWithoutEvent(stage *models.Record) error {

	return ss.dao.Clone().SaveRecord(stage)
}
