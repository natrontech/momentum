package yaml_test

import (
	"fmt"
	"momentum-core/yaml"
	"os"
	"testing"
)

const YAML_TEST_FILE = "./data/yml-parser-test.yml" // "./../data/example.yml"
const YAML_TEST_OUTPUT_FILE = "./data/example-output.yml"

func TestEncodeDecodeRoundTrip(t *testing.T) {

	root := yaml.DecodeToTree(YAML_TEST_FILE)

	printAll(root.Content, 0)

	_, err := yaml.EncodeToFile(root, YAML_TEST_OUTPUT_FILE, false)
	if err != nil {
		fmt.Println("failed writing yaml:", err.Error())
		t.FailNow()
	}

	os.Remove(YAML_TEST_OUTPUT_FILE)
}

func printAll(rs []*yaml.Node, level int) {
	for _, r := range rs {
		print(r, level)
		if r.Content != nil && len(r.Content) > 0 {
			printAll(r.Content, level+1)
		}
	}
}

func print(r *yaml.Node, level int) {
	fmt.Println(spaces(level), "Tag:", r.Tag, " | Value:", r.Value, " | Children:", len(r.Content))
}

func spaces(i int) string {
	s := ""
	for i > 0 {
		s += " "
		i--
	}
	return s
}
