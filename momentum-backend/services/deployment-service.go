package services

import (
	"errors"
	"momentum-core/config"
	"momentum-core/models"
	"momentum-core/tree"
	"momentum-core/utils"
)

type DeploymentService struct {
	config          *config.MomentumConfig
	stageService    *StageService
	templateService *TemplateService
	treeService     *TreeService
}

func NewDeploymentService(config *config.MomentumConfig, stageService *StageService, templateService *TemplateService, treeService *TreeService) *DeploymentService {

	deploymentService := new(DeploymentService)

	deploymentService.config = config
	deploymentService.stageService = stageService
	deploymentService.templateService = templateService
	deploymentService.treeService = treeService

	return deploymentService
}

func (ds *DeploymentService) GetDeployments(repositoryName string, traceId string) ([]*models.Deployment, error) {

	repo, err := ds.treeService.repository(repositoryName, traceId)
	if err != nil {
		config.LOGGER.LogWarning("no repository found", err, traceId)
		return nil, err
	}

	deployments := repo.AllDeployments()

	mappedDeployments := make([]*models.Deployment, 0)
	for _, deployment := range deployments {
		mapped, err := models.ToDeploymentFromNode(deployment, repo.Id)
		if err != nil {
			config.LOGGER.LogWarning("failed mapping deployment from node", err, traceId)
			return nil, err
		}
		mappedDeployments = append(mappedDeployments, mapped)
	}

	return mappedDeployments, nil
}

func (ds *DeploymentService) GetDeployment(repositoryName string, deploymentId string, traceId string) (*models.Deployment, error) {

	deployment, err := ds.treeService.deployment(repositoryName, deploymentId, traceId)
	if err != nil {
		config.LOGGER.LogWarning("unable to find deployment "+deploymentId, err, traceId)
		return nil, err
	}
	if deployment == nil {
		return nil, errors.New("no deployment with id " + deploymentId)
	}

	return models.ToDeploymentFromNode(deployment, deployment.Root().Id)
}

func (ds *DeploymentService) AddDeployment(request *models.DeploymentCreateRequest, traceId string) (*models.Deployment, error) {

	stage, err := ds.stageService.FindStageById(request.ParentStageId, request.RepositoryName, traceId)
	if err != nil {
		config.LOGGER.LogWarning("unable to find stage with id "+request.ParentStageId, err, traceId)
		return nil, err
	}

	if ds.deploymentExists(request.Name, stage.Id, request.RepositoryName, traceId) {
		return nil, errors.New("unable to create deployment because name already in use")
	}

	deploymentYamlDestinationName := request.Name + ".yaml"
	deploymentFolderDestinationPath := utils.BuildPath(stage.Path, "_deploy", request.Name)
	deploymentFileDestinationPath := utils.BuildPath(stage.Path, deploymentYamlDestinationName)

	deploymentFolderDestinationPath, err = utils.DirCopy(ds.config.DeploymentTemplateFolderPath(), deploymentFolderDestinationPath)
	if err != nil {
		config.LOGGER.LogWarning("failed copying deployment template from "+ds.config.DeploymentTemplateFolderPath()+" to "+deploymentFolderDestinationPath, err, traceId)
		return nil, err
	}

	fileCopySuccess := utils.FileCopy(ds.config.DeploymentTemplateFilePath(), deploymentFileDestinationPath)
	if !fileCopySuccess {
		return nil, errors.New("failed copying deployment file")
	}

	releaseYamlPath := utils.BuildPath(deploymentFolderDestinationPath, "release.yaml")
	deploymentKustomizationYamlPath := utils.BuildPath(deploymentFolderDestinationPath, KUSTOMIZATION_FILE_NAME)

	template := ds.templateService.NewDeploymentTemplate(deploymentKustomizationYamlPath, deploymentFileDestinationPath, releaseYamlPath, deploymentYamlDestinationName, request.ApplicationName, request.RepositoryName)
	err = ds.templateService.ApplyDeploymentTemplate(template)
	if err != nil {
		config.LOGGER.LogWarning("failed applying deployment template", err, traceId)
		return nil, err
	}

	parentStageKustomizationResources, err := ds.treeService.find(traceId, request.RepositoryName, stage.Path, KUSTOMIZATION_FILE_NAME, "resources")
	if err != nil {
		config.LOGGER.LogError("failed searching for parent stage kustomization yaml", err, traceId)
		return nil, err
	}

	if parentStageKustomizationResources != nil {
		err = parentStageKustomizationResources.AddValue(deploymentYamlDestinationName, 0)
		if err != nil {
			config.LOGGER.LogWarning("failed adding deployment to resources", err, traceId)
			return nil, err
		}

		tree.Print(parentStageKustomizationResources)

		err = parentStageKustomizationResources.Write(true)
		if err != nil {
			config.LOGGER.LogError("failed writing deployment to resources", err, traceId)
			return nil, err
		}
	} else {
		parentStageKustomization, err := ds.treeService.find(traceId, request.RepositoryName, stage.Path, KUSTOMIZATION_FILE_NAME)
		if err != nil {
			config.LOGGER.LogError("unable to add resources sequence beacuse not found parents stage kustomization", err, traceId)
			return nil, err
		}

		err = parentStageKustomization.AddSequence("resources", []string{deploymentYamlDestinationName}, 0)
		if err != nil {
			config.LOGGER.LogError("unable to add resources sequence to parents stage kustomization", err, traceId)
			return nil, err
		}

		tree.Print(parentStageKustomization)

		err = parentStageKustomization.Write(true)
		if err != nil {
			config.LOGGER.LogError("failed writing deployment to resources", err, traceId)
			return nil, err
		}
	}

	deployment, err := ds.treeService.find(traceId, request.RepositoryName, deploymentFileDestinationPath)
	if err != nil {
		config.LOGGER.LogWarning("unable to find deployment", err, traceId)
		return nil, err
	}

	depl, err := models.ToDeploymentFromNode(deployment, request.RepositoryName)
	if err != nil {
		config.LOGGER.LogWarning("failed mapping deployment from node", err, traceId)
		return nil, err
	}
	return depl, nil
}

func (ds *DeploymentService) deploymentExists(deploymentName string, stageId string, repositoryName string, traceId string) bool {

	deployments, err := ds.findDeploymentsByStage(stageId, repositoryName, traceId)
	if err != nil {
		config.LOGGER.LogWarning("unable to find deployments of stage "+stageId, err, traceId)
		return true
	}

	for _, depl := range deployments {
		if depl.Name == deploymentName {
			return true
		}
	}

	return false
}

func (ds *DeploymentService) findDeploymentsByStage(stageId string, repositoryName string, traceId string) ([]*models.Deployment, error) {

	stage, err := ds.treeService.stage(repositoryName, stageId, traceId)
	if err != nil {
		config.LOGGER.LogWarning("unable to find stage "+stageId, err, traceId)
		return nil, err
	}
	deployments := stage.Deployments()

	depls := make([]*models.Deployment, 0)
	for _, deployment := range deployments {
		depl, err := models.ToDeploymentFromNode(deployment, repositoryName)
		if err != nil {
			config.LOGGER.LogWarning("failed mapping deployment from node", err, traceId)
			return nil, err
		}
		depls = append(depls, depl)
	}

	return depls, nil
}
