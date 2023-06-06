package tree

import (
	"errors"
	"fmt"
)

type OverwriteStrategy interface {
	apply(n *Node) error
}

type OverwriteParentStrategy struct {
	root *Node
}

func (ops *OverwriteParentStrategy) apply(n *Node) error {

	// Children will overwrite their parents -> Bottom over Top
	fmt.Println(ops.root)
	return errors.New("not yet implemented")
}
