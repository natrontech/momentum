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
	logDir           string

	transactionToken *gittransaction.Token
}

func (m *MomentumConfig) DataDir() string {
	return m.dataDir
}

func (m *MomentumConfig) RepoDir() string {
	return filepath.Join(m.DataDir(), "repository")
}

func (m *MomentumConfig) ValidationTmpDir() string {
	return m.validationTmpDir
}

func (m *MomentumConfig) LogDir() string {
	return m.logDir
}

func TemplateDir(config *MomentumConfig) string {
	return filepath.Join(config.RepoDir(), "templates")
}

func ApplicationTemplatesPath(config *MomentumConfig) string {
	return filepath.Join(TemplateDir(config), "applications")
}

func StageTemplatesPath(config *MomentumConfig) string {
	return filepath.Join(TemplateDir(config), "stages")
}

func DeploymentTemplatesPath(config *MomentumConfig) string {
	return filepath.Join(TemplateDir(config), "deployments")
}

func (m *MomentumConfig) TransactionToken() *gittransaction.Token {
	return m.transactionToken
}

func checkMandatoryTemplates(config *MomentumConfig) error {

	errs := make([]error, 0)

	templatePath := TemplateDir(config)
	if !utils.FileExists(templatePath) {
		err := utils.DirCreate(templatePath)
		if err != nil {
			errs = append(errs, err)
		}
	}

	appTemplatePath := ApplicationTemplatesPath(config)
	if !utils.FileExists(appTemplatePath) {
		err := utils.DirCreate(appTemplatePath)
		if err != nil {
			errs = append(errs, err)
		}
	}

	stageTemplatePath := StageTemplatesPath(config)
	if !utils.FileExists(stageTemplatePath) {
		err := utils.DirCreate(stageTemplatePath)
		if err != nil {
			errs = append(errs, err)
		}
	}

	deploymentTemplatePath := DeploymentTemplatesPath(config)
	if !utils.FileExists(deploymentTemplatePath) {
		err := utils.DirCreate(deploymentTemplatePath)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func initializeRepository(config *MomentumConfig) error {

	_, err := os.Stat(config.RepoDir())
	if !os.IsNotExist(err) {
		LOGGER.LogInfo("will not clone repository because one present", "STARTUP")
		return nil
	}

	repoUrl := os.Getenv(MOMENTUM_GIT_REPO_URL)
	if repoUrl == "" {
		return errors.New("failed initializing momentum because " + MOMENTUM_GIT_REPO_URL + " was not set")
	}

	cloneRepoTo(repoUrl, "", "", config.RepoDir())

	if !utils.FileExists(filepath.Join(config.RepoDir(), MOMENTUM_ROOT)) {
		return errors.New("invalid momentum repository")
	}

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
	logDir := momentumDir

	createPathIfNotPresent(dataDir, momentumDir)
	createPathIfNotPresent(validationTmpDir, momentumDir)

	config := new(MomentumConfig)

	config.logDir = logDir
	config.dataDir = dataDir
	config.validationTmpDir = validationTmpDir

	LOGGER, err = NewLogger(config.LogDir())
	if err != nil {
		panic("failed initializing logger: " + err.Error())
	}

	err = config.initializeGitAccessToken()
	if err != nil {
		return nil, err
	}

	err = initializeRepository(config)
	if err != nil {
		return nil, err
	}

	err = checkMandatoryTemplates(config)
	if err != nil {
		return nil, err
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
