package services

import (
	"errors"
	"momentum-core/config"
	"momentum-core/models"
	"strings"
)

type ValueService struct {
	treeService *TreeService
}

func NewValueService(treeService *TreeService) *ValueService {

	valueService := new(ValueService)

	valueService.treeService = treeService

	return valueService
}

func (vs *ValueService) ValueById(repositoryName string, valueId string, traceId string) (*models.Value, error) {

	value, err := vs.treeService.value(repositoryName, valueId, traceId)
	if err != nil {
		config.LOGGER.LogWarning("unable to find value "+valueId, err, traceId)
		return nil, err
	}
	if value == nil {
		return nil, errors.New("no value with id " + valueId)
	}

	return models.ToValueFromNode(value)
}

func (vs *ValueService) ValuesByApplication(repositoryName string, applicationId string, traceId string) ([]*models.ValueWrapper, error) {

	application, err := vs.treeService.application(repositoryName, applicationId, traceId)
	if err != nil {
		return nil, err
	}

	wrappedValues := make([]*models.ValueWrapper, 0)
	files := application.Files()
	for _, f := range files {
		if strings.EqualFold(f.NormalizedPath(), KUSTOMIZATION_FILE_NAME) {
			wrappedKustomization, err := models.ToValueWrapperFromNode(f, models.APPLICATION)
			if err != nil {
				return nil, err
			}
			wrappedValues = append(wrappedValues, wrappedKustomization)
		}
		if strings.EqualFold(f.PathWithoutEnding(), "ns") {
			wrappedNamespace, err := models.ToValueWrapperFromNode(f, models.APPLICATION)
			if err != nil {
				return nil, err
			}
			wrappedValues = append(wrappedValues, wrappedNamespace)
		}
		if strings.EqualFold(f.PathWithoutEnding(), "repository") {
			wrappedRepository, err := models.ToValueWrapperFromNode(f, models.APPLICATION)
			if err != nil {
				return nil, err
			}
			wrappedValues = append(wrappedValues, wrappedRepository)
		}
	}

	return wrappedValues, nil
}

func (vs *ValueService) ValuesByStage(repositoryName string, stageId string, traceId string) ([]*models.ValueWrapper, error) {

	stage, err := vs.treeService.stage(repositoryName, stageId, traceId)
	if err != nil {
		return nil, err
	}

	files := stage.Files()
	for _, f := range files {
		if strings.EqualFold(f.NormalizedPath(), KUSTOMIZATION_FILE_NAME) {
			wrappedKustomization, err := models.ToValueWrapperFromNode(f, models.STAGE)
			if err != nil {
				return nil, err
			}
			return []*models.ValueWrapper{wrappedKustomization}, nil
		}
	}

	return make([]*models.ValueWrapper, 0), nil
}

func (vs *ValueService) ValuesByDeployment(repositoryName string, deploymentId string, traceId string) ([]*models.ValueWrapper, error) {

	deployment, err := vs.treeService.deployment(repositoryName, deploymentId, traceId)
	if err != nil {
		return nil, err
	}

	wrappers := make([]*models.ValueWrapper, 0)

	wrappedDeploymentFile, err := models.ToValueWrapperFromNode(deployment, models.DEPLOYMENT)
	if err != nil {
		return nil, err
	}
	wrappers = append(wrappers, wrappedDeploymentFile)

	for _, f := range deployment.DeploymentFolderFiles() {
		wrappedFile, err := models.ToValueWrapperFromNode(f, models.DEPLOYMENT)
		if err != nil {
			return nil, err
		}
		wrappers = append(wrappers, wrappedFile)
	}

	return wrappers, nil
}
