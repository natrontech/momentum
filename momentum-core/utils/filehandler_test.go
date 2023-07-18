package utils_test

import (
	"momentum-core/utils"
	"testing"
)

func TestFileEmpty(t *testing.T) {
	testFile := "./testdata/emptyfile.yaml"

	if utils.FileExists(testFile) {

		isEmpty := utils.FileEmpty(testFile)
		if !isEmpty {
			t.Errorf("Test file is empty but not recongized as such.")
		}
		return
	}

	t.Errorf("Test File %s is not existent", testFile)
}
