package tree

import (
	"errors"
	"fmt"
	utils "momentum/momentum-core/momentum-utils"
	yaml "momentum/momentum-core/momentum-yaml"
	"strings"
)

/*
* BE AWARE: Write will overwrite everything without asking if allowOverwrite=true is set.
* You have to check first, that you don't overwrite stuff
* which shall not be overwritten.
*
* If the given node is not of Kind File or Directory, it will
* look for any parent which matches the criteria and write the parent
 */
func WriteNode(n *Node, allowOverwrite bool) (string, error) {

	writableNode, err := writableParent(n)
	if err != nil {
		return "", err
	}

	if writableNode.Kind == Directory {
		return writeDir(writableNode, allowOverwrite)
	} else if writableNode.Kind == File {
		return writeFile(writableNode, allowOverwrite)
	}

	return "", errors.New("unknown writable node kind")
}

func writeDir(n *Node, allowOverwrite bool) (string, error) {

	if n.Kind != Directory {
		return "", errors.New("kind directory expected")
	}

	fmt.Println("Writing DIR", n.FullPath())
	err := utils.DirCreate(n.FullPath())
	if err != nil {
		return "", err
	}

	for _, child := range n.Children {
		if isWritable(child) {
			WriteNode(child, allowOverwrite)
		}
	}

	return n.Path, nil
}

func writeFile(n *Node, allowOverwrite bool) (string, error) {

	if n.Kind != File {
		return "", errors.New("kind file expected")
	}

	if n.YamlNode == nil {
		return "", errors.New("file has no yaml node. cannot write file")
	}

	if strings.HasSuffix(n.NormalizedPath(), ".yml") || strings.HasSuffix(n.NormalizedPath(), ".yaml") {
		fmt.Println("Writing FILE", n.FullPath())
		_, err := yaml.EncodeToFile(n.YamlNode, n.FullPath(), allowOverwrite)
		if err != nil {
			return "", err
		}
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
