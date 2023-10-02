package artefacts

import (
	"strings"
)

type OverwriteAdvice func(string) []*Artefact

// extend this list to add more overwriting rules
var ActiveOverwriteAdvices []OverwriteAdvice = []OverwriteAdvice{overwritesByFilenamePriorityAsc}

// gets all files which are higher up in the structure with the same name as the given file path.
func overwritesByFilenamePriorityAsc(path string) []*Artefact {

	overwritable := FindArtefactByPath(path)
	if overwritable != nil {

		overwritesOrderedByPriorityAsc := make([]*Artefact, 0)
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

	return make([]*Artefact, 0)
}
