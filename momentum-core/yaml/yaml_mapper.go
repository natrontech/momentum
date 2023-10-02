package yaml

import (
	"errors"
	"momentum-core/utils"
	"strings"
)

const EMPTY_SCALAR = ""

func MomentumTreeFromYaml(n *Node, filePath string) (*ViewNode, error) {

	if n == nil {

		return nil, errors.New("yaml ViewNode is nil")
	}

	if n.Kind != DocumentNode {

		return nil, errors.New("only yaml nodes of kind documentnode can be loaded into the tree")
	}

	fileNode := NewNode(File, n.Value, n.Value, nil, nil, n)

	for _, child := range n.Content {
		_, err := momentumFileTree(child, fileNode)
		if err != nil {

			return nil, err
		}
	}
	fileNode.Path = strings.ReplaceAll(utils.LastPartOfPath(filePath), ".", FILE_ENDING_SEPARATOR_REPLACEMENT)

	return fileNode, nil
}

func momentumFileTree(yamlNode *Node, parent *ViewNode) (*ViewNode, error) {

	if yamlNode == nil {

		return nil, errors.New("yaml ViewNode is nil")
	}

	if parent == nil {

		return nil, errors.New("parent ViewNode is nil")
	}

	mapped, parent, doAddChild, err := mapYamlNode(yamlNode, parent)
	if err != nil {

		return nil, err
	}

	if doAddChild {

		parent.AddChild(mapped)
	}

	if len(yamlNode.Content) > 0 {
		for _, child := range yamlNode.Content {
			_, err := momentumFileTree(child, mapped)
			if err != nil {
				return nil, err
			}
		}
	}

	return parent, nil
}

func mapYamlNode(n *Node, parent *ViewNode) (*ViewNode, *ViewNode, bool, error) {

	if n == nil {

		return nil, parent, false, errors.New("node is nil")
	}

	momentumKind, err := momentumKind(n, parent)
	if err != nil {

		return nil, parent, false, err
	}

	doAddChild := true
	var momentumNode *ViewNode

	switch momentumKind {

	case Mapping:

		momentumNode, doAddChild, err = handleMapEntries(n, parent)
		if err != nil {
			return nil, parent, false, err
		}
	case Sequence:

		momentumNode, err = handleSequenceEntries(n, parent)
		if err != nil {
			return nil, parent, false, err
		}
	case Property:

		momentumNode, doAddChild, err = handlePropertyEntries(n, parent)
		if err != nil {
			return nil, parent, false, err
		}
	case Value:

		momentumNode, doAddChild, err = handleValueEntries(n, parent)
		if err != nil {
			return nil, parent, false, err
		}
	default:

		return nil, parent, false, errors.New("unallowed kind")
	}

	return momentumNode, parent, doAddChild, nil
}

func YamlTree(n *ViewNode) (*Node, error) {

	if n.Kind == Directory {

		return nil, errors.New("unable to create yaml tree from directory")
	}

	if n.YamlNode == nil {

		return nil, errors.New("for this node no yaml node exists")
	}

	return n.YamlNode, nil
}

func handleValueEntries(n *Node, parent *ViewNode) (*ViewNode, bool, error) {

	if n.Kind != ScalarNode {

		return nil, false, errors.New("expecting scalar ViewNode")
	}

	if parent.Kind == Sequence {

		return NewNode(Value, EMPTY_SCALAR, n.Value, nil, nil, n), true, nil
	}

	lastChildIndex := len(parent.Children) - 1
	if (parent.Kind == Mapping || parent.Kind == File) &&
		lastChildIndex > -1 &&
		parent.Children[lastChildIndex].Kind == Property {

		parent.Children[lastChildIndex].Value = n.Value
		parent.Children[lastChildIndex].YamlNode = n

		return parent.Children[lastChildIndex], false, nil
	}

	return nil, false, errors.New("unallowed combination of kind value and ViewNode parents kind (parent kind: " + ToHumanReadableKind(parent.Kind) + ")")
}

