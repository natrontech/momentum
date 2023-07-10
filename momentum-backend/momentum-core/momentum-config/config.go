package momentumconfig

import (
	"errors"
	utils "momentum/momentum-core/momentum-utils"
	"os"
)

type MomentumConfig struct {
	dataDir                      string
	validationTmpDir             string
	templatesDir                 string
	deploymentTemplateFilePath   string
	deploymentTemplateFolderPath string
}

func (m *MomentumConfig) DataDir() string {
	return m.dataDir
}

func (m *MomentumConfig) ValidationTmpDir() string {
	return m.validationTmpDir
}

func (m *MomentumConfig) TemplateDir() string {
	return m.templatesDir
}

func (m *MomentumConfig) DeploymentTemplateFolderPath() string {
	return m.deploymentTemplateFolderPath
}

func (m *MomentumConfig) DeploymentTemplateFilePath() string {
	return m.deploymentTemplateFilePath
}

func (m *MomentumConfig) checkMandatoryTemplates() error {

	if !utils.FileExists(m.DeploymentTemplateFolderPath()) {
		return errors.New("provide mandatory template for deployment folders at " + m.DeploymentTemplateFolderPath())
	}

	if !utils.FileExists(m.DeploymentTemplateFilePath()) {
		return errors.New("provide mandatory template for deployment files at " + m.DeploymentTemplateFilePath())
	}

	return nil
}

func InitializeMomentumCore() (*MomentumConfig, error) {

	usrHome, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	momentumDir := utils.BuildPath(usrHome, ".momentum")
	dataDir := utils.BuildPath(momentumDir, "data")
	validationTmpDir := utils.BuildPath(momentumDir, "validation")
	templatesDir := utils.BuildPath(momentumDir, "templates")

	createPathIfNotPresent(dataDir, momentumDir)
	createPathIfNotPresent(validationTmpDir, momentumDir)

	config := new(MomentumConfig)

	config.dataDir = dataDir

	config.validationTmpDir = validationTmpDir

	config.templatesDir = templatesDir
	config.deploymentTemplateFolderPath = utils.BuildPath(templatesDir, "deployments", "deploymentName")
	config.deploymentTemplateFilePath = utils.BuildPath(templatesDir, "deployments", "deploymentName.yaml")

	err = config.checkMandatoryTemplates()

	return config, err
}

func createPathIfNotPresent(path string, parentDir string) {

	if !utils.FileExists(path) {
		if !utils.FileExists(parentDir) {
			utils.DirCreate(parentDir)
		}
		utils.DirCreate(path)
	}
}
