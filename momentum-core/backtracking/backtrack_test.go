package backtracking_test

import (
	"fmt"
	"momentum-core/backtracking"
	"os"
	"path/filepath"
	"testing"
)

func TestBacktracker(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	testFilePath := filepath.Join(wd, "test.yaml")

	backtracker := backtracking.NewPropertyBacktracker("test.for", testFilePath, backtracking.NewYamlPropertyParser())
	result := backtracker.RunBacktrack()

	if len(result) != 1 {
		fmt.Println("expected one result")
		t.FailNow()
	}

	for _, res := range result {
		fmt.Println(res)
	}
}
