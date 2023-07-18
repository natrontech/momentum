package config

import (
	"errors"
	"momentum-core/utils"
	"os"
)

const MOMENTUM_ROOT = "momentum-root"

type ILoggerClient interface {
	LogTrace(msg string, traceId string)
	LogInfo(msg string, traceId string)
	LogWarning(msg string, err error, traceId string)
	LogError(msg string, err error, traceId string)
	TraceId() string
}

var LOGGER ILoggerClient

type MomentumConfig struct {
	dataDir          string
	validationTmpDir string
	templatesDir     string
	logDir           string

	applicationTemplateFolderPath string
	stageTemplateFolderPath       string
	deploymentTemplateFilePath    string
	deploymentTemplateFolderPath  string
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

func (m *MomentumConfig) LogDir() string {
	return m.logDir
}

func (m *MomentumConfig) ApplicationTemplateFolderPath() string {
	return m.applicationTemplateFolderPath
}

func (m *MomentumConfig) StageTemplateFolderPath() string {
	return m.stageTemplateFolderPath
}

func (m *MomentumConfig) DeploymentTemplateFolderPath() string {
	return m.deploymentTemplateFolderPath
}

func (m *MomentumConfig) DeploymentTemplateFilePath() string {
	return m.deploymentTemplateFilePath
}

func (m *MomentumConfig) checkMandatoryTemplates() error {

	if !utils.FileExists(m.ApplicationTemplateFolderPath()) {
		return errors.New("provide mandatory template for application folder at " + m.TemplateDir())
	}

	if !utils.FileExists(m.StageTemplateFolderPath()) {
		return errors.New("provide mandatory template for stage folder at " + m.TemplateDir())
	}

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
	logDir := momentumDir

	createPathIfNotPresent(dataDir, momentumDir)
	createPathIfNotPresent(validationTmpDir, momentumDir)

	config := new(MomentumConfig)

	config.logDir = logDir
	config.dataDir = dataDir
	config.validationTmpDir = validationTmpDir
	config.templatesDir = templatesDir
	config.applicationTemplateFolderPath = utils.BuildPath(templatesDir, "applications")
	config.stageTemplateFolderPath = utils.BuildPath(templatesDir, "stages")
	config.deploymentTemplateFolderPath = utils.BuildPath(templatesDir, "deployments", "deploymentName")
	config.deploymentTemplateFilePath = utils.BuildPath(templatesDir, "deployments", "deploymentName.yaml")

	err = config.checkMandatoryTemplates()
	if err != nil {
		return nil, err
	}

	LOGGER, err = NewLogger(config.LogDir())
	if err != nil {
		panic("failed initializing logger: " + err.Error())
	}

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