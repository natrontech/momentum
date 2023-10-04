package yaml

import (
	"errors"
	"fmt"
	"momentum-core/config"
	"momentum-core/utils"
	"strings"
)

type NodeKind int

const (
	Directory NodeKind = 1 << iota
	File
	Property
	Mapping
	Sequence
	Value
)

const FILE_ENDING_SEPARATOR_REPLACEMENT = "::"

type ViewNode struct {
	Id   string
	Kind NodeKind
	// indicates wether the node is wrapping or not
	// (a wrapping node has an empty path and must be ignored while searching)
	IsWrapping bool
	Path       string
	Value      string // Value only set for Kind == Value
	Parent     *ViewNode
	Children   []*ViewNode
	// only sub-trees of (Yaml-)File nodes have YamlNode. otherwise nil.
	// facilitates writing to writable yaml tree.
	YamlNode *Node
}

func NewNode(kind NodeKind, path string, value string, parent *ViewNode, children []*ViewNode, yamlNode *Node) *ViewNode {

	n := new(ViewNode)

	n.Kind = kind
	n.Path = strings.ReplaceAll(utils.LastPartOfPath(path), ".", FILE_ENDING_SEPARATOR_REPLACEMENT)
	n.Value = value
	n.Parent = parent
	n.YamlNode = yamlNode

	if children == nil || len(children) < 1 {
		n.Children = make([]*ViewNode, 0)
	} else {
		n.Children = children
	}

	if kind == Sequence || kind == Mapping {
		// when setting nodes path, this must be changed
		n.IsWrapping = true
	}

	id, err := utils.GenerateId(config.IdGenerationPath(n.FullPath()))
	if err != nil {
		fmt.Println("generating id failed:", err.Error())
	}
	n.Id = id

	return n
}

func ParseFile(path string) (*ViewNode, error) {

	if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
		if utils.FileEmpty(path) {
			return nil, nil
		}
		parseTree := DecodeToTree(path)
		mappedTree, err := MomentumTreeFromYaml(parseTree, path)
		if err != nil {
			return nil, err
		}
		return mappedTree, nil
	}
	return nil, errors.New("unsupported file type")
}

func FindNodeByLine(n *ViewNode, line int) *ViewNode {

	if n.Kind == Directory {
		fmt.Println("cannot find node of type directory")
		return nil
	}

	if n.YamlNode.Line == line {
		return n
	}

	for _, child := range n.Children {
		result := FindNodeByLine(child, line)
		if result != nil {
			return result
		}
	}

	return nil
}

// Returns full path from root to this node.
func (n *ViewNode) FullPath() string {
	path := ""
	current := n
	if current == nil {
		return path
	}
	for current.Parent != nil {
		path = utils.BuildPath(current.Path, path)
		current = current.Parent
	}
	path = utils.BuildPath(current.Path, path)
	return strings.ReplaceAll(path, FILE_ENDING_SEPARATOR_REPLACEMENT, ".")
}

func ToMatchableSearchTerm(fullPath string) string {

	return strings.ReplaceAll(strings.ReplaceAll(fullPath, ".", "::"), "/", ".")
}

func (n *ViewNode) Remove() {

	if n.Parent == nil || len(n.Parent.Children) < 1 {
		return
	}

	newChilds := make([]*ViewNode, 0)
	for _, child := range n.Parent.Children {
		if child != n {
			newChilds = append(newChilds, child)
		}
	}
	n.Parent.Children = newChilds

	n.Parent.RemoveYamlChild(n.Path)

	n.Parent = nil
}

func (n *ViewNode) AddChild(child *ViewNode) {

	if child == nil {
		return
	}

	if child.Parent != nil {
		child.Remove()
	}

	child.Parent = n

	n.Children = append(n.Children, child)
}

func (n *ViewNode) RemoveYamlChild(path string) error {

	updated := make([]*ViewNode, 0)
	for _, child := range n.Children {
		if child.Path != path {
			updated = append(updated, child)
		}

		updatedYaml := make([]*Node, 0)
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
