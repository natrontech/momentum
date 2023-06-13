package momentumservices

import (
	"errors"
	"fmt"
	consts "momentum/momentum-core/momentum-config"
	tree "momentum/momentum-core/momentum-tree"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type DeploymentService struct {
	dao *daos.Dao
}

func NewDeploymentService(dao *daos.Dao) *DeploymentService {

	if dao == nil {
		panic("cannot initialize service with nil dao")
	}

	deplyomentService := new(DeploymentService)

	deplyomentService.dao = dao

	return deplyomentService
}

func (ds *DeploymentService) SyncDeploymentsFromDisk(n *tree.Node) ([]string, error) {

	deployments := n.AllDeployments()

	deploymentIds := make([]string, 0)
	for _, deployment := range deployments {

		deploymentId, err := ds.createWithoutEvent(deployment.Path)
		if err != nil {
			return nil, err
		}

		deploymentIds = append(deploymentIds, deploymentId)
	}

	return deploymentIds, nil
}

func (ds *DeploymentService) AddRepository(repositoryRecord *models.Record, deployments []*models.Record) error {

	if repositoryRecord.Collection().Name != consts.TABLE_REPOSITORIES_NAME {
		return errors.New("repositoryRecord is not record of repositories collection")
	}

	fmt.Println("adding repository", repositoryRecord.Id, "to deplyoments", deployments)

	for _, depl := range deployments {

		depl.Set(consts.TABLE_DEPLOYMENTS_FIELD_REPOSITORIES, append(depl.Get(consts.TABLE_DEPLOYMENTS_FIELD_REPOSITORIES).([]string), repositoryRecord.Id))
		err := ds.dao.SaveRecord(depl)
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

func (ds *DeploymentService) createWithoutEvent(name string) (string, error) {

	deploymentCollection, err := ds.GetDeploymentsCollection()
	if err != nil {
		return "", err
	}

	deploymentRecord := models.NewRecord(deploymentCollection)
	deploymentRecord.Set(consts.TABLE_DEPLOYMENTS_FIELD_NAME, name)

	err = ds.dao.Clone().SaveRecord(deploymentRecord)

	return deploymentRecord.Id, nil
}
