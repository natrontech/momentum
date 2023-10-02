package config

import (
	"errors"
	"momentum-core/utils"
	"os"
	"path/filepath"
	"strings"

	git "gopkg.in/src-d/go-git.v4"

	gittransaction "github.com/Joel-Haeberli/git-transaction"
)

const MOMENTUM_ROOT = "momentum-root"

const TRANSACTION_MODE = gittransaction.DEBUG

const MOMENTUM_GIT_USER = "MOMENTUM_GIT_USER"         // the git username associated with the repositories token
const MOMENTUM_GIT_EMAIL = "MOMENTUM_GIT_EMAIL"       // the email belonging to the user associated with the repository token
const MOMENTUM_GIT_TOKEN = "MOMENTUM_GIT_TOKEN"       // the access token associated with the repository token
const MOMENTUM_GIT_REPO_URL = "MOMENTUM_GIT_REPO_URL" // the HTTP url to the git repository the instance is working on

var GLOBAL *MomentumConfig = nil // set on initialization, otherwise crash

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

	transactionToken *gittransaction.Token

	applicationTemplateFolderPath string
	stageTemplateFolderPath       string
	deploymentTemplateFilePath    string
	deploymentTemplateFolderPath  string
}

func (m *MomentumConfig) DataDir() string {
	return m.dataDir
}

func (m *MomentumConfig) RepoDir() string {
	return filepath.Join(m.dataDir, "repository")
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

func (m *MomentumConfig) TransactionToken() *gittransaction.Token {
	return m.transactionToken
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

func (m *MomentumConfig) initializeRepository() error {

	_, err := os.Stat(m.RepoDir())
	if !os.IsNotExist(err) {
		LOGGER.LogInfo("will not clone repository because one present", "STARTUP")
		return nil
	}

	repoUrl := os.Getenv(MOMENTUM_GIT_REPO_URL)
	if repoUrl == "" {
		return errors.New("failed initializing momentum because " + MOMENTUM_GIT_REPO_URL + " was not set")
	}

	cloneRepoTo(repoUrl, "", "", m.RepoDir())

	return nil
}

func (m *MomentumConfig) initializeGitAccessToken() error {

	if TRANSACTION_MODE == gittransaction.DEBUG {
		return nil
	}

	m.transactionToken = new(gittransaction.Token)

	m.transactionToken.Username = os.Getenv(MOMENTUM_GIT_USER)
	m.transactionToken.Email = os.Getenv(MOMENTUM_GIT_EMAIL)
	m.transactionToken.Token = os.Getenv(MOMENTUM_GIT_TOKEN)

	if m.transactionToken == nil ||
		m.transactionToken.Username == "" ||
		m.transactionToken.Email == "" ||
		m.transactionToken.Token == "" {

		return errors.New("failed initializing git transaction user please make sure to set " + MOMENTUM_GIT_USER + " and " + MOMENTUM_GIT_EMAIL + " and " + MOMENTUM_GIT_TOKEN)
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

	err = config.initializeGitAccessToken()
	if err != nil {
		return nil, err
	}

	err = config.checkMandatoryTemplates()
	if err != nil {
		return nil, err
	}

	LOGGER, err = NewLogger(config.LogDir())
	if err != nil {
		panic("failed initializing logger: " + err.Error())
	}

	GLOBAL = config

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

func IdGenerationPath(path string) string {
	relevantForId, _ := strings.CutPrefix(path, GLOBAL.RepoDir())
	return relevantForId
}

func cloneRepoTo(url string, username string, password string, location string) error {
	_, err := git.PlainClone(location, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		// Auth: ..., TODO in case not public dir
	})

	return err
}
