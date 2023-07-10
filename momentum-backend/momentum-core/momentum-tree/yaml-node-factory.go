package tree

import (
	"errors"
	yaml "momentum/momentum-core/momentum-yaml"
)

// copied from yaml, because module private there
const (
	nullTag      = "!!null"
	boolTag      = "!!bool"
	strTag       = "!!str"
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

	if parent.YamlNode.Kind != yaml.DocumentNode|yaml.MappingNode|yaml.SequenceNode {
		return errors.New("can only add content to document mapping and sequence yaml nodes")
	}

	if parent.Kind != File|Mapping|Sequence {
		return errors.New("can only add content to file mapping and sequence nodes")
	}

	parent.YamlNode.Content = append(parent.YamlNode.Content, content.YamlNode)
	parent.AddChild(content)

	return nil
}

func CreateSequenceValueNode(value string, tag string, style yaml.Style) *yaml.Node {

	n := new(yaml.Node)

	n.Kind = yaml.ScalarNode
	n.Tag = tag
	n.Value = value
	n.Style = style

	return n
}
