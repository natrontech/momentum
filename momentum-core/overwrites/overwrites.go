package overwrites

import (
	"errors"
	"momentum-core/artefacts"
	"momentum-core/backtracking"
	"momentum-core/yaml"
	"strconv"
	"strings"
)

type OverwriteProvider func(origin *artefacts.Artefact, originLineNumber int) ([]*Overwrite, error)

// extend this list to add more overwriting rules
var ActiveOverwriteProviders []OverwriteProvider = []OverwriteProvider{defaultOverwriteBehaviour, ruleEngineOverwriteBehaviour}

func OverwriteConfigByArtefact(artefactId string) *OverwriteConfig {

	return nil
}

func defaultOverwriteBehaviour(origin *artefacts.Artefact, originLineNumber int) ([]*Overwrite, error) {

	fileNode, err := yaml.ParseFile(artefacts.FullPath(origin))
	if err != nil {
		return make([]*Overwrite, 0), err
	}

	lineNode := yaml.FindNodeByLine(fileNode, originLineNumber)
	if lineNode == nil {
		err := errors.New("could not find line " + strconv.Itoa(originLineNumber) + " in file " + artefacts.FullPath(origin))
		return make([]*Overwrite, 0), err
	}

	overwritingFiles := overwritesByFilenamePriorityAsc(artefacts.FullPath(origin))

	if len(overwritingFiles) > 0 {

		predicate := yaml.ToMatchableSearchTerm(lineNode.FullPath())
		predicate = strings.Join(strings.Split(predicate, ".")[1:], ".") // remove filename prefix

		overwrites := make([]*Overwrite, 0)
		for _, overwriting := range overwritingFiles {

			backtracker := backtracking.NewPropertyBacktracker(predicate, artefacts.FullPath(overwriting), backtracking.NewYamlPropertyParser())
			var result []*backtracking.Match[string, yaml.ViewNode] = backtracker.RunBacktrack()

			for _, match := range result {

				overwrite := new(Overwrite)
				overwrite.OriginFileId = origin.Id
				overwrite.OriginFileLine = originLineNumber
				overwrite.OverwriteFileLine = match.MatchNode.Pointer.YamlNode.Line
				overwrite.OverwriteFileId = overwriting.Id

				overwrites = append(overwrites, overwrite)
			}
		}

		return overwrites, nil
	}

	return make([]*Overwrite, 0), nil
}

// gets all files which are higher up in the structure with the same name as the given file path.
// first item in result is most important
func overwritesByFilenamePriorityAsc(path string) []*artefacts.Artefact {

	overwritable := artefacts.FindArtefactByPath(path)
	if overwritable != nil {

		overwritesOrderedByPriorityAsc := make([]*artefacts.Artefact, 0)
		current := overwritable
		for current != nil {

			for _, child := range current.Content {

				if strings.EqualFold(overwritable.Name, child.Name) && !strings.EqualFold(overwritable.Id, child.Id) {
					overwritesOrderedByPriorityAsc = append(overwritesOrderedByPriorityAsc, child)
				}
			}

			current = current.Parent
		}

		return overwritesOrderedByPriorityAsc
	}

	return make([]*artefacts.Artefact, 0)
}

func ruleEngineOverwriteBehaviour(origin *artefacts.Artefact, originLineNumber int) ([]*Overwrite, error) {

	return make([]*Overwrite, 0), nil
}

func rulesByFilename(path string) []*artefacts.Artefact {

	return make([]*artefacts.Artefact, 0)
}
