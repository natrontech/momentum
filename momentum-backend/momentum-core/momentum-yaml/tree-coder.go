package yaml

import (
	"fmt"
	utils "momentum/momentum-core/momentum-utils"
	"os"
)

func EncodeToFile(tree *Node, path string) *os.File {

	if tree == nil || tree.Kind != DocumentNode || path == "" || utils.FileExists(path) {
		panic("Illegal arguments. Either tree is nil, not a DocumentNode or the path is empty or file does already exist.")
	}

	file, err := utils.FileOpen(path, os.O_CREATE|os.O_WRONLY)
	if err != nil {
		panic(err.Error())
	}

	writer := newEncoderWithWriter(file)
	writer.init()
	writer.node(tree, "")

	return nil
}

func DecodeToTree(path string) *Node {

	file, err := utils.FileOpen(path, os.O_RDONLY)
	if err != nil {
		fmt.Println("could not read file:", err.Error())
	}

	parser := newParserFromReader(file)
	parser.init()

	return parser.document()
}
