package templates

import (
	"errors"
	"io/fs"
	"momentum-core/config"
	"momentum-core/files"
	"momentum-core/utils"
	"os"
	"path/filepath"
	"strings"
)

const TEMPLATE_CONFIG_FILENAME = "momentum-template-config.yaml"

func LoadSpec(templateName string, templateKind TemplateKind) []*TemplateSpec {

	return make([]*TemplateSpec, 0)
}

func LoadTemplate(templateName string) (*Template, error) {

	if !TemplateExists(templateName) {
		return nil, errors.New("no template name '" + templateName + "'")
	}

	//files.

	template := new(Template)
	//template.Kind =

	return template, nil
}

func CreateTemplate(templateRequest *CreateTemplateRequest) (*Template, error) {

	templateAnchorPath := filepath.Join(templatePathForKind(templateRequest.TemplateKind))

	err := createTemplateDir(templateAnchorPath, templateRequest.Template)
	if err != nil {
		return nil, err
	}

	_, err = templateConfigToFile(filepath.Join(templateAnchorPath, templateRequest.Template.Name, TEMPLATE_CONFIG_FILENAME), templateRequest.TemplateConfig)
	if err != nil {
		return nil, err
	}

	return LoadTemplate(templateRequest.Template.Name)
}

func createTemplateDir(anchorPath string, templateDir *TemplateDir) error {

	templateAnchorPath := filepath.Join(anchorPath, templateDir.Name)
	err := utils.DirCreate(templateAnchorPath)
	if err != nil {
		return err
	}

	for _, dir := range templateDir.Directories {
		err := createTemplateDir(templateAnchorPath, dir)
		if err != nil {
			return err
		}
	}

	for _, file := range templateDir.Files {
		err := createTemplateFile(templateAnchorPath, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func createTemplateFile(anchorPath string, templateFile *TemplateFile) error {

	if !strings.HasSuffix(templateFile.Name, ".yaml") && !strings.HasSuffix(templateFile.Name, ".yml") {
		return errors.New("only yaml files supported at the moment (you sent " + templateFile.Name + ")")
	}

	path := filepath.Join(anchorPath, templateFile.Name)
	body, err := files.FileToRaw(templateFile.TemplateBody)
	if err != nil {
		return err
	}

	success := utils.FileWrite(path, body)
	if !success {
		return errors.New("unable to write file '" + path + "'")
	}
	return nil
}

func templatePathForKind(kind TemplateKind) string {

	switch kind {
	case APPLICATION:
		return config.ApplicationTemplatesPath(config.GLOBAL)
	case STAGE:
		return config.StageTemplatesPath(config.GLOBAL)
	case DEPLOYMENT:
		return config.DeploymentTemplatesPath(config.GLOBAL)
	default:
		return ""
	}
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

// validates a given config. A config is valid if all contained children are present.
func validTemplateConfig(templateConfig *TemplateConfig, templateKind TemplateKind) (bool, error) {

	if templateKind != templateConfig.Kind {
		return false, errors.New("template config must have same kind as request")
	}

	for _, temp := range templateConfig.Children {

		if !TemplateExists(temp) {
			return false, errors.New("template with name '" + temp + "' does not exist")
		}
	}

	return true, nil
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
