package services

import (
	"errors"
	"momentum-core/config"
	"momentum-core/tree"
	"momentum-core/utils"
)

type TreeService struct {
	config *config.MomentumConfig
}

func NewTreeService(config *config.MomentumConfig) *TreeService {

	service := new(TreeService)

	service.config = config

	return service
}

func (ts *TreeService) repository(repositoryName string, traceId string) (*tree.Node, error) {

	repoPath := utils.BuildPath(ts.config.DataDir(), repositoryName)
	root, err := tree.Parse(repoPath)
	if err != nil {
		config.LOGGER.LogWarning("failed adding stage to resources", err, traceId)
		return nil, err
	}
	return root.Repo(), nil
}

func (ts *TreeService) application(repositoryName string, applicationId string, traceId string) (*tree.Node, error) {

	repo, err := ts.repository(repositoryName, traceId)
	if err != nil {
		return nil, err
	}

	apps := repo.Apps()
	for _, app := range apps {
		if app.Id == applicationId {
			return app, nil
		}
	}
	return nil, errors.New("no application with id " + applicationId)
}

func (ts *TreeService) stage(repositoryName string, stageId string, traceId string) (*tree.Node, error) {

	repo, err := ts.repository(repositoryName, traceId)
	if err != nil {
		return nil, err
	}

	stages := repo.AllStages()
	for _, stage := range stages {
		if stage.Id == stageId {
			return stage, nil
		}
	}
	return nil, errors.New("no stage with id " + stageId)
}

func (ts *TreeService) deployment(repositoryName string, deploymentId string, traceId string) (*tree.Node, error) {

	repo, err := ts.repository(repositoryName, traceId)
	if err != nil {
		return nil, err
	}

	deployments := repo.AllDeployments()
	for _, deployment := range deployments {
		if deployment.Id == deploymentId {
			return deployment, nil
		}
	}
	return nil, errors.New("no deployment with id " + deploymentId)
}

func (ts *TreeService) value(repositoryName string, valueId string, traceId string) (*tree.Node, error) {

	repo, err := ts.repository(repositoryName, traceId)
	if err != nil {
		return nil, err
	}

	values := repo.AllValues()
	for _, value := range values {
		if value.Id == valueId {
			return value, nil
		}
	}
	return nil, errors.New("no value with id " + valueId)
}

func (ts *TreeService) find(traceId string, repositoryName string, terms ...string) (*tree.Node, error) {

	repo, err := ts.repository(repositoryName, traceId)
	if err != nil {
		return nil, err
	}

	searchTerm := tree.ToMatchableSearchTerm(utils.BuildPath(terms...))

	config.LOGGER.LogTrace("searching for: "+searchTerm, traceId)

	result, _ := repo.FindFirst(searchTerm)

	return result, nil
}

func (ts *TreeService) findAll(traceId string, repositoryName string, terms ...string) ([]*tree.Node, error) {

	repo, err := ts.repository(repositoryName, traceId)
	if err != nil {
		return nil, err
	}

	searchTerm := tree.ToMatchableSearchTerm(utils.BuildPath(terms...))

	result := repo.Search(searchTerm)

	return result, nil
}
