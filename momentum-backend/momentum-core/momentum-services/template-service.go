package momentumservices

import (
	utils "momentum/momentum-core/momentum-utils"
	"text/template"
)

const KUSTOMIZATION_FILE_NAME = "kustomization.yaml"

type DeploymentKustomizationTemplate struct {
	DeploymentName string
}

type DeploymentStageKustomizationTemplate struct {
	DeploymentName string
	RepositoryName string
}

type DeploymentReleaseTemplate struct {
	ApplicationName string
}

type TemplateService struct{}

func NewTemplateService() *TemplateService {
	return new(TemplateService)
}

func (ts *TemplateService) ApplyDeploymentStageKustomization(path string, deploymentName string, repositoryName string) error {
	depStageTemp := DeploymentStageKustomizationTemplate{deploymentName, repositoryName}
	return ts.ParseReplaceWrite(path, depStageTemp)
}

func (ts *TemplateService) ApplyDeploymentKustomizationTemplate(path string, deploymentName string) error {
	depTemp := DeploymentKustomizationTemplate{deploymentName}
	return ts.ParseReplaceWrite(path, depTemp)
}

func (ts *TemplateService) ApplyDeploymentReleaseTemplate(path string, applicationName string) error {
	depReleaseTemp := DeploymentReleaseTemplate{applicationName}
	return ts.ParseReplaceWrite(path, depReleaseTemp)
}

func (ts *TemplateService) ParseReplaceWrite(path string, data any) error {

	parsed, err := ts.parse(path)
	if err != nil {
		return err
	}

	err = ts.replace(path, parsed, data)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TemplateService) parse(path string) (*template.Template, error) {

	temp, err := template.ParseFiles(path)

	if err != nil {
		return nil, err
	}

	return temp, nil
}

func (ts *TemplateService) replace(path string, t *template.Template, data any) error {

	writer, err := utils.FileOpen(path, utils.FILE_ALLOW_READ_WRITE_ALL)
	if err != nil {
		return err
	}
	err = t.Execute(writer, data)
	writer.Close()
	return err
}
