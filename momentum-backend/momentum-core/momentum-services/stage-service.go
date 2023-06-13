package momentumservices

import (
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

func (ss *StageService) SyncStagesFromDisk(n *tree.Node) ([]string, error) {

	stages := n.AllStages()
	stageIds := make([]string, 0)
	for _, stage := range stages {

		deploymentIds, err := ss.deploymentService.SyncDeploymentsFromDisk(n)
		if err != nil {
			return nil, err
		}

		stageId, stageRecord, err := ss.createWithoutEvent(stage.NormalizedPath(), deploymentIds)
		if err != nil {
			return nil, err
		}

		if stage.Kind == tree.Directory {
			stageFiles := stage.Files()
			for _, f := range stageFiles {
				err = ss.keyValueService.SyncFile(f, stageRecord)
				if err != nil {
					return nil, err
				}
			}
		}

		stageIds = append(stageIds, stageId)
	}

	return stageIds, nil
}

func (ss *StageService) GetStagesCollection() (*models.Collection, error) {

	return ss.dao.FindCollectionByNameOrId(consts.TABLE_STAGES_NAME)
}

func (ss *StageService) createWithoutEvent(name string, deploymentIds []string) (string, *models.Record, error) {

	stageCollection, err := ss.GetStagesCollection()
	if err != nil {
		return "", nil, err
	}

	stageRecord := models.NewRecord(stageCollection)
	stageRecord.Set(consts.TABLE_STAGES_FIELD_NAME, name)
	stageRecord.Set(consts.TABLE_STAGES_FIELD_DEPLOYMENTS, deploymentIds)

	err = ss.dao.Clone().SaveRecord(stageRecord)

	return stageRecord.Id, stageRecord, err
}
