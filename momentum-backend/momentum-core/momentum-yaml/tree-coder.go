package yaml

import (
	"errors"
	"fmt"
	utils "momentum/momentum-core/momentum-utils"
	"os"
)

func EncodeToFile(tree *Node, path string, allowOverwrite bool) (*os.File, error) {

	if tree == nil || tree.Kind != DocumentNode || path == "" {
		return nil, errors.New("illegal arguments either tree is nil not a DocumentNode or the path is empty")
	}

	if utils.FileExists(path) && !allowOverwrite {
		return nil, errors.New("file does already exist you can allow overwrite")
	}

	file, err := utils.FileOpen(path, os.O_CREATE|os.O_WRONLY)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	writer := newEncoderWithWriter(file)
	writer.init()
	writer.node(tree, "")

	return file, nil
}

func DecodeToTree(path string) *Node {

	file, err := utils.FileOpen(path, os.O_RDONLY)
	if err != nil {
		fmt.Println("could not read file:", err.Error())
	}
	defer file.Close()

	parser := newParserFromReader(file)
	parser.init()

	return parser.document()
}
