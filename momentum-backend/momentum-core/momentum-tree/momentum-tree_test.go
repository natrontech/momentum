package tree_test

import (
	"fmt"
	tree "momentum/momentum-core/momentum-tree"
	"testing"
)

const basePath = "testdata"

func FILESYSTEMTEST_TestAllApps(t *testing.T) {

	n, err := tree.Parse(basePath, []string{".git"})
	if err != nil {
		fmt.Println("unable to parse tree:", err.Error())
		t.FailNow()
	}

	apps := n.Apps()
	if len(apps) != 1 {
		fmt.Println("expected exactly one app")
		t.FailNow()
	}

	if apps[0].Path != "testdata_sub" {
		fmt.Println("expected app to be 'testdata_sub', but was", apps[0].Path)
		t.FailNow()
	}
}

func FILESYSTEMTEST_TestAllStages(t *testing.T) {

	n, err := tree.Parse(basePath, []string{".git"})
	if err != nil {
		fmt.Println("unable to parse tree:", err.Error())
		t.FailNow()
	}

	stages := n.AllStages()
	if len(stages) != 3 {
		fmt.Println("expected exactly three stages")
		t.FailNow()
	}

	expectOneToBe(stages, "bronze", t)
	expectOneToBe(stages, "silver", t)
	expectOneToBe(stages, "gold", t)
}

func FILESYSTEMTEST_TestAllDeployments(t *testing.T) {

	n, err := tree.Parse(basePath, []string{".git"})
	if err != nil {
		fmt.Println("unable to parse tree:", err.Error())
		t.FailNow()
	}

	deployments := n.AllDeployments()
	if len(deployments) != 2 {
		fmt.Println("expected exactly two deployments but were", len(deployments))
		t.FailNow()
	}

	expectOneToBe(deployments, "deployment2.yaml", t)
	expectOneToBe(deployments, "deployment3.yaml", t)
}

func expectOneToBe(n []*tree.Node, expected string, t *testing.T) {

	for _, e := range n {

		fmt.Println("checking if", e.NormalizedPath(), "matches expectation", expected)
		if e.NormalizedPath() == expected {
			return
		}
	}

	fmt.Println("expected", expected, "but could not find in", n)
	t.FailNow()
}
