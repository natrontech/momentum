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
	// only sub-trees of (Yaml-)File nodes have yamlNode. otherwise nil.
	// facilitates writing to writable yaml tree.
	yamlNode *yaml.Node
}

func NewNode(kind NodeKind, path string, value string, parent *Node, children []*Node, yamlNode *yaml.Node) *Node {

	n := new(Node)

	n.Kind = kind
	n.Path = strings.ReplaceAll(utils.LastPartOfPath(path), ".", FILE_ENDING_SEPARATOR_REPLACEMENT)
	n.Value = value
	n.Parent = parent
	n.yamlNode = yamlNode

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

	if n.yamlNode != nil {
		n.yamlNode.Value = v
	}
}

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

func (n *Node) Write() error {

	_, err := WriteNode(n)
	if err != nil {
		return err
	}
	return nil
}

func (n *Node) IsRoot() bool {
	return n.Parent == nil
}

func (n *Node) Root() *Node {

	current := n
	for !current.IsRoot() {
		current = n.Parent
	}
	return current
}

func (n *Node) Apps() []*Node {

	root := n.Root()
	apps := make([]*Node, 0)
	for _, app := range root.Children {
		if app.Kind == Directory {
			apps = append(apps, app)
		}
	}
	return apps
}

func (n *Node) AllStages() []*Node {

	apps := n.Apps()
	stgs := make([]*Node, 0)
	for _, app := range apps {
		stgs = append(stgs, stages(app.Children, stgs)...)
	}
	return stgs
}

func (n *Node) AllDeployments() []*Node {

	stgs := n.AllStages()
	depls := make([]*Node, 0)

	for _, stage := range stgs {
		depls = append(depls, deployments(stage)...)
	}

	return depls
}

// selects stages relative to this node
func (n *Node) Stages() ([]*Node, error) {

	if n.Kind != Directory || strings.HasPrefix(n.Path, META_PREFIX) {

		return nil, errors.New("cannot read stages of given node")
	}

	if n.Parent == nil {
		return n.AllStages(), nil
	}

	stgs := make([]*Node, 0)
	for _, stg := range n.Children {
		stgs = append(stgs, stages(stg.Children, stgs)...)
	}
	return stgs, nil
}

// selects deployments relative to the current stage
func (n *Node) Deployments() ([]*Node, error) {

	if n.Kind != Directory || n.Parent == nil {
		return nil, errors.New("cannot read deployments relative to stage if node is not stage")
	}

	return deployments(n), nil
}

func (n *Node) AppExists(app string) (bool, *Node) {

	for _, a := range n.Apps() {

		if strings.EqualFold(a.Path, app) {
			return true, a
		}
	}
	return false, nil
}

func (n *Node) StageExists(app string, stage string) (bool, *Node) {

	appExists, appNode := n.AppExists(app)
	if !appExists {
		return false, nil
	}

	stages, err := appNode.Stages()
	if err != nil {
		return false, nil
	}

	for _, stg := range stages {

		if stg.Path == stage {

			return true, stg
		}
	}

	return false, nil
}

func (n *Node) DeploymentExists(app string, stage string, deployment string) (bool, *Node) {

	stageExists, stageNode := n.StageExists(app, stage)
	if !stageExists {
		return false, nil
	}

	depls, err := stageNode.Deployments()
	if err != nil {
		return false, nil
	}

	for _, dep := range depls {

		if dep.Path == deployment {
			return true, dep
		}
	}

	return false, nil
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

	return DepthFirstSearch(term, n)
}

func stages(parents []*Node, stgs []*Node) []*Node {

	if len(parents) == 0 {
		return stgs
	}

	for _, node := range parents {
		if node.Kind == Directory && !strings.HasPrefix(node.Path, META_PREFIX) {
			stgs = stages(node.Children, stgs)
			stgs = append([]*Node{node}, stgs...)
		}
	}

	return stgs
}

func deployments(stage *Node) []*Node {

	files := stage.Files()
	deployFolders := stage.Search("_deploy")
	if len(deployFolders) < 1 {
		return make([]*Node, 0)
	}
	deployFolder := deployFolders[0]
	deployFolders = deployFolder.Directories()
	exclusions := []string{"kustomization.yml", "kustomization.yaml"}

	depls := make([]*Node, 0)
	for _, f := range files {

		fileIsDeployment := true
		if !strings.HasSuffix(f.Path, ".yml") && !strings.HasSuffix(f.Path, ".yaml") {
			continue
		}

		for _, excl := range exclusions {

			if strings.EqualFold(f.Path, excl) {

				fileIsDeployment = false
			}
		}

		if fileIsDeployment {

			fileNameWithoutEnding := strings.Split(f.Path, ".")[0]
			for _, dir := range deployFolders {

				if strings.EqualFold(dir.Path, fileNameWithoutEnding) {
					depls = append(depls, f)
				}
			}
		}
	}

	return depls
}
