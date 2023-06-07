package momentumconfig

import (
	utils "momentum/momentum-core/momentum-utils"
	"os"
)

type MomentumConfig struct {
	dataDir string
}

func (m *MomentumConfig) DataDir() string {
	return m.dataDir
}

func InitializeMomentumCore() (*MomentumConfig, error) {

	usrHome, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	momentumDir := utils.BuildPath(usrHome, ".momentum")
	dataDir := utils.BuildPath(momentumDir, "data")

	if !utils.FileExists(dataDir) {
		if !utils.FileExists(momentumDir) {
			utils.DirCreate(momentumDir)
		}
		utils.DirCreate(dataDir)
	}

	config := new(MomentumConfig)
	config.dataDir = dataDir

	return config, err
}
