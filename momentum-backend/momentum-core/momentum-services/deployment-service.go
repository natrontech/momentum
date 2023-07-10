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
	dao             *daos.Dao
	config          *config.MomentumConfig
	keyValueService *KeyValueService
	templateService *TemplateService
}

func NewDeploymentService(
	dao *daos.Dao,
	config *config.MomentumConfig,
	keyValueService *KeyValueService,
	templateService *TemplateService) *DeploymentService {

	if dao == nil {
		panic("cannot initialize service with nil dao")
	}

	deplyomentService := new(DeploymentService)

	deplyomentService.dao = dao
	deplyomentService.config = config
	deplyomentService.keyValueService = keyValueService
	deplyomentService.templateService = templateService

	return deplyomentService
}

func (ds *DeploymentService) GetById(deploymentId string) (model.IDeployment, error) {

	record, err := ds.dao.FindRecordById(model.TABLE_REPOSITORIES_NAME, deploymentId)
	if err != nil {
		return nil, err
	}

	return model.ToDeployment(record)
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
	deployment model.IDeployment,
	stagesSorted []model.IStage,
	application model.IApplication,
	repository model.IRepository,
	isStagelessDeployment bool) error {

	repoPath := utils.BuildPath(ds.config.DataDir(), repository.Name())
	deploymentStagePath := utils.BuildPath(repoPath, application.Name())
	for _, s := range stagesSorted {
		deploymentStagePath = utils.BuildPath(deploymentStagePath, s.Name())
	}

	repoTree, err := tree.Parse(repoPath, []string{".git"})
	if err != nil {
		fmt.Println("failed parsing repository:", err.Error())
		return err
	}

	stageFound, stageNode := repoTree.FindStage(deploymentStagePath)
	if !stageFound {
		return errors.New("unable to find stage in tree structure")
	}

	existingDeployments := stageNode.Deployments()
	for _, depl := range existingDeployments {
		if depl.Path == deployment.Name() {
			return errors.New("unable to create deployment because name already in use")
		}
	}

	kustomizationResources, found := stageNode.FindFirst(tree.ToMatchableSearchTerm(utils.BuildPath(stageNode.FullPath(), KUSTOMIZATION_FILE_NAME, "resources")))
	if !found {
		return errors.New("unable to find kustomization resources for stage or application of new deployment")
	}

	deploymentYamlDestinationName := deployment.Name() + ".yaml"
	deploymentFolderDestinationPath := utils.BuildPath(deploymentStagePath, "_deploy", deployment.Name())
	deploymentFileDestinationPath := utils.BuildPath(deploymentStagePath, deploymentYamlDestinationName)

	deploymentFolderDestinationPath, err = utils.DirCopy(ds.config.DeploymentTemplateFolderPath(), deploymentFolderDestinationPath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fileCopySuccess := utils.FileCopy(ds.config.DeploymentTemplateFilePath(), deploymentFileDestinationPath)
	if !fileCopySuccess {
		return errors.New("failed copying deployment file")
	}

	releaseYamlPath := utils.BuildPath(deploymentFolderDestinationPath, "release.yaml")
	deploymentKustomizationYamlPath := utils.BuildPath(deploymentFolderDestinationPath, KUSTOMIZATION_FILE_NAME)

	err = ds.templateService.ApplyDeploymentKustomizationTemplate(deploymentKustomizationYamlPath, deployment.Name())
	if err != nil {
		fmt.Println("template for", deploymentKustomizationYamlPath, "failed:", err.Error())
		return err
	}
	err = ds.templateService.ApplyDeploymentReleaseTemplate(releaseYamlPath, application.Name())
	if err != nil {
		fmt.Println("template for", deploymentKustomizationYamlPath, "failed:", err.Error())
		return err
	}
	err = ds.templateService.ApplyDeploymentStageDeploymentDescriptionTemplate(deploymentFileDestinationPath, deployment.Name(), repository.Name())
	if err != nil {
		fmt.Println("template for", deploymentKustomizationYamlPath, "failed:", err.Error())
		return err
	}

	err = kustomizationResources.AddSequenceValue(deploymentYamlDestinationName, 0)
	if err != nil {
		fmt.Println("failed adding deployment to resources:", err.Error())
		return err
	}

	err = kustomizationResources.Write(true)
	if err != nil {
		fmt.Println("failed writing deployment to resources:", err.Error())
		return err
	}

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
	if err != nil {
		return nil, err
	}

	return deploymentRecord, nil
}

func (ds *DeploymentService) saveWithoutEvent(record *models.Record) error {

	return ds.dao.Clone().SaveRecord(record)
}