func handlePropertyEntries(n *Node, parent *ViewNode) (*ViewNode, bool, error) {

	if n.Kind != ScalarNode {

		return nil, false, errors.New("expecting scalar ViewNode")
	}

	return NewNode(Property, n.Value, EMPTY_SCALAR, nil, nil, nil), true, nil
}

/*
* Most of the Sequence and Mapping entries are relevant to the parser, but not to us.
* This function handles edge cases for sequences.
 */
func handleSequenceEntries(n *Node, parent *ViewNode) (*ViewNode, error) {

	if n.Kind != SequenceNode {

		return nil, errors.New("expecting sequence ViewNode")
	}

	lastChildIndex := len(parent.Children) - 1
	if lastChildIndex > -1 &&
		parent.Children[lastChildIndex].Kind == Property &&
		parent.Children[lastChildIndex].Value == EMPTY_SCALAR {

		emptyPropertyNode := parent.Children[lastChildIndex]
		newNode := NewNode(Sequence, emptyPropertyNode.Path, emptyPropertyNode.Value, nil, nil, n)
		if newNode.Path != EMPTY_SCALAR {
			newNode.IsWrapping = false
		}
		parent.AddChild(newNode)
		emptyPropertyNode.Remove()
		return newNode, nil
	}

	return NewNode(Sequence, n.Value, EMPTY_SCALAR, parent, nil, n), nil
}

/*
* Most of the Sequence and Mapping entries are relevant to the parser, but not to us.
* This function handles edge cases for maps.
 */
func handleMapEntries(n *Node, parent *ViewNode) (*ViewNode, bool, error) {

	if n.Kind != MappingNode {

		return nil, false, errors.New("expecting mapping ViewNode")
	}

	lastChildIndex := len(parent.Children) - 1
	if lastChildIndex > -1 &&
		parent.Children[lastChildIndex].Kind == Property &&
		parent.Children[lastChildIndex].Value == EMPTY_SCALAR {

		emptyPropertyNode := parent.Children[lastChildIndex]
		newNode := NewNode(Mapping, emptyPropertyNode.Path, emptyPropertyNode.Value, nil, nil, n)
		if newNode.Path != EMPTY_SCALAR {
			newNode.IsWrapping = false
		}
		parent.AddChild(newNode)
		emptyPropertyNode.Remove()
		return newNode, false, nil
	}

	if parent.Kind == Sequence {
		newNode := NewNode(Mapping, "", "", nil, nil, n)
		newNode.IsWrapping = true
		return newNode, true, nil
	}

	newNode := NewNode(Mapping, n.Value, EMPTY_SCALAR, nil, nil, n)
	if parent.Kind == File {
		// first mapping of file is always wrapping.
		newNode.IsWrapping = true
	}

	return newNode, true, nil
}

func yamlNodeIsMapValue(parent *ViewNode) bool {

	if parent == nil || parent.Kind != Mapping {
		return false
	}

	lastChildIndex := len(parent.Children) - 1
	if lastChildIndex < 0 {
		// no child yet. -> first child is always a property
		return false
	}

	return parent.Children[lastChildIndex].Kind == Property && parent.Children[lastChildIndex].Value == EMPTY_SCALAR
}

func momentumKind(n *Node, parent *ViewNode) (NodeKind, error) {

	switch n.Kind {

	case DocumentNode:

		return File, nil
	case MappingNode:

		return Mapping, nil
	case SequenceNode:

		return Sequence, nil
	case AliasNode:

		return -1, errors.New("kind alias currently not supported by momentum")
	case ScalarNode:

		if yamlNodeIsMapValue(parent) {
			return Value, nil
		} else if parent.Kind == Sequence {
			return Value, nil
		} else {
			return Property, nil
		}
	default:

		return -1, errors.New("unknown yaml kind")
	}
}

func ToHumanReadableKind(k NodeKind) string {

	switch k {
	case Directory:
		return "Directory"
	case File:
		return "File"
	case Mapping:
		return "Mapping"
	case Sequence:
		return "Sequence"
	case Property:
		return "Property"
	case Value:
		return "Value"
	default:
		return "UNKNOWN"
	}
}
