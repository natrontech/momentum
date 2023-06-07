package tree

import (
	"errors"
	"fmt"
	"strings"
)

/*
* BE AWARE: Write will overwrite everything without asking.
* You have to check first, that you don't overwrite stuff
* which shall not be overwritten.
*
* If the given node is not of Kind File or Directory, it will
* look for any parent which matches the criteria and write the parent
 */
func WriteNode(n *Node) (string, error) {

	writableNode, err := writableParent(n)
	if err != nil {
		return "", err
	}

	if writableNode.Kind == Directory {
		return writeDir(n)
	} else if writableNode.Kind == File {
		return writeFile(n)
	}

	return "", errors.New("unknown writable node kind")
}

func writeDir(n *Node) (string, error) {

	if n.Kind != Directory {
		return "", errors.New("kind directory expected")
	}

	fmt.Println("Writing DIR", n.FullPath())
	// err := utils.DirCreate(n.FullPath())
	// if err != nil {
	// 	return "", err
	// }

	for _, child := range n.Children {
		if isWritable(child) {
			WriteNode(child)
		}
	}

	return n.Path, nil
}

func writeFile(n *Node) (string, error) {

	if n.Kind != File {
		return "", errors.New("kind file expected")
	}

	if n.yamlNode == nil {
		return "", errors.New("file has no yaml node. cannot write file")
	}

	if strings.HasSuffix(n.Path, ".yml") || strings.HasSuffix(n.Path, ".yaml") {
		fmt.Println("Writing FILE", n.FullPath())
		// yaml.EncodeToFile(n.yamlNode, n.FullPath())
	} else {
		return "", errors.New("only files which end with .yml or .yaml can be written")
	}

	return n.Path, nil
}

func writableParent(n *Node) (*Node, error) {

	if n == nil {
		return nil, errors.New("node supplied was nil")
	}

	if isWritable(n) {
		return n, nil
	}

	if n.Parent == nil {
		return nil, errors.New("no writable node inside tree")
	}

	return writableParent(n.Parent)
}

func isWritable(n *Node) bool {
	return n.Kind == Directory || n.Kind == File
}
