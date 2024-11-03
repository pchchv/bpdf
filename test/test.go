package test

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/node"
)

var (
	goModFile = "go.mod"
	BPDFFile  = ".bpdf.yml"
)

type Node struct {
	Value   interface{}            `json:"value,omitempty"`
	Type    string                 `json:"type"`
	Details map[string]interface{} `json:"details,omitempty"`
	Nodes   []*Node                `json:"nodes,omitempty"`
}

// BPDFTest is the unit test instance.
type BPDFTest struct {
	t    *testing.T
	node *node.Node[core.Structure]
}

func getParentDir(path string) (newPath string) {
	dirs := strings.Split(path, "/")
	dirs = dirs[:len(dirs)-2]
	for _, dir := range dirs {
		newPath += dir + "/"
	}

	return
}

func hasFileInPath(file string, path string) (bool, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return false, err
	}

	for _, entry := range entries {
		if entry.Name() == file {
			return true, nil
		}
	}

	return false, nil
}

func getBPDFConfigFilePathRecursive(path string) (string, error) {
	hasBPDF, err := hasFileInPath(BPDFFile, path)
	if err != nil {
		return "", err
	}

	if hasBPDF {
		return path, nil
	}

	hasGoMod, err := hasFileInPath(goModFile, path)
	if err != nil {
		return "", err
	}

	if hasGoMod {
		return "", errors.New("found go.mod but not .bpdf.yml")
	}

	parentPath := getParentDir(path)
	return getBPDFConfigFilePathRecursive(parentPath)
}

func getBPDFConfigFilePath() (string, error) {
	path, _ := os.Getwd()
	path += "/"

	return getBPDFConfigFilePathRecursive(path)
}
