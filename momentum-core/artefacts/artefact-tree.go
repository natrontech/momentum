package artefacts

import (
	"errors"
	"momentum-core/config"
	"momentum-core/utils"
	"os"
	"path/filepath"
	"strings"
)

const META_PREFIX = "_"
const DEPLOYMENT_WRAP_FOLDER_NAME = "_deploy"

const LOAD_ENTIRE_TREE = -1

func LoadArtefactTree() (*Artefact, error) {

	return LoadArtefactTreeFrom(filepath.Join(config.GLOBAL.RepoDir(), config.MOMENTUM_ROOT))
}

func LoadArtefactTreeFrom(root string) (*Artefact, error) {

	return loadArtefactTreeUntilDepth(root, nil, LOAD_ENTIRE_TREE)
}

func WriteToString(artefact *Artefact) string {

	return writeToString(artefact, 0, "")
}

func Root(artefact *Artefact) *Artefact {

	current := artefact
	for current.Parent != nil {
		current = current.Parent
	}

	return current
}

func NewArtefact(path string, parent *Artefact) (*Artefact, error) {

	f, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, errors.New("path " + path + " does not exist")
	}

	artefact := new(Artefact)
	artefact.Id, err = utils.GenerateId(config.IdGenerationPath(path))
	if err != nil {
		return nil, err
	}

	artefact.Name = f.Name()
	artefact.Content = make([]*Artefact, 0)

	if parent != nil {
		artefact.Parent = parent
		artefact.ParentId = artefact.Parent.Id
		artefact.Parent.Content = append(artefact.Parent.Content, artefact)
		artefact.Parent.ContentIds = append(artefact.Parent.ContentIds, artefact.Id)
	}

	if artefact.Parent == nil {
		artefact.ArtefactType = ROOT
	} else if f.Mode().IsRegular() {
		artefact.ArtefactType = FILE
	} else if strings.HasPrefix(f.Name(), META_PREFIX) {
		artefact.ArtefactType = META
	} else if IsApplication(artefact) {
		artefact.ArtefactType = APPLICATION
	} else if IsDeployment(artefact) {
		artefact.ArtefactType = DEPLOYMENT
	} else if IsStage(artefact) {
		artefact.ArtefactType = STAGE
	} else {
		return nil, errors.New("unable to decide artefact type")
	}

	return artefact, nil
}

func Applications() []*Artefact {

	t, err := LoadArtefactTree()
	if err != nil {
		return make([]*Artefact, 0)
	}

	applications := make([]*Artefact, 0)
	for _, chld := range t.Content {
		if chld.ArtefactType == APPLICATION {
			applications = append(applications, chld)
		}
	}

	return applications
}

func StagesByApplication(applicationOrStageId string) []*Artefact {

	t, err := LoadArtefactTree()
	if err != nil {
		return make([]*Artefact, 0)
	}

	artefacts := FlatPreorder(t, make([]*Artefact, 0))
	stages := make([]*Artefact, 0)
	for _, artefact := range artefacts {
		if artefact.ArtefactType == STAGE && artefact.Parent.Id == applicationOrStageId {
			stages = append(stages, artefact)
		}
	}

	return stages
}

func DeploymentsByStage(stageId string) []*Artefact {

	t, err := LoadArtefactTree()
	if err != nil {
		return make([]*Artefact, 0)
	}

	artefacts := FlatPreorder(t, make([]*Artefact, 0))
	deployments := make([]*Artefact, 0)
	for _, artefact := range artefacts {
		if artefact.ArtefactType == DEPLOYMENT && artefact.Parent.Id == stageId {
			deployments = append(deployments, artefact)
		}
	}

	return deployments
}

func FileById(id string) (*Artefact, error) {

	artefacts, err := LoadArtefactTree() // origin

	if err != nil {
		return nil, err
	}

	files := Files(artefacts)
	for _, file := range files {
		if file.Id == id {
			return file, nil
		}
	}

	return nil, errors.New("no file with id " + id)
}

func DirectoryById(id string) (*Artefact, error) {

	artefacts, err := LoadArtefactTree() // origin

	if err != nil {
		return nil, err
	}

	directories := Directories(artefacts)
	for _, dir := range directories {
		if dir.Id == id {
			return dir, nil
		}
	}

	return nil, errors.New("no file with id " + id)
}

func Directories(origin *Artefact) []*Artefact {

	flattened := FlatPreorder(origin, make([]*Artefact, 0))
	directories := make([]*Artefact, 0)
	for _, artefact := range flattened {
		if artefact.ArtefactType != FILE && artefact.ArtefactType != META {
			directories = append(directories, artefact)
		}
	}

	return directories
}

func Files(origin *Artefact) []*Artefact {

	flattened := FlatPreorder(origin, make([]*Artefact, 0))
	files := make([]*Artefact, 0)
	for _, artefact := range flattened {
		if artefact.ArtefactType == FILE {
			files = append(files, artefact)
		}
	}

	return files
}

