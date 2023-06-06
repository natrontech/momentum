package tree

import (
	"errors"
	"fmt"
	"os"
	"strings"

	utils "momentum/momentum-core/momentum-utils"
	yaml "momentum/momentum-core/momentum-yaml"
)

func Parse(path string, excludes []string) (*Node, error) {

	dir, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stat, err := dir.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if !stat.IsDir() {
		rootFile, err := parseFile(path)
		if err != nil {
			return nil, err
		}
		rootFile.Path = path // overwrite root with absolute path
		return rootFile, nil
	}

	rootDir, err := parseDir(path, excludes)
	if err != nil {
		return nil, err
	}
	rootDir.Path = path // overwrite root with absolute path
	return rootDir, nil
}

func parseDir(path string, excludes []string) (*Node, error) {

	dir, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	stat, err := dir.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if !stat.IsDir() {
		e := errors.New("path must be directory")
		fmt.Println(e.Error())
		return nil, e
	}

	entries, err := dir.ReadDir(-1)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	dirNode := NewNode(Directory, utils.LastPartOfPath(path), "", nil, nil, nil)

	for _, entry := range entries {

		if contains(excludes, entry.Name()) {
			fmt.Println("EXCLUDE", entry.Name())
			continue
		}

		entryPath := utils.BuildPath(dir.Name(), entry.Name())
		if entry.IsDir() {
			parsed, err := parseDir(entryPath, excludes)
			if err != nil {
				return nil, err
			}
			dirNode.AddChild(parsed)
		} else {

			if strings.HasSuffix(entryPath, ".yml") || strings.HasSuffix(entryPath, ".yaml") {
				fileNode, err := parseFile(entryPath)
				if err != nil {
					return nil, err
				}
				if fileNode != nil {
					dirNode.AddChild(fileNode)
				}
			}
		}
	}

	return dirNode, nil
}

// fileNode == nil indicates empty file.
func parseFile(path string) (*Node, error) {

	if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
		if utils.FileEmpty(path) {
			return nil, nil
		}
		parseTree := yaml.DecodeToTree(path)
		mappedTree, err := MomentumTreeFromYaml(parseTree, path)
		if err != nil {
			return nil, err
		}
		return mappedTree, nil
	}
	return nil, errors.New("unsupported file type")
}

func contains(col []string, s string) bool {

	for _, c := range col {
		if s == c {
			return true
		}
	}
	return false
}
