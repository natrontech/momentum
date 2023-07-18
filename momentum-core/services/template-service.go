package services

import (
	"fmt"
	"momentum-core/config"
	"momentum-core/utils"
	"os"
	"strings"
	"text/template"
)

const KUSTOMIZATION_FILE_NAME = "kustomization.yaml"

type ApplicationTemplate struct {
	applicationRepositoryPath                string
	applicationNamespacePath                 string
	applicationBaseKustomizationTemplatePath string
	applicationBaseReleaseTemplatePath       string
	applicationKustomizationTemplate         *ApplicationKustomizationTemplate
	applicationReleaseTemplate               *ApplicationReleaseTemplate
}

type ApplicationKustomizationTemplate struct {
	ApplicationName string
}

type ApplicationReleaseTemplate struct {
	ApplicationName      string
	ApplicationChartName string
}

type StageTemplate struct {
	stageBaseKustomizationPath string
	stageBaseReleasePath       string
	stageKustomizationTemplate *StageKustomizationTemplate
	stageReleaseTemplate       *StageReleaseTemplate
}

type StageKustomizationTemplate struct {
	StageName string
}

type StageReleaseTemplate struct {
	StageName       string
	ApplicationName string
}

type DeploymentTemplate struct {
	deploymentKustomizationTemplatePath              string
	deploymentStageDeploymentDescriptionTemplatePath string
	deploymentReleaseTemplatePath                    string
	deploymentKustomizationTemplate                  *DeploymentKustomizationTemplate
	deploymentStageDeploymentDescriptionTemplate     *DeploymentStageDeploymentDescriptionTemplate
	deploymentReleaseTemplate                        *DeploymentReleaseTemplate
}

type DeploymentKustomizationTemplate struct {
	DeploymentNameWithoutEnding string
}

type DeploymentStageDeploymentDescriptionTemplate struct {
	DeploymentNameWithoutEnding string
	DeploymentName              string
	RepositoryName              string
}

type DeploymentReleaseTemplate struct {
	ApplicationName string
}

type TemplateService struct{}

func NewTemplateService() *TemplateService {
	return new(TemplateService)
}

func (ts *TemplateService) NewApplicationTemplate(
	applicationRepositoryPath string,
	applicationNamespacePath string,
	applicationBaseKustomizationPath string,
	applicationBaseReleasePath string,
	applicationName string,
	applicationChartName string,
) *ApplicationTemplate {

	template := new(ApplicationTemplate)

	template.applicationRepositoryPath = applicationRepositoryPath
	template.applicationNamespacePath = applicationNamespacePath
	template.applicationBaseKustomizationTemplatePath = applicationBaseKustomizationPath
	template.applicationBaseReleaseTemplatePath = applicationBaseReleasePath

	applicationKustomizationTemplate := new(ApplicationKustomizationTemplate)
	applicationKustomizationTemplate.ApplicationName = applicationName

	applicationReleaseTemplate := new(ApplicationReleaseTemplate)
	applicationReleaseTemplate.ApplicationName = applicationName
	applicationReleaseTemplate.ApplicationChartName = applicationChartName

	template.applicationKustomizationTemplate = applicationKustomizationTemplate
	template.applicationReleaseTemplate = applicationReleaseTemplate

	return template
}

func (ts *TemplateService) NewStageTemplate(
	stageBaseKustomizationPath string,
	stageBaseReleasePath string,
	stageName string,
	applicationName string,
) *StageTemplate {

	template := new(StageTemplate)

	template.stageBaseKustomizationPath = stageBaseKustomizationPath
	template.stageBaseReleasePath = stageBaseReleasePath

	stageBaseKustomization := new(StageKustomizationTemplate)
	stageBaseKustomization.StageName = stageName

	stageBaseRelease := new(StageReleaseTemplate)
	stageBaseRelease.StageName = stageName
	stageBaseRelease.ApplicationName = applicationName

	template.stageKustomizationTemplate = stageBaseKustomization
	template.stageReleaseTemplate = stageBaseRelease

	return template
}

func (ts *TemplateService) NewDeploymentTemplate(
	deploymentKustomizationTemplatePath string,
	deploymentStageDeploymentDescriptionTemplatePath string,
	deploymentReleaseTemplatePath string,
	deploymentFileName string,
	applicationName string,
	repositoryName string) *DeploymentTemplate {

	template := new(DeploymentTemplate)

	template.deploymentKustomizationTemplatePath = deploymentKustomizationTemplatePath
	template.deploymentStageDeploymentDescriptionTemplatePath = deploymentStageDeploymentDescriptionTemplatePath
	template.deploymentReleaseTemplatePath = deploymentReleaseTemplatePath

	deploymentNameWithoutEnding, _ := strings.CutSuffix(deploymentFileName, ".yml")
	deploymentNameWithoutEnding, _ = strings.CutSuffix(deploymentNameWithoutEnding, ".yaml")

	templateStageDeployment := new(DeploymentStageDeploymentDescriptionTemplate)
	templateStageDeployment.DeploymentName = deploymentFileName
	templateStageDeployment.DeploymentNameWithoutEnding = deploymentNameWithoutEnding
	templateStageDeployment.RepositoryName = repositoryName

	deploymentKustomizationTemplate := new(DeploymentKustomizationTemplate)
	deploymentKustomizationTemplate.DeploymentNameWithoutEnding = deploymentNameWithoutEnding

	deploymentReleaseTemplate := new(DeploymentReleaseTemplate)
	deploymentReleaseTemplate.ApplicationName = applicationName

	template.deploymentKustomizationTemplate = deploymentKustomizationTemplate
	template.deploymentStageDeploymentDescriptionTemplate = templateStageDeployment
	template.deploymentReleaseTemplate = deploymentReleaseTemplate

	return template
}

