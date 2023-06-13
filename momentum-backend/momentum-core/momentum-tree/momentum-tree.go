package tree

import (
	"strings"
)

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

	stgsAndApps := n.AllStages()
	stgsAndApps = append(stgsAndApps, n.Apps()...)
	depls := make([]*Node, 0)

	for _, stage := range stgsAndApps {
		depls = append(depls, deployments(stage)...)
	}

	return depls
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
