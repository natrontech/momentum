package momentumconfig

import (
	utils "momentum/momentum-core/momentum-utils"
	"os"
)

type MomentumConfig struct {
	dataDir          string
	validationTmpDir string
}

func (m *MomentumConfig) DataDir() string {
	return m.dataDir
}

func (m *MomentumConfig) ValidationTmpDir() string {
	return m.validationTmpDir
}

func InitializeMomentumCore() (*MomentumConfig, error) {

	usrHome, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	momentumDir := utils.BuildPath(usrHome, ".momentum")
	dataDir := utils.BuildPath(momentumDir, "data")
	validationTmpDir := utils.BuildPath(momentumDir, "validation")

	createPathIfNotPresent(dataDir, momentumDir)
	createPathIfNotPresent(validationTmpDir, momentumDir)

	config := new(MomentumConfig)
	config.dataDir = dataDir
	config.validationTmpDir = validationTmpDir

	return config, err
}

func createPathIfNotPresent(path string, momentumDir string) {

	if !utils.FileExists(path) {
		if !utils.FileExists(momentumDir) {
			utils.DirCreate(momentumDir)
		}
		utils.DirCreate(path)
	}
}
