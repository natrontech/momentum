package tree_test

import (
	"strings"
	"testing"

	tree "momentum/momentum-core/momentum-tree"
	utils "momentum/momentum-core/momentum-utils"
)

const testdataDir = "./testdata"
const baseKustomization = "kustomization.yaml"

func TestTreeBfs(t *testing.T) {

	testPath := utils.BuildPath(testdataDir, baseKustomization)
	searchTerm := "kind"

	ftree, err := tree.Parse(testPath, make([]string, 0))
	if err != nil {
		t.Error(err, err.Error())
	}

	tree.Print(ftree)

	result := ftree.Search(searchTerm)

	if len(result) > 1 {
		t.Error("Too many results. Expected only one.")
		t.FailNow()
	}

	if len(result) < 1 {
		t.Error("Expected a result.")
		t.FailNow()
	}

	if result[0].Path != searchTerm || result[0].Value != "Kustomization" {
		t.Error("Expected other result.")
	}
}

func TestMultiLevelPathBfs(t *testing.T) {

	testPath := "testdata"
	searchTerm := "testdata.testdata_sub.kustomization::yaml.kind"

	ftree, err := tree.Parse(testPath, make([]string, 0))
	if err != nil {
		t.Error(err, err.Error())
	}

	tree.Print(ftree)

	result := ftree.Search(searchTerm)

	if len(result) > 1 {
		t.Error("Too many results. Expected only one.")
		t.FailNow()
	}

	if len(result) < 1 {
		t.Error("Expected a result.")
		t.FailNow()
	}

	if strings.Contains(result[0].FullPath(), searchTerm) || result[0].Value != "Kustomization" {
		t.Error("Expected other result.")
	}
}
