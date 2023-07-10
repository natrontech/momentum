package momentumservices

import (
	"fmt"
	utils "momentum/momentum-core/momentum-utils"
	"os"
	"strings"
	"text/template"
)

const KUSTOMIZATION_FILE_NAME = "kustomization.yaml"

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

func (ts *TemplateService) ApplyDeploymentStageDeploymentDescriptionTemplate(path string, deploymentFileName string, repositoryName string) error {
	deploymentNameWithoutEnding, _ := strings.CutSuffix(deploymentFileName, ".yml")
	deploymentNameWithoutEnding, _ = strings.CutSuffix(deploymentNameWithoutEnding, ".yaml")
	depStageTemp := DeploymentStageDeploymentDescriptionTemplate{deploymentNameWithoutEnding, deploymentFileName, repositoryName}
	return ts.ParseReplaceWrite(path, depStageTemp)
}

func (ts *TemplateService) ApplyDeploymentKustomizationTemplate(path string, deploymentFileName string) error {
	deploymentNameWithoutEnding, _ := strings.CutSuffix(deploymentFileName, ".yml")
	deploymentNameWithoutEnding, _ = strings.CutSuffix(deploymentNameWithoutEnding, ".yaml")
	depTemp := DeploymentKustomizationTemplate{deploymentNameWithoutEnding}
	return ts.ParseReplaceWrite(path, depTemp)
}

func (ts *TemplateService) ApplyDeploymentReleaseTemplate(path string, applicationName string) error {
	depReleaseTemp := DeploymentReleaseTemplate{applicationName}
	return ts.ParseReplaceWrite(path, depReleaseTemp)
}

func (ts *TemplateService) ParseReplaceWrite(path string, data any) error {

	parsed, err := ts.parse(path)
	if err != nil {
		fmt.Println("failed parsing template:", err.Error())
		return err
	}

	fmt.Println("template parsed", path)

	err = ts.replace(path, parsed, data)
	if err != nil {
		fmt.Println("failed replacing template strings:", err.Error())
		return err
	}

	fmt.Println("template replaced", path)

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
	fmt.Println("executing template", path)
	err = t.Execute(writer, data)
	return err
}
