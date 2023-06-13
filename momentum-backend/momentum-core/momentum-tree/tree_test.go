package tree_test

import (
	tree "momentum/momentum-core/momentum-tree"
	"testing"
)

func FILESYSTEMTEST_TestTreeCreation(test *testing.T) {

	tDir := tree.NewNode(tree.Directory, "/testpath", "", nil, nil, nil)
	tFile := tree.NewNode(tree.File, "testfile.yaml", "", tDir, nil, nil)
	tMapping := tree.NewNode(tree.Mapping, "mapping", "", tFile, nil, nil)
	tSequence := tree.NewNode(tree.Sequence, "sequence", "", tFile, nil, nil)
	tProperty := tree.NewNode(tree.Property, "property", "", tMapping, nil, nil)
	tValue := tree.NewNode(tree.Value, "", "value", tSequence, nil, nil)

	if tProperty.FullPath() != "testpath/testfile.yaml/mapping/property" {
		test.FailNow()
	}

	if tValue.FullPath() != "testpath/testfile.yaml/sequence" {
		test.FailNow()
	}
}
