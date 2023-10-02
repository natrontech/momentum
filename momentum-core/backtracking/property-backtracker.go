package backtracking

import (
	"fmt"
	"momentum-core/yaml"
	"strings"
)

const PROPERTY_SEPARATOR = "."

type PropertyParser interface {
	ParseProperties(string) *Node[string, yaml.ViewNode]
}

// The property backtracker finds matches within a property structure.
// A property search is concepted as hierarchical string separated by dots,
// where each level of the
type PropertyBacktracker struct {
	Backtracker[string, yaml.ViewNode]
	Input[string, string, yaml.ViewNode]
	Search[string, yaml.ViewNode]

	inputParser PropertyParser
	predicate   string
	path        string // path to the yaml file which shall be processed
}

func NewPropertyBacktracker(predicate string, path string, parser PropertyParser) *PropertyBacktracker {

	backtracker := new(PropertyBacktracker)
	backtracker.predicate = predicate
	backtracker.inputParser = parser
	backtracker.path = path

	return backtracker
}

func (backtracker *PropertyBacktracker) RunBacktrack() []*Match[string, yaml.ViewNode] {

	return Backtrack[string, string, yaml.ViewNode](backtracker)
}

func (backtracker *PropertyBacktracker) GetInput() string {

	return backtracker.path
}

func (backtracker *PropertyBacktracker) Parse(input string) *Node[string, yaml.ViewNode] {

	return backtracker.inputParser.ParseProperties(backtracker.path)
}

func (backtracker *PropertyBacktracker) GetSearch() Search[string, yaml.ViewNode] {

	return backtracker
}

func (backtracker *PropertyBacktracker) Predicate() *string {

	return &backtracker.predicate
}

func (backtracker *PropertyBacktracker) Comparable(n *Node[string, yaml.ViewNode]) *string {

	propertyString := ""
	current := n
	for current != nil {
		propertyString = strings.Join([]string{*current.value, propertyString}, PROPERTY_SEPARATOR)
		current = current.parent
	}

	cutted, _ := strings.CutPrefix(propertyString, PROPERTY_SEPARATOR)
	cutted, _ = strings.CutSuffix(cutted, PROPERTY_SEPARATOR)
	//fmt.Println("cutted propertyString:", propertyString)
	return &cutted
}

func (backtracker *PropertyBacktracker) IsMatch(predicate *string, argument *string) bool {

	fmt.Println("trying to match:", *predicate, "=?", *argument)
	return strings.EqualFold(*predicate, *argument)
}

func (backtracker *PropertyBacktracker) StopEarly(predicate *string, argument *string) bool {

	return !strings.HasPrefix(*predicate, *argument)
}
