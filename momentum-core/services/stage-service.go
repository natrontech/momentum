package services

import (
	"errors"
	"momentum-core/config"
	"momentum-core/models"
	"momentum-core/tree"
	"momentum-core/utils"
)

type StageService struct {
	config          *config.MomentumConfig
	treeService     *TreeService
	templateService *TemplateService
}

func NewStageService(config *config.MomentumConfig, treeService *TreeService, templateService *TemplateService) *StageService {

	stageService := new(StageService)

	stageService.config = config
	stageService.treeService = treeService
	stageService.templateService = templateService

	return stageService
}

func (as *StageService) GetStages(repositoryName string, traceId string) ([]*models.Stage, error) {

	repo, err := as.treeService.repository(repositoryName, traceId)
	if err != nil {
		config.LOGGER.LogWarning("failed adding deployment to resources", err, traceId)
		return nil, err
	}

	stages := repo.AllStages()

	mappedStages := make([]*models.Stage, 0)
	for _, stage := range stages {
		mapped, err := models.ToStageFromNode(stage, stage.Parent.Id)
		if err != nil {
			return nil, err
		}
		mappedStages = append(mappedStages, mapped)
	}

	return mappedStages, nil
}

func (as *StageService) GetStage(repositoryName string, stageId string, traceId string) (*models.Stage, error) {

	result, err := as.treeService.stage(repositoryName, stageId, traceId)
	if err != nil {
		config.LOGGER.LogWarning("unable to find stage "+stageId, err, traceId)
		return nil, err
	}

	return models.ToStageFromNode(result, result.Parent.Id)
}

func (s *StageService) FindStageById(stageId string, repositoryName string, traceId string) (*models.Stage, error) {

	stageNode, err := s.treeService.stage(repositoryName, stageId, traceId)
	if err != nil {
		config.LOGGER.LogWarning("unable to find stage "+stageId, err, traceId)
		return nil, err
	}

	return models.ToStageFromNode(stageNode, stageNode.Parent.Id)
}

func (s *StageService) AddStage(request *models.StageCreateRequest, traceId string) (*models.Stage, error) {

	parentAppNode, err := s.treeService.application(request.RepositoryName, request.ParentApplicationId, traceId)
	if err != nil {
		config.LOGGER.LogWarning("unable to find application "+request.ParentApplicationId, err, traceId)
		return nil, err
	}

	var stageMountNode *tree.Node
	if request.ParentStageId == "" {
		stageMountNode = parentAppNode
	} else {
		stageMountNode, err = s.treeService.stage(request.RepositoryName, request.ParentStageId, traceId)
		if err != nil {
			return nil, err
		}
	}

	dirs := stageMountNode.Directories()
	for _, dir := range dirs {
		if dir.Path == request.Name {
			return nil, errors.New("stage with this name already exists")
		}
	}

	stageMountPath := utils.BuildPath(stageMountNode.FullPath(), request.Name)
	stageBasePath := utils.BuildPath(stageMountPath, "_base")

	_, err = utils.DirCopy(s.config.StageTemplateFolderPath(), stageMountPath)
	if err != nil {
		return nil, err
	}

	stageBaseKustomizationPath := utils.BuildPath(stageBasePath, KUSTOMIZATION_FILE_NAME)
	stageBaseReleasePath := utils.BuildPath(stageBasePath, "release.yaml")

	template := s.templateService.NewStageTemplate(stageBaseKustomizationPath, stageBaseReleasePath, request.Name, parentAppNode.Path)
	err = s.templateService.ApplyStageTemplate(template)
	if err != nil {
		return nil, err
	}

	parentKustomizationResources, err := s.treeService.find(traceId, request.RepositoryName, stageMountNode.FullPath(), KUSTOMIZATION_FILE_NAME, "resources")
	if err != nil {
		config.LOGGER.LogError(err.Error(), err, traceId)
		return nil, err
	}

	if parentKustomizationResources == nil {
		config.LOGGER.LogTrace("unable to find kustomization resources for parent stage or application of new stage", traceId)
		kustomizationFileNode, err := s.treeService.find(traceId, request.RepositoryName, stageMountNode.FullPath(), KUSTOMIZATION_FILE_NAME)
		if err != nil {
			config.LOGGER.LogError(err.Error(), err, traceId)
			return nil, err
		}

		err = kustomizationFileNode.AddSequence("resources", []string{request.Name}, 0)
		if err != nil {
			config.LOGGER.LogError(err.Error(), err, traceId)
			return nil, err
		}

		err = kustomizationFileNode.Write(true)
		if err != nil {
			config.LOGGER.LogWarning("failed writing stage to resources", err, traceId)
			return nil, err
		}
	} else {
		err = parentKustomizationResources.AddValue(request.Name, 0)
		if err != nil {
			config.LOGGER.LogWarning("failed adding stage to resources", err, traceId)
			return nil, err
		}

		err = parentKustomizationResources.Write(true)
		if err != nil {
			config.LOGGER.LogWarning("failed writing stage to resources", err, traceId)
			return nil, err
		}
	}

	stage, err := s.treeService.find(traceId, request.RepositoryName, stageMountPath)
	if err != nil {
		config.LOGGER.LogWarning("unable to find "+stageMountPath, err, traceId)
		return nil, err
	}

	stg, err := models.ToStageFromNode(stage, stageMountNode.Id)
	if err != nil {
		config.LOGGER.LogWarning("failed mapping stage from node", err, traceId)
		return nil, err
	}
	return stg, nil
}
