package tree_test

import (
	"fmt"
	"os"
	"testing"

	tree "momentum/momentum-core/momentum-tree"
	utils "momentum/momentum-core/momentum-utils"
)

func TestNewTree(t *testing.T) {

	path, err := os.Getwd()
	path = utils.BuildPath(path, "testdata")
	if err != nil {
		t.Errorf(err.Error())
	}

	n, err := tree.Parse(path, []string{".git"})
	if err != nil {
		t.Errorf(err.Error())
	}

	tree.Print(n)
}

func TestParseFile(t *testing.T) {

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	testfile := utils.BuildPath(currentDir, "testdata", "kustomization.yaml")

	parsed, err := tree.Parse(testfile, make([]string, 0))
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	tree.Print(parsed)

	if len(parsed.Children) != 1 {
		fmt.Println("expected root to have exactly 1 child")
		t.FailNow()
	}

	if parsed.Path != testfile {
		fmt.Println("expected root to have exactly path", testfile)
		t.FailNow()
	}

	if parsed.Children[0].Kind != tree.Mapping || !parsed.Children[0].IsWrapping {
		fmt.Println("expected root to have kind mapping and be wrapping")
		t.FailNow()
	}

	if len(parsed.Children[0].Children) != 4 {
		fmt.Println("expected root mapping to have exactly four children")
		t.FailNow()
	}

	parsed = parsed.Children[0]

	if parsed.Children[0].Kind != tree.Property || parsed.Children[1].Kind != tree.Property {
		fmt.Println("expected child one and two to be of kind property")
		t.FailNow()
	}

	if parsed.Children[0].Path != "apiVersion" || parsed.Children[1].Path != "kind" {
		fmt.Println("paths didn't match expectations")
		t.FailNow()
	}

	if parsed.Children[0].Value != "kustomize.config.k8s.io/v1beta1" || parsed.Children[1].Value != "Kustomization" {
		fmt.Println("values didn't match expectations")
		t.FailNow()
	}

	if parsed.Children[2].Kind != tree.Sequence {
		fmt.Println("expected third child to be of kind sequence")
		t.FailNow()
	}

	if parsed.Children[2].Path != "resources" {
		fmt.Println("expected third child to have path resources")
		t.FailNow()
	}

	if len(parsed.Children[2].Children) != 1 {
		fmt.Println("expected third child to have exactly one child")
		t.FailNow()
	}

	if parsed.Children[2].Children[0].Kind != tree.Value {
		fmt.Println("expected first child of third child to have kind value")
		t.FailNow()
	}

	if parsed.Children[2].Children[0].Value != "./mywebserver" {
		fmt.Println("value didn't match expectation")
		t.FailNow()
	}

	if parsed.Children[2].Children[0].Value != "./mywebserver" {
		fmt.Println("value didn't match expectation")
		t.FailNow()
	}
}
