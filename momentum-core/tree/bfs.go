package tree

import (
	"strings"
)

func ToMatchableSearchTerm(fullPath string) string {

	return strings.ReplaceAll(strings.ReplaceAll(fullPath, ".", "::"), "/", ".")
}

func BreathFirstSearch(term string, startNode *Node) []*Node {

	if strings.Contains(term, HIERARCHY_SEPARATOR) {

		return bfsMultilevelPathMatch(SortedTermsPathList(term), startNode, make([]*Node, 0))
	} else {

		return bfsPathMatch(term, startNode, make([]*Node, 0))
	}
}

/*
* Finds all nodes which paths match the search term.
 */
func bfsPathMatch(term string, n *Node, result []*Node) []*Node {

	if n == nil {
		return result
	}

	if !n.IsWrapping && n.Path == term {
		result = append(result, n)
	}

	if n.Children != nil || len(n.Children) > 0 {

		for _, child := range n.Children {
			result = bfsPathMatch(term, child, result)
		}
	}

	return result
}

/*
* Finds all nodes which pathes correspond to the given hierarchy relative to given start node.
 */
func bfsMultilevelPathMatch(sortedTerms []string, startNode *Node, result []*Node) []*Node {

	term := BuildSortedTermsPath(sortedTerms...)

	return bfsMultilevelPathMatchByTerm(term, startNode, result)
}

func bfsMultilevelPathMatchByTerm(term string, n *Node, result []*Node) []*Node {

	// fmt.Println("Searching in", ToMatchableSearchTerm(n.FullPath()), "for", term)

	if strings.Contains(ToMatchableSearchTerm(n.FullPath()), term) {

		if len(result) > 1 {
			for _, res := range result {

				if strings.LastIndex(ToMatchableSearchTerm(n.FullPath()), term) != strings.LastIndex(ToMatchableSearchTerm(res.FullPath()), term) {
					result = append(result, n)
				}
			}
		} else {
			result = append(result, n)
		}
	}

	for _, child := range n.Children {
		result = bfsMultilevelPathMatchByTerm(term, child, result)
	}

	return result
}
