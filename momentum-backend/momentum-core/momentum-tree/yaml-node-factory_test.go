package tree_test

import (
	"fmt"
	tree "momentum/momentum-core/momentum-tree"
	utils "momentum/momentum-core/momentum-utils"
	"os"
	"strings"
	"testing"
)

func FILESYSTEMTEST_TestAddSequenceValue(t *testing.T) {

	wdir, err := os.Getwd()
	if err != nil {
		fmt.Println("cannot read workdir")
		t.FailNow()
	}

	TEST_FILE_PATH := utils.BuildPath(wdir, "testdata/node_manipulations.yaml")

	success := writeTestFile(TEST_FILE_PATH)
	if !success {
		fmt.Println("unable to instantiate testfile")
		t.FailNow()
	}

	parsed, err := tree.Parse(TEST_FILE_PATH, []string{})
	if err != nil {
		fmt.Println("failed to parse testfile")
		t.FailNow()
	}

	utils.FileDelete(TEST_FILE_PATH)

	resourcesSequence, found := parsed.FindFirst("resources")
	if !found {
		fmt.Println("unable to find resources sequence in testfile")
		t.FailNow()
	}

	resourcesSequence.Children[0].SetValue("Ciao")
	err = resourcesSequence.AddSequenceValue("World")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}

	tree.Print(resourcesSequence)

	err = resourcesSequence.Write(true)
	if err != nil {
		fmt.Println("writing new resources failed:", err.Error())
		t.FailNow()
	}

	afterWriteTree, err := tree.Parse(TEST_FILE_PATH, nil)
	if err != nil {
		fmt.Println("parsing tree failed after adding resources")
		t.FailNow()
	}

	resourcesSequence, found = afterWriteTree.FindFirst("resources")
	if !found {
		fmt.Println("unable to find resources sequence in testfile after update")
		t.FailNow()
	}

	tree.Print(resourcesSequence)

	if len(resourcesSequence.Children) != 2 {
		fmt.Println("expected two children in resources")
		t.FailNow()
	}

	if resourcesSequence.Children[0].Value != "Ciao" {
		fmt.Println("expected first child to be 'Hello'")
		t.FailNow()
	}

	if resourcesSequence.Children[1].Value != "World" {
		fmt.Println("expected first child to be 'World'")
		t.FailNow()
	}

	cleanup(t, TEST_FILE_PATH)
}

func writeTestFile(p string) bool {
	return utils.FileWriteLines(p, strings.Split("kind: Test\nresources:\n- \"Hello\"", "\n"))
}

func cleanup(t *testing.T, p string) {
	err := os.Remove(p)

	if err != nil {
		fmt.Println("unable to clean up after test")
		t.FailNow()
	}
}
