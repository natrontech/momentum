package tree

import (
	"errors"
	"strings"

	utils "momentum/momentum-core/momentum-utils"
	yaml "momentum/momentum-core/momentum-yaml"
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

type Node struct {
	Kind NodeKind
	// indicates wether the node is wrapping or not
	// (a wrapping node has an empty path and must be ignored while searching)
	IsWrapping bool
	Path       string
	Value      string // Value only set for Kind == Value
	Parent     *Node
	Children   []*Node
	// only sub-trees of (Yaml-)File nodes have YamlNode. otherwise nil.
	// facilitates writing to writable yaml tree.
	YamlNode *yaml.Node
}

func NewNode(kind NodeKind, path string, value string, parent *Node, children []*Node, yamlNode *yaml.Node) *Node {

	n := new(Node)

	n.Kind = kind
	n.Path = strings.ReplaceAll(utils.LastPartOfPath(path), ".", FILE_ENDING_SEPARATOR_REPLACEMENT)
	n.Value = value
	n.Parent = parent
	n.YamlNode = yamlNode

	if children == nil || len(children) < 1 {
		n.Children = make([]*Node, 0)
	} else {
		n.Children = children
	}

	if kind == Sequence || kind == Mapping {
		// when setting nodes path, this must be changed
		n.IsWrapping = true
	}

	return n
}

func (n *Node) Remove() {

	if n.Parent == nil || len(n.Parent.Children) < 1 {
		return
	}

	newChilds := make([]*Node, 0)
	for _, child := range n.Parent.Children {
		if child != n {
			newChilds = append(newChilds, child)
		}
	}
	n.Parent.Children = newChilds
	n.Parent = nil
}

func (n *Node) AddChild(child *Node) {

	if child == nil {
		return
	}

	if child.Parent != nil {
		child.Remove()
	}

	child.Parent = n

	n.Children = append(n.Children, child)
}

func (n *Node) SetValue(v string) {

	n.Value = v

	if n.YamlNode != nil {
		n.YamlNode.Value = v
	}
}

/*
* Returns full path from root to this node.
 */
func (n *Node) FullPath() string {
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

/*
* Returns path for use in filesystem interactions.
 */
func (n *Node) NormalizedPath() string {
	return strings.ReplaceAll(n.Path, FILE_ENDING_SEPARATOR_REPLACEMENT, ".")
}

/*
* If path has special ending '::[anything]' its cut otherwise just n.Path
 */
func (n *Node) PathWithoutEnding() string {
	return strings.Split(n.Path, FILE_ENDING_SEPARATOR_REPLACEMENT)[0]
}

func (n *Node) Write(allowOverwrite bool) error {

	_, err := WriteNode(n, allowOverwrite)
	if err != nil {
		return err
	}
	return nil
}

func (n *Node) IsRoot() bool {
	return n.Parent == nil || n == n.Parent // if nodes parent is reference to itself its the parent.
}

func (n *Node) Root() *Node {

	current := n
	for !current.IsRoot() {
		current = current.Parent
	}
	return current
}

func (n *Node) Files() []*Node {

	if n == nil || n.Kind != Directory {
		return make([]*Node, 0)
	}

	files := make([]*Node, 0)
	for _, child := range n.Children {
		if child.Kind == File {
			files = append(files, child)
		}
	}
	return files
}

func (n *Node) Directories() []*Node {

	if n == nil || n.Kind != Directory {
		return make([]*Node, 0)
	}

	directories := make([]*Node, 0)
	for _, child := range n.Children {
		if child.Kind == Directory {
			directories = append(directories, child)
		}
	}
	return directories
}

func (n *Node) Search(term string) []*Node {

	return BreathFirstSearch(term, n)
}

func (n *Node) FindFirst(term string) (*Node, bool) {

	results := n.Search(term)
	if results == nil || len(results) < 1 {
		return nil, false
	}
	return results[0], true
}

func (n *Node) AddSequenceValue(value string, style yaml.Style) error {

	if n.Kind != Sequence || n.YamlNode.Kind != yaml.SequenceNode {
		return errors.New("can only add sequence value to node of type sequence")
	}

	sequenceValue := CreateSequenceValueNode(value, strTag, style)
	n.YamlNode.Content = append(n.YamlNode.Content, sequenceValue)
	momentumNode := NewNode(Value, "", value, n, nil, sequenceValue)
	n.AddChild(momentumNode)

	return nil
}
