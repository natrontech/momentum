package tree

import (
	"momentum-core/config"
	"strings"
)

func (n *Node) Repo() *Node {

	return n.Root()
}

func (n *Node) MomentumRoot() *Node {

	momentumRoot, found := n.Repo().FindFirst(config.MOMENTUM_ROOT)
	if !found {
		return nil
	}
	return momentumRoot
}

func (n *Node) Apps() []*Node {

	root := n.MomentumRoot()
	apps := make([]*Node, 0)
	for _, app := range root.Directories() {
		apps = append(apps, app)
	}
	return apps
}

func (n *Node) AllStages() []*Node {

	apps := n.Apps()
	stgs := make([]*Node, 0)
	for _, app := range apps {
		stgs = append(stgs, app.stages()...)
	}
	return stgs
}

func (n *Node) stages() []*Node {

	nodesStages := make([]*Node, 0)
	for _, possibleStage := range n.Directories() {
		if possibleStage.Kind == Directory && !strings.HasPrefix(possibleStage.Path, META_PREFIX) {
			childStages := possibleStage.stages()
			nodesStages = append([]*Node{possibleStage}, childStages...)
		}
	}

	return nodesStages
}

func (n *Node) IsStage() bool {

	stages := n.AllStages()
	for _, stage := range stages {
		if n.FullPath() == stage.FullPath() {
			return true
		}
	}
	return false
}

func (n *Node) FindStage(path string) (bool, *Node) {

	stages := n.AllStages()
	for _, stage := range stages {
		if stage.FullPath() == path {
			return true, stage
		}
	}
	return false, nil
}

func (n *Node) AllDeployments() []*Node {

	stgsAndApps := n.AllStages()
	stgsAndApps = append(stgsAndApps, n.Apps()...)
	depls := make([]*Node, 0)

	for _, stage := range stgsAndApps {
		depls = append(depls, deployments(stage)...)
	}

	return depls
}

func (n *Node) Deployments() []*Node {
	return deployments(n)
}

func (n *Node) Deployment(deploymentId string) *Node {

	for _, deployment := range deployments(n) {
		if deployment.Id == deploymentId {
			return deployment
		}
	}

	return nil
}

func (n *Node) DeploymentFolderFiles() []*Node {

	deploymentFolders := childrenWithName(n.Parent, "_deploy")
	for _, depl := range deploymentFolders {

		if strings.EqualFold(n.PathWithoutEnding(), depl.PathWithoutEnding()) {
			return depl.Files()
		}
	}

	return make([]*Node, 0)
}

func (n *Node) Values() []*Node {

	if n == nil || n.Kind != File {
		return make([]*Node, 0)
	}

	return n.flat(make([]*Node, 0))
}

func (n *Node) flat(result []*Node) []*Node {

	if len(n.Children) > 0 {
		result = append(result, n.Children...)
		for _, child := range n.Children {
			result = append(result, child.flat(result)...)
		}
	}

	return result
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
		if !strings.HasSuffix(f.NormalizedPath(), ".yml") && !strings.HasSuffix(f.NormalizedPath(), ".yaml") {
			continue
		}

		for _, excl := range exclusions {

			if strings.EqualFold(f.NormalizedPath(), excl) {

				fileIsDeployment = false
			}
		}

		if fileIsDeployment {

			for _, dir := range deployFolders {

				if strings.EqualFold(dir.Path, f.PathWithoutEnding()) {
					depls = append(depls, f)
				}
			}
		}
	}

	return depls
}

func childrenWithName(n *Node, name string) []*Node {

	matches := make([]*Node, 0)
	for _, child := range n.Children {
		if strings.EqualFold(child.Path, name) {
			matches = append(matches, child)
		}
	}

	return matches
}
