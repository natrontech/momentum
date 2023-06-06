package tree

import (
	"fmt"
)

func Print(n *Node) {

	print(n, 0)
}

func ToHumanReadableKind(k NodeKind) string {

	switch k {
	case Directory:
		return "Directory"
	case File:
		return "File"
	case Mapping:
		return "Mapping"
	case Sequence:
		return "Sequence"
	case Property:
		return "Property"
	case Value:
		return "Value"
	default:
		return "UNKNOWN"
	}
}

func print(n *Node, level int) {

	if n == nil {
		return
	}

	if n.Kind == Directory {
		fmt.Println(spaces(level), ToHumanReadableKind(n.Kind), ":", n.Path)
	} else if n.Kind == File {
		fmt.Println(spaces(level), ToHumanReadableKind(n.Kind), ":", n.Path)
	} else if n.Kind == Sequence || n.Kind == Mapping {
		fmt.Println(spaces(level), ToHumanReadableKind(n.Kind), ":", n.Path, n.Value)
	} else if n.Kind == Property {
		fmt.Println(spaces(level), ToHumanReadableKind(n.Kind), ":", n.Path, "->", n.Value)
	} else if n.Kind == Value {
		fmt.Println(spaces(level), ToHumanReadableKind(n.Kind), ":", n.Value)
	} else {
		return
	}

	if n != nil || n.Children != nil || len(n.Children) > 0 {
		for _, child := range n.Children {
			print(child, level+1)
		}
	}
}

func spaces(i int) string {
	s := ""
	for i > 0 {
		s += " "
		i--
	}
	return s
}