func IsApplication(artefact *Artefact) bool {

	if artefact != nil && artefact.Parent != nil {
		if artefact.Parent.Name == config.MOMENTUM_ROOT {
			return true
		}
	}

	return false
}

func IsStage(artefact *Artefact) bool {

	if artefact.Parent != nil &&
		artefact.ArtefactType != FILE &&
		!IsApplication(artefact) &&
		!IsDeployment(artefact) {

		return true
	}

	return false
}

func IsDeployment(artefact *Artefact) bool {

	if artefact != nil && artefact.Parent != nil {
		if artefact.Parent.Name == DEPLOYMENT_WRAP_FOLDER_NAME {
			return true
		}
	}

	return false
}

func FlatPreorder(root *Artefact, result []*Artefact) []*Artefact {

	if root == nil {
		return result
	}

	result = append(result, root)

	if len(root.Content) > 0 {
		for _, child := range root.Content {
			result = FlatPreorder(child, result)
		}
	}

	return result
}

// checks wether the given artefact has a predecessor with given id
func hasAnyPredecessor(artefact *Artefact, predecessorId string) bool {

	current := artefact
	for current != nil {
		if current.Id == predecessorId {
			return true
		}
		current = current.Parent
	}

	return false
}

// selects all predecessors of an artefact, which have the same name
func predecessorsByName(artefact *Artefact) []*Artefact {

	predecessors := make([]*Artefact, 0)
	if artefact == nil || artefact.Parent == nil {
		return predecessors
	}

	current := artefact.Parent
	for current != nil {

		if strings.EqualFold(current.Name, artefact.Name) {
			predecessors = append(predecessors, current)
		}

		for _, currentChild := range current.Content {

			if strings.EqualFold(currentChild.Name, artefact.Name) {
				predecessors = append(predecessors, currentChild)
			}
		}

		current = current.Parent
	}

	return predecessors
}

func FindArtefactByPath(path string) *Artefact {

	artefactTree, err := LoadArtefactTree()
	if err != nil {
		config.LOGGER.LogError(err.Error(), err, "NO-TRACEID")
		return nil
	}

	return findArtefactByPath(artefactTree, path)
}

func findArtefactByPath(artefact *Artefact, path string) *Artefact {

	if artefact == nil {
		return nil
	}

	if strings.EqualFold(FullPath(artefact), path) {
		return artefact
	}

	for _, next := range artefact.Content {

		match := findArtefactByPath(next, path)
		if match != nil {
			return match
		}
	}

	return nil
}

func FullPath(artefact *Artefact) string {

	current := artefact
	path := ""
	for current != nil {
		path = filepath.Join(current.Name, path)
		current = current.Parent
	}

	return filepath.Join(config.GLOBAL.RepoDir(), path)
}

func loadArtefactTreeUntilDepth(path string, parent *Artefact, depth int) (*Artefact, error) {

	artefact, err := NewArtefact(path, parent)
	if err != nil {
		return nil, err
	}

	if artefact.ArtefactType != FILE && (depth > 0 || depth <= LOAD_ENTIRE_TREE) {

		dir, err := utils.FileOpen(path, int(os.ModeDir.Perm()))
		if err != nil {
			return nil, err
		}
		defer dir.Close()

		dirEntries, err := dir.ReadDir(-1) // reads all entries
		for _, entry := range dirEntries {
			if entry.Type().IsDir() || entry.Type().IsRegular() {
				// we are only interested in "normal" files and directories
				childPath := filepath.Join(path, entry.Name())
				_, err := loadArtefactTreeUntilDepth(childPath, artefact, depth-1)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return artefact, nil
}

func children(artefact *Artefact) ([]*Artefact, error) {

	childs := make([]*Artefact, 0)
	path := FullPath(artefact)
	dir, err := utils.FileOpen(path, int(os.ModeDir.Perm()))
	if err != nil {
		return make([]*Artefact, 0), err
	}
	defer dir.Close()

	dirEntries, err := dir.ReadDir(-1) // reads all entries
	for _, entry := range dirEntries {
		if entry.Type().IsDir() || entry.Type().IsRegular() {
			// we are only interested in "normal" files and directories
			childPath := filepath.Join(path, entry.Name())
			chld, err := NewArtefact(childPath, artefact)
			if err != nil {
				return make([]*Artefact, 0), err
			}
			childs = append(childs, chld)
		}
	}

	return childs, nil
}

func writeToString(artefact *Artefact, level int, representation string) string {

	internalIntendation := level
	for internalIntendation > 0 {
		representation += "   "
		internalIntendation--
	}

	representation += artefact.Name
	if artefact.ArtefactType == FILE {
		representation += " (FILE, ID: " + artefact.Id + ")"
	} else {
		representation += " (DIR, ID: " + artefact.Id + ")"
	}
	representation += "\n"

	for _, child := range artefact.Content {

		representation = writeToString(child, level+1, representation)
	}

	return representation
}
