package tree

import (
	"errors"
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

func AddContent(parent *Node, content *Node) error {

	if parent == nil || content == nil {
		return errors.New("parent or content is nil")
	}

	if parent.YamlNode.Kind != yaml.DocumentNode &&
		parent.YamlNode.Kind != yaml.MappingNode &&
		parent.YamlNode.Kind != yaml.SequenceNode {
		return errors.New("can only add content to document mapping and sequence yaml nodes")
	}

	if parent.Kind != File &&
		parent.Kind != Mapping &&
		parent.Kind != Sequence {
		return errors.New("can only add content to file mapping and sequence nodes")
	}

	var err error = nil
	p := parent
	if parent.Kind == File {
		p, err = parent.FileMapping()
		if err != nil {
			return errors.New("unable to retrieve parent files mapping")
		}
	}

	p.YamlNode.Content = append(p.YamlNode.Content, content.YamlNode)
	p.AddChild(content)

	return nil
}

func CreatePropertyNodes(key string, value string, valueTag string, style yaml.Style) (*yaml.Node, *yaml.Node) {

	keyNode := CreateScalarNode(key, StrTag, style)
	valueNode := CreateScalarNode(value, valueTag, style)

	return keyNode, valueNode
}

func CreateSequenceNode(key string, style yaml.Style) *yaml.Node {

	n := new(yaml.Node)

	n.Kind = yaml.SequenceNode
	n.Tag = seqTag
	n.Value = key
	n.Style = style

	return n
}

func CreateMappingNode(key string, style yaml.Style) *yaml.Node {

	n := new(yaml.Node)

	n.Kind = yaml.MappingNode
	n.Tag = mapTag
	n.Value = key
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
