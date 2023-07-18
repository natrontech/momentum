package tree

import "strings"

const META_PREFIX = "_"
const HIERARCHY_SEPARATOR = "."

func BuildSortedTermsPath(parts ...string) string {

	p := ""
	for _, part := range parts {
		p += part + HIERARCHY_SEPARATOR
	}
	p, _ = strings.CutSuffix(p, HIERARCHY_SEPARATOR)

	return p
}

func SortedTermsPathList(term string) []string {
	return strings.Split(term, HIERARCHY_SEPARATOR)
}

/*
* Builds a match tree from a given term which separates nodes with the HIERARCHY_SEPARATOR
 */
func BuildMatchTree(term string) *Node {

	parts := strings.Split(term, HIERARCHY_SEPARATOR)

	root := new(Node)
	root.Path = parts[0]

	current := root
	if len(parts) > 1 {
		for i := 1; i < len(parts); i++ {

			newNode := new(Node)
			newNode.Path = parts[i]
			current.Children = append(current.Children, newNode)
			current = newNode
		}
	}

	return root
}