func (ts *TemplateService) ApplyApplicationTemplate(template *ApplicationTemplate) error {

	err := ts.ApplyApplicationKustomizationTemplate(template.applicationBaseKustomizationTemplatePath, template.applicationKustomizationTemplate)
	if err != nil {
		config.LOGGER.LogError("failed applying base kustomization template", err, "")
		return err
	}

	err = ts.ApplyApplicationReleaseTemplate(template.applicationBaseReleaseTemplatePath, template.applicationReleaseTemplate)
	if err != nil {
		config.LOGGER.LogError("failed applying base release template", err, "")
		return err
	}

	err = ts.ApplyApplicationKustomizationTemplate(template.applicationRepositoryPath, template.applicationKustomizationTemplate)
	if err != nil {
		config.LOGGER.LogError("failed applying release template", err, "")
		return err
	}

	err = ts.ApplyApplicationKustomizationTemplate(template.applicationNamespacePath, template.applicationKustomizationTemplate)
	if err != nil {
		config.LOGGER.LogError("failed applying namespace template", err, "")
		return err
	}

	return nil
}

func (ts *TemplateService) ApplyStageTemplate(template *StageTemplate) error {

	err := ts.ApplyStageKustomizationTemplate(template.stageBaseKustomizationPath, template.stageKustomizationTemplate)
	if err != nil {
		return err
	}

	err = ts.ApplyStageReleaseTemplate(template.stageBaseReleasePath, template.stageReleaseTemplate)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TemplateService) ApplyDeploymentTemplate(template *DeploymentTemplate) error {

	err := ts.ApplyDeploymentStageDeploymentDescriptionTemplate(template.deploymentStageDeploymentDescriptionTemplatePath, template.deploymentStageDeploymentDescriptionTemplate)
	if err != nil {
		return err
	}

	err = ts.ApplyDeploymentKustomizationTemplate(template.deploymentKustomizationTemplatePath, template.deploymentKustomizationTemplate)
	if err != nil {
		return err
	}

	err = ts.ApplyDeploymentReleaseTemplate(template.deploymentReleaseTemplatePath, template.deploymentReleaseTemplate)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TemplateService) ApplyApplicationKustomizationTemplate(path string, template *ApplicationKustomizationTemplate) error {
	return ts.ParseReplaceWrite(path, template)
}

func (ts *TemplateService) ApplyApplicationReleaseTemplate(path string, template *ApplicationReleaseTemplate) error {
	return ts.ParseReplaceWrite(path, template)
}

func (ts *TemplateService) ApplyStageKustomizationTemplate(path string, template *StageKustomizationTemplate) error {
	return ts.ParseReplaceWrite(path, template)
}

func (ts *TemplateService) ApplyStageReleaseTemplate(path string, template *StageReleaseTemplate) error {
	return ts.ParseReplaceWrite(path, template)
}

func (ts *TemplateService) ApplyDeploymentStageDeploymentDescriptionTemplate(path string, template *DeploymentStageDeploymentDescriptionTemplate) error {
	return ts.ParseReplaceWrite(path, template)
}

func (ts *TemplateService) ApplyDeploymentKustomizationTemplate(path string, template *DeploymentKustomizationTemplate) error {
	return ts.ParseReplaceWrite(path, template)
}

func (ts *TemplateService) ApplyDeploymentReleaseTemplate(path string, template *DeploymentReleaseTemplate) error {
	return ts.ParseReplaceWrite(path, template)
}

func (ts *TemplateService) ParseReplaceWrite(path string, data any) error {

	parsed, err := ts.parse(path)
	if err != nil {
		fmt.Println("failed parsing template:", err.Error())
		return err
	}

	err = ts.replace(path, parsed, data)
	if err != nil {
		fmt.Println("failed replacing template strings:", err.Error())
		return err
	}

	return nil
}

func (ts *TemplateService) parse(path string) (*template.Template, error) {

	reader, err := utils.FileOpen(path, utils.FILE_ALLOW_READ_WRITE_ALL)
	if err != nil {
		fmt.Println("failed opening template:", err.Error())
		return nil, err
	}
	defer reader.Close()
	templateAsString := utils.FileAsString(reader)

	temp, err := template.New(path).Parse(templateAsString)
	if err != nil {
		fmt.Println("failed parsing path:", err.Error())
		return nil, err
	}

	return temp, nil
}

func (ts *TemplateService) replace(path string, t *template.Template, data any) error {

	utils.FileDelete(path)

	writer, err := utils.FileOpen(path, utils.FILE_ALLOW_READ_WRITE_ALL|os.O_CREATE)
	if err != nil {
		fmt.Println("failed opening template:", err.Error())
		return err
	}
	defer writer.Close()
	err = t.Execute(writer, data)
	return err
}