package services

import (
	"errors"
	"momentum-core/config"
	"momentum-core/models"
	"momentum-core/utils"
)

type ApplicationService struct {
	config          *config.MomentumConfig
	treeService     *TreeService
	templateService *TemplateService
}

func NewApplicationService(config *config.MomentumConfig, treeService *TreeService, templateService *TemplateService) *ApplicationService {

	appService := new(ApplicationService)

	appService.config = config
	appService.treeService = treeService
	appService.templateService = templateService

	return appService
}

func (as *ApplicationService) GetApplications(repositoryName string, traceId string) ([]*models.Application, error) {

	repo, err := as.treeService.repository(repositoryName, traceId)
	if err != nil {
		config.LOGGER.LogWarning("no repository found", err, traceId)
		return nil, err
	}

	apps := repo.Apps()
	mappedApps := make([]*models.Application, 0)
	for _, app := range apps {
		mapped, err := models.ToApplicationFromNode(app, repositoryName)
		if err != nil {
			config.LOGGER.LogWarning("failed mapping application from node", err, traceId)
			return nil, err
		}
		mappedApps = append(mappedApps, mapped)
	}

	return mappedApps, nil
}

func (as *ApplicationService) GetApplication(repositoryName string, applicationId string, traceId string) (*models.Application, error) {

	result, err := as.treeService.application(repositoryName, applicationId, traceId)
	if err != nil {
		config.LOGGER.LogWarning("no application found", err, traceId)
		return nil, err
	}

	return models.ToApplicationFromNode(result, repositoryName)
}

func (as *ApplicationService) AddApplication(request *models.ApplicationCreateRequest, traceId string) (*models.Application, error) {

	repo, err := as.treeService.repository(request.RepositoryName, traceId)
	if err != nil {
		config.LOGGER.LogWarning("no repository found", err, traceId)
		return nil, err
	}

	apps := repo.Apps()
	for _, app := range apps {
		if app.Path == request.Name {
			return nil, errors.New("application with this name already exists")
		}
	}

	appMountPath := utils.BuildPath(repo.MomentumRoot().FullPath(), request.Name)
	appBasePath := utils.BuildPath(appMountPath, "_base")

	_, err = utils.DirCopy(as.config.ApplicationTemplateFolderPath(), appMountPath)
	if err != nil {
		config.LOGGER.LogWarning("failed copying from "+as.config.ApplicationTemplateFolderPath()+" to "+appMountPath, err, traceId)
		return nil, err
	}

	appRepositoryPath := utils.BuildPath(appMountPath, "repository.yaml")
	appNamespacePath := utils.BuildPath(appMountPath, "ns.yaml")
	appBaseKustomizationPath := utils.BuildPath(appBasePath, KUSTOMIZATION_FILE_NAME)
	appBaseReleasePath := utils.BuildPath(appBasePath, "release.yaml")

	template := as.templateService.NewApplicationTemplate(appRepositoryPath, appNamespacePath, appBaseKustomizationPath, appBaseReleasePath, request.Name, request.ReconcileInterval, request.ChartVersion)
	err = as.templateService.ApplyApplicationTemplate(template)
	if err != nil {
		config.LOGGER.LogWarning("failed applying application template", err, traceId)
		return nil, err
	}

	repositoryKustomizationResources, err := as.treeService.find(traceId, request.RepositoryName, config.MOMENTUM_ROOT, KUSTOMIZATION_FILE_NAME, "resources")
	if repositoryKustomizationResources == nil || err != nil {
		return nil, errors.New("unable to find kustomization resources for repository of new application")
	}

	err = repositoryKustomizationResources.AddYamlValue(request.Name, 0)
	if err != nil {
		config.LOGGER.LogWarning("failed adding application to resources", err, traceId)
		return nil, err
	}

	err = repositoryKustomizationResources.Write(true)
	if err != nil {
		config.LOGGER.LogWarning("failed writing application to resources", err, traceId)
		return nil, err
	}

	application, err := as.treeService.find(traceId, request.RepositoryName, appMountPath)
	if err != nil {
		config.LOGGER.LogWarning("unable to find application in "+appMountPath, err, traceId)
		return nil, err
	}

	app, err := models.ToApplicationFromNode(application, request.RepositoryName)
	if err != nil {
		config.LOGGER.LogWarning("failed mapping application from node", err, traceId)
		return nil, err
	}
	return app, nil
}
