package backtracking

type Backtracker[T any, P any] interface {
	// The function executes the backtrack algorithm as configured and returns the result
	RunBacktrack() []*Match[T, P]
}

type Input[I any, T any, P any] interface {
	// GetInput shall return the input of a configured source, such that the Parse function can parse the input into the Node structure
	GetInput() I
	// Parse shall map the input source into the node structure
	Parse(I) *Node[T, P]
	// GetSearch shall return the instance of the search
	GetSearch() Search[T, P]
}

type Search[T any, P any] interface {
	// The predicate is the fragment, the backtracking algorithm searches for
	Predicate() *T
	// The Comparable shall process a node and generate an object comparable to the predicate
	Comparable(*Node[T, P]) *T
	// IsMatch compares the predicate with a given compare value and returns true if they match and false if they do not
	IsMatch(*T, *T) bool
	// StopEarly shall return true if the algorithm must backtrack and reject the current branch.
	// False if the algorithm shall process further even if no full match has been reached.
	// Be aware that a slightly wrong implementation can lead to ommitted results.
	StopEarly(*T, *T) bool
}

// Node is the structure used inside the backtrack algorithm and describes a tree structure.
type Node[T any, P any] struct {
	value    *T
	Pointer  *P // pointer to the actual element which belongs to the node
	parent   *Node[T, P]
	children []*Node[T, P]
}

type Match[T any, P any] struct {
	Input     *T
	Match     *T
	MatchNode *Node[T, P]
}

// Implementation of a backtracking algorithm to find matches on path as described on Wikipedia: https://en.wikipedia.org/wiki/Backtracking
func Backtrack[I any, T any, P any](inp Input[I, T, P]) []*Match[T, P] {

	result := make([]*Match[T, P], 0)
	return backtrack(inp.GetSearch(), root(inp.Parse(inp.GetInput())), result)
}

func backtrack[T any, P any](problem Search[T, P], n *Node[T, P], result []*Match[T, P]) []*Match[T, P] {

	if n == nil {
		return result
	}

	// Currently no rejection makes sense
	// if reject(problem, n) {
	// 	return backtrack(problem, next(problem, n), result)
	// }

	if accept(problem, n) {
		result = appendOutput(problem, n, result)
	}

	return backtrack(problem, next(problem, n), result)
}

func root[T any, P any](n *Node[T, P]) *Node[T, P] {

	current := n
	for current.parent != nil {
		current = current.parent
	}

	return current
}

func reject[T any, P any](search Search[T, P], n *Node[T, P]) bool {

	// can be dangerous, because results are ommited possibly ommitted
	return search.StopEarly(search.Predicate(), search.Comparable(n))
}

func accept[T any, P any](search Search[T, P], n *Node[T, P]) bool {

	return search.IsMatch(search.Predicate(), search.Comparable(n))
}

func first[T any, P any](search Search[T, P], n *Node[T, P]) *Node[T, P] {

	return n.children[0]
}

func next[T any, P any](search Search[T, P], n *Node[T, P]) *Node[T, P] {

	if n == nil {
		return nil
	}

	if len(n.children) == 0 {
		// backtrack
		return walkUp(n.parent, n)
	}

	return first(search, n)
}

func walkUp[T any, P any](parent *Node[T, P], current *Node[T, P]) *Node[T, P] {

	found := false
	if parent != nil && current != nil && current.parent == parent {
		for _, chld := range parent.children {
			if found {
				return chld
			} else if chld == current {
				found = true
			}
		}
	}

	if found && parent.parent != nil {
		// this indicates that the current node was the last of this level, thus we need to walk up further
		// if the parent is nil, this indicates we are at root and traversed the whole tree
		if parent.parent == nil {
			return nil
		}

		return walkUp(parent.parent, parent)
	}

	return nil
}

func appendOutput[T any, P any](search Search[T, P], n *Node[T, P], matches []*Match[T, P]) []*Match[T, P] {

	match := new(Match[T, P])
	match.Input = search.Predicate()
	match.Match = n.value
	match.MatchNode = n

	matches = append(matches, match)

	return matches
}
