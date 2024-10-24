// The core package contains all core interfaces and basic implementations.
package core

import (
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/node"
)

// Node is the interface that wraps the basic methods of a node.
type Node interface {
	SetConfig(config *entity.Config)
	GetStructure() *node.Node[Structure]
}

// Component is the interface that wraps the basic methods of a component.
type Component interface {
	Node
	Render(provider Provider, cell *entity.Cell)
	GetHeight(provider Provider, cell *entity.Cell) float64
}
