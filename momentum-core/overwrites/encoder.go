package overwrites

import (
	"momentum-core/utils"
	"os"

	"gopkg.in/yaml.v3"
)

func rulesFromFile(path string) (*OverwriteConfig, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	conf := new(OverwriteConfig)
	err = yaml.NewDecoder(f).Decode(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func rulesToFile(path string, rules *OverwriteConfig) (string, error) {

	f, err := utils.FileOpen(path, os.O_WRONLY)
	if err != nil {
		return "", err
	}
	defer f.Close()

	err = yaml.NewEncoder(f).Encode(rules)
	if err != nil {
		return "", err
	}

	return path, nil
}
