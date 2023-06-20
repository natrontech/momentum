package momentumservices

import (
	"errors"
	"fmt"
	config "momentum/momentum-core/momentum-config"
	model "momentum/momentum-core/momentum-model"
	tree "momentum/momentum-core/momentum-tree"
	utils "momentum/momentum-core/momentum-utils"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

type DeploymentService struct {
	dao                *daos.Dao
	config             *config.MomentumConfig
	repositoryService  *RepositoryService
	applicationService *ApplicationService
	stageService       *StageService
	keyValueService    *KeyValueService
}

func NewDeploymentService(
	dao *daos.Dao,
	config *config.MomentumConfig,
	keyValueService *KeyValueService) *DeploymentService {

	if dao == nil {
		panic("cannot initialize service with nil dao")
	}

	deplyomentService := new(DeploymentService)

	deplyomentService.dao = dao
	deplyomentService.config = config
	deplyomentService.keyValueService = keyValueService

	return deplyomentService
}

func (ds *DeploymentService) GetById(deploymentId string) (*models.Record, error) {

	return ds.dao.FindRecordById(model.TABLE_REPOSITORIES_NAME, deploymentId)
}

func (ds *DeploymentService) AddParentStage(stage *models.Record, deployments []*models.Record) error {

	if stage.Collection().Name != model.TABLE_STAGES_NAME {
		return errors.New("stage is not record of stages collection")
	}

	for _, deployment := range deployments {
		deployment.Set(model.TABLE_DEPLOYMENTS_FIELD_PARENTSTAGE, stage.Id)
		err := ds.saveWithoutEvent(deployment)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ds *DeploymentService) AddRepository(repositoryRecord *models.Record, deployments []*models.Record) error {

	if repositoryRecord.Collection().Name != model.TABLE_REPOSITORIES_NAME {
		return errors.New("repositoryRecord is not record of repositories collection")
	}

	for _, depl := range deployments {

		depl.Set(model.TABLE_DEPLOYMENTS_FIELD_REPOSITORIES, append(depl.Get(model.TABLE_DEPLOYMENTS_FIELD_REPOSITORIES).([]string), repositoryRecord.Id))
		err := ds.saveWithoutEvent(depl)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ds *DeploymentService) CreateDeployment(
	deploymentRecord *models.Record,
	stageNamesSorted []string,
	appName string,
	repositoryName string,
	isStagelessDeployment bool) error {

	name := deploymentRecord.GetString(model.TABLE_DEPLOYMENTS_FIELD_NAME)

	repoPath := utils.BuildPath(ds.config.DataDir(), repositoryName)
	deploymentStagePath := utils.BuildPath(repoPath, appName)
	for _, s := range stageNamesSorted {
		deploymentStagePath = utils.BuildPath(deploymentStagePath, s)
	}
	fmt.Println("adding deployment for stage:", deploymentStagePath)

	repoTree, err := tree.Parse(repoPath, []string{".git"})
	if err != nil {
		return err
	}

	stageFound, stageNode := repoTree.FindStage(deploymentStagePath)
	if !stageFound {
		return errors.New("unable to find stage in tree structure")
	}

	existingDeployments := stageNode.Deployments()
	for _, deployment := range existingDeployments {
		if deployment.Path == name {
			return errors.New("unable to create deployment because name already in use")
		}
	}

	// 	1. copy template to destination (with deploymentName)
	deploymentStagePath, err = utils.DirCopy(ds.config.DeploymentTemplateDir(), utils.BuildPath(deploymentStagePath))
	if err != nil {
		return err
	}

	// 	2. replace deploymentName in files
	// 	3. replace applicationName in release yaml
	// 	4. replace repositoryName in kustomization yaml
	// 	5. add deployment yaml to kustomization.yaml of parent stage OR application (see isStagelessDeployment toggle)

	return nil
}

func (ds *DeploymentService) GetDeploymentsCollection() (*models.Collection, error) {

	coll, err := ds.dao.FindCollectionByNameOrId(model.TABLE_DEPLOYMENTS_NAME)
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
	deploymentRecord.Set(model.TABLE_DEPLOYMENTS_FIELD_NAME, name)

	err = ds.saveWithoutEvent(deploymentRecord)

	return deploymentRecord, nil
}

func (ds *DeploymentService) saveWithoutEvent(record *models.Record) error {

	return ds.dao.Clone().SaveRecord(record)
}
