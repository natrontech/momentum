package tree

import (
	"fmt"
	"strings"
)

func DepthFirstSearch(term string, startNode *Node) []*Node {

	if strings.Contains(term, HIERARCHY_SEPARATOR) {

		return dfsMultilevelPathMatch(SortedTermsPathList(term), startNode, make([]*Node, 0))
	} else {

		return dfsPathMatch(term, startNode, make([]*Node, 0))
	}
}

/*
* Finds all nodes which paths match the search term.
 */
func dfsPathMatch(term string, n *Node, result []*Node) []*Node {

	if n == nil {
		return result
	}

	if !n.IsWrapping && n.Path == term {
		result = append(result, n)
	}

	if n.Children != nil || len(n.Children) > 0 {

		for _, child := range n.Children {
			result = dfsPathMatch(term, child, result)
		}
	}

	return result
}

/*
* Finds all nodes which pathes correspond to the given hierarchy relative to given start node.
 */
func dfsMultilevelPathMatch(sortedTerms []string, startNode *Node, result []*Node) []*Node {

	term := BuildSortedTermsPath(sortedTerms...)

	return dfsMultilevelPathMatchByTerm(term, startNode, result)
}

func dfsMultilevelPathMatchByTerm(term string, n *Node, result []*Node) []*Node {

	fmt.Println(toMatchableFullPath(n.FullPath()), term)

	if strings.Contains(toMatchableFullPath(n.FullPath()), term) {

		if len(result) > 1 {
			for _, res := range result {

				if strings.LastIndex(toMatchableFullPath(n.FullPath()), term) != strings.LastIndex(toMatchableFullPath(res.FullPath()), term) {
					result = append(result, n)
				}
			}
		} else {
			result = append(result, n)
		}
	}

	for _, child := range n.Children {
		result = dfsMultilevelPathMatchByTerm(term, child, result)
	}

	return result
}

func toMatchableFullPath(fullPath string) string {

	return strings.ReplaceAll(strings.ReplaceAll(fullPath, ".", "::"), "/", ".")
}
