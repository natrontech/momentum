package momentumservices

import (
	"errors"
	consts "momentum/momentum-core/momentum-config"
	tree "momentum/momentum-core/momentum-tree"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type DeploymentService struct {
	dao             *daos.Dao
	keyValueService *KeyValueService
}

func NewDeploymentService(dao *daos.Dao, keyValueService *KeyValueService) *DeploymentService {

	if dao == nil {
		panic("cannot initialize service with nil dao")
	}

	deplyomentService := new(DeploymentService)

	deplyomentService.dao = dao
	deplyomentService.keyValueService = keyValueService

	return deplyomentService
}

func (ds *DeploymentService) SyncDeploymentsFromDisk(n *tree.Node) ([]*models.Record, error) {

	deployments := n.AllDeployments()

	deploymentIds := make([]*models.Record, 0)
	for _, deployment := range deployments {

		deploymentRecord, err := ds.createWithoutEvent(deployment.NormalizedPath())
		if err != nil {
			return nil, err
		}

		if deployment.Kind == tree.File {

			err := ds.keyValueService.SyncFile(deployment, deploymentRecord)
			if err != nil {
				return nil, err
			}
		}

		deploymentIds = append(deploymentIds, deploymentRecord)
	}

	return deploymentIds, nil
}

func (ds *DeploymentService) AddParentStage(stage *models.Record, deployments []*models.Record) error {

	if stage.Collection().Name != consts.TABLE_STAGES_NAME {
		return errors.New("stage is not record of stages collection")
	}

	for _, deployment := range deployments {
		deployment.Set(consts.TABLE_DEPLOYMENTS_FIELD_PARENTSTAGE, stage.Id)
		err := ds.saveWithoutEvent(deployment)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ds *DeploymentService) AddRepository(repositoryRecord *models.Record, deployments []*models.Record) error {

	if repositoryRecord.Collection().Name != consts.TABLE_REPOSITORIES_NAME {
		return errors.New("repositoryRecord is not record of repositories collection")
	}

	for _, depl := range deployments {

		depl.Set(consts.TABLE_DEPLOYMENTS_FIELD_REPOSITORIES, append(depl.Get(consts.TABLE_DEPLOYMENTS_FIELD_REPOSITORIES).([]string), repositoryRecord.Id))
		err := ds.saveWithoutEvent(depl)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ds *DeploymentService) GetDeploymentsCollection() (*models.Collection, error) {

	coll, err := ds.dao.FindCollectionByNameOrId(consts.TABLE_DEPLOYMENTS_NAME)
	if err != nil {
		return nil, err
	}

	return coll, nil
}

func (ds *DeploymentService) createWithoutEvent(name string) (*models.Record, error) {

	deploymentCollection, err := ds.GetDeploymentsCollection()
	if err != nil {
		return nil, err
	}

	deploymentRecord := models.NewRecord(deploymentCollection)
	deploymentRecord.Set(consts.TABLE_DEPLOYMENTS_FIELD_NAME, name)

	err = ds.saveWithoutEvent(deploymentRecord)

	return deploymentRecord, nil
}

func (ds *DeploymentService) saveWithoutEvent(record *models.Record) error {

	return ds.dao.Clone().SaveRecord(record)
}
