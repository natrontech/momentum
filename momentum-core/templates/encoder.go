package templates

import (
	"momentum-core/utils"
	"os"

	"gopkg.in/yaml.v3"
)

func templateConfigFromFile(path string) (*TemplateConfig, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	conf := new(TemplateConfig)
	err = yaml.NewDecoder(f).Decode(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func templateConfigToFile(path string, rules *TemplateConfig) (string, error) {

	f, err := utils.FileOpen(path, os.O_CREATE|os.O_WRONLY)
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
