package backtracking

import (
	"fmt"
	"momentum-core/utils"
	"momentum-core/yaml"
)

type YamlPropertyParser struct {
	PropertyParser
}

func NewYamlPropertyParser() *YamlPropertyParser {

	parser := new(YamlPropertyParser)

	return parser
}

func (parser *YamlPropertyParser) ParseProperties(path string) *Node[string, yaml.ViewNode] {

	if !utils.FileExists(path) {
		return nil
	}

	root, err := yaml.ParseFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if root.Children[0].Kind == yaml.Mapping {

		return viewNodeToNode(root.Children[0], nil)
	}

	return nil
}

func viewNodeToNode(n *yaml.ViewNode, parent *Node[string, yaml.ViewNode]) *Node[string, yaml.ViewNode] {

	r := new(Node[string, yaml.ViewNode])
	r.parent = parent
	r.value = &n.Path
	r.Pointer = n

	children := make([]*Node[string, yaml.ViewNode], 0)
	for _, child := range n.Children {
		children = append(children, viewNodeToNode(child, r))
	}
	r.children = children

	return r
}
