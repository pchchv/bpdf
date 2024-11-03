package test

import (
	"testing"

	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/node"
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
