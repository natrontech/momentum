package templates

import (
	"io/fs"
	"momentum-core/config"
	"momentum-core/utils"
	"os"
	"strings"
)

func LoadSpec(templateName string, templateKind TemplateKind) []*TemplateSpec {

	return make([]*TemplateSpec, 0)
}

func CreateTemplate(templateRequest *CreateTemplateRequest) {

}

func TemplateNames(path string) []string {

	entrs, err := entries(path)
	if err != nil {
		return make([]string, 0)
	}

	names := make([]string, 0)
	for _, entry := range entrs {
		names = append(names, entry.Name())
	}

	return names
}

func TemplateExists(name string) bool {

	return applicationExists(name) || stageExists(name) || deploymentExists(name)
}

func applicationExists(name string) bool {

	appTemplates := config.ApplicationTemplatesPath(config.GLOBAL)
	entries, err := entries(appTemplates)
	if err != nil {
		return false
	}
	return directoryContains(name, entries)
}

func stageExists(name string) bool {

	stageTemplates := config.ApplicationTemplatesPath(config.GLOBAL)
	entries, err := entries(stageTemplates)
	if err != nil {
		return false
	}
	return directoryContains(name, entries)
}

func deploymentExists(name string) bool {

	deploymentTemplates := config.ApplicationTemplatesPath(config.GLOBAL)
	entries, err := entries(deploymentTemplates)
	if err != nil {
		return false
	}
	return directoryContains(name, entries)
}

func entries(dirPath string) ([]fs.DirEntry, error) {

	dir, err := utils.FileOpen(dirPath, int(os.ModeDir.Perm()))
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	return dir.ReadDir(-1) // reads all entries
}

func directoryContains(name string, entries []fs.DirEntry) bool {

	for _, entry := range entries {

		if strings.EqualFold(entry.Name(), name) {
			return true
		}
	}

	return false
}
