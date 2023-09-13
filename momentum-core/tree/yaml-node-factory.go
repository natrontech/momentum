package tree

import (
	"errors"
	"momentum-core/utils"
	"momentum-core/yaml"
)

// copied from yaml, because module private there
const (
	nullTag      = "!!null"
	boolTag      = "!!bool"
	StrTag       = "!!str"
	intTag       = "!!int"
	floatTag     = "!!float"
	timestampTag = "!!timestamp"
	seqTag       = "!!seq"
	mapTag       = "!!map"
	binaryTag    = "!!binary"
	mergeTag     = "!!merge"
)

func (n *Node) RemoveYamlChildren() error {

	errs := make([]error, 0)
	for _, chld := range n.Children {
		errs = append(errs, chld.RemoveYamlChild(chld.Path))
	}

	return errors.Join(errs...)
}

func (n *Node) RemoveYamlChild(path string) error {

	updated := make([]*Node, 0)
	for _, child := range n.Children {
		if child.Path != path {
			updated = append(updated, child)
		}

		updatedYaml := make([]*yaml.Node, 0)
		if child.Path == path {
			if child.Parent != nil && child.Parent.YamlNode != nil {
				if child.Parent.YamlNode != nil {
					for _, yamlChild := range child.Parent.YamlNode.Content {
						if yamlChild.Value != utils.LastPartOfPath(path) {
							updatedYaml = append(updatedYaml, yamlChild)
						}
					}
					child.Parent.YamlNode.Content = updatedYaml
				}
			}
		}

		return nil
	}

	n.Children = updated

	return nil
}

func (n *Node) AddYamlSequence(key string, values []string, style yaml.Style) error {

	if len(values) < 1 {
		return errors.New("sequence must have at least one value")
	}

	var err error = nil
	var anchor *Node = n
	if n.Kind == File {
		anchor, err = n.FileMapping()
		if err != nil {
			return errors.New("failed retrieving file mapping of file node")
		}
	}

	yamlSequenceName := CreateScalarNode(key, StrTag, 0)
	yamlSequenceNode := CreateSequenceNode(style)
	anchor.YamlNode.Content = append(anchor.YamlNode.Content, yamlSequenceName, yamlSequenceNode)

	sequenceNode := NewNode(Sequence, key, "", nil, nil, yamlSequenceNode)
	anchor.AddChild(sequenceNode)

	for _, val := range values {
		err := sequenceNode.AddYamlValue(val, 0)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *Node) AddYamlMapping(key string, style yaml.Style) (*Node, error) {

	var err error = nil
	var anchor *Node = n
	if n.Kind == File {
		anchor, err = n.FileMapping()
		if err != nil {
			return nil, err
		}
	}

	yamlMappingName := CreateScalarNode(key, StrTag, 0)
	yamlMappingNode := CreateMappingNode(key, style)

	anchor.YamlNode.Content = append(anchor.YamlNode.Content, yamlMappingName, yamlMappingNode)

	mappingNode := NewNode(Mapping, key, "", nil, nil, yamlMappingNode)
	anchor.AddChild(mappingNode)

	return mappingNode, nil
}

func (n *Node) AddYamlProperty(key string, value string, valueTag string, style yaml.Style) error {

	if n.Kind != Mapping || n.YamlNode.Kind != yaml.MappingNode {
		return errors.New("properties can only be added to mapping nodes")
	}

	yamlKeyNode, yamlValueNode := CreatePropertyNodes(key, value, valueTag, style)

	mappingNode := NewNode(Property, key, value, n, nil, yamlValueNode)

	n.YamlNode.Content = append(n.YamlNode.Content, yamlKeyNode, yamlValueNode)
	n.AddChild(mappingNode)

	return nil
}

func (n *Node) AddYamlValue(value string, style yaml.Style) error {

	if n.Kind != Sequence || n.YamlNode.Kind != yaml.SequenceNode {
		return errors.New("can only add sequence value to node of type sequence")
	}

	sequenceValue := CreateScalarNode(value, StrTag, style)
	n.YamlNode.Content = append(n.YamlNode.Content, sequenceValue)
	momentumNode := NewNode(Value, "", value, n, nil, sequenceValue)
	n.AddChild(momentumNode)

	return nil
}

func CreatePropertyNodes(key string, value string, valueTag string, style yaml.Style) (*yaml.Node, *yaml.Node) {

	keyNode := CreateScalarNode(key, StrTag, style)
	valueNode := CreateScalarNode(value, valueTag, style)

	return keyNode, valueNode
}

func CreateSequenceNode(style yaml.Style) *yaml.Node {

	n := new(yaml.Node)

	n.Kind = yaml.SequenceNode
	n.Tag = seqTag
	n.Value = ""
	n.Style = style

	return n
}

func CreateMappingNode(key string, style yaml.Style) *yaml.Node {

	n := new(yaml.Node)

	n.Kind = yaml.MappingNode
	n.Tag = mapTag
	n.Value = ""
	n.Style = style

	return n
}

func CreateScalarNode(value string, tag string, style yaml.Style) *yaml.Node {

	n := new(yaml.Node)

	n.Kind = yaml.ScalarNode
	n.Tag = tag
	n.Value = value
	n.Style = style

	return n
}
