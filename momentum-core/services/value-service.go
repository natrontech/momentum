package services

import (
	"errors"
	"momentum-core/config"
	"momentum-core/models"
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

func (vs *ValueService) ValuesByRepository(repositoryName string, traceId string) ([]*models.ValueWrapper, error) {
	return make([]*models.ValueWrapper, 0), nil
}

func (vs *ValueService) ValuesByApplication(repositoryName string, applicationId string, traceId string) ([]*models.ValueWrapper, error) {
	return make([]*models.ValueWrapper, 0), nil
}

func (vs *ValueService) ValuesByStage(repositoryName string, stageId string, traceId string) ([]*models.ValueWrapper, error) {
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

// This method removes all parents which shall not be read in certain types.
// For example, when reading values of a stage, we do not want to read the
// value of each deployment.
func (vs *ValueService) cleanParents(parents []string, reading models.ValueType) []string {

	switch reading {

	case models.REPOSITORY:

	case models.APPLICATION:

	case models.STAGE:

	case models.DEPLOYMENT:

	}

	return make([]string, 0)
}
