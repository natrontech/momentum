package momentumconfig

import (
	utils "momentum/momentum-core/momentum-utils"
	"os"
)

type MomentumConfig struct {
	dataDir                string
	validationTmpDir       string
	templatesDir           string
	deploymentTemplatePath string
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

func (m *MomentumConfig) DeploymentTemplateDir() string {
	return m.deploymentTemplatePath
}

func InitializeMomentumCore() (*MomentumConfig, error) {

	usrHome, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	momentumDir := utils.BuildPath(usrHome, ".momentum")
	dataDir := utils.BuildPath(momentumDir, "data")
	validationTmpDir := utils.BuildPath(momentumDir, "validation")
	templatesDir := "./templates"

	createPathIfNotPresent(dataDir, momentumDir)
	createPathIfNotPresent(validationTmpDir, momentumDir)

	config := new(MomentumConfig)
	config.dataDir = dataDir
	config.validationTmpDir = validationTmpDir
	config.templatesDir = templatesDir
	config.deploymentTemplatePath = utils.BuildPath(templatesDir, "deployments")

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
