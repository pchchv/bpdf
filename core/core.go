// The core package contains all core interfaces and basic implementations.
package core

import (
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/node"
	"github.com/pchchv/bpdf/properties"
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

// Col is the interface that wraps the basic methods of a col.
type Col interface {
	Node
	Add(components ...Component) Col
	GetSize() int
	GetHeight(provider Provider, cell *entity.Cell) float64
	WithStyle(style *properties.Cell) Col
	Render(provider Provider, cell entity.Cell, createCell bool)
}

// Row is the interface that wraps the basic methods of a row.
type Row interface {
	Node
	Add(cols ...Col) Row
	GetHeight(provider Provider, cell *entity.Cell) float64
	GetColumns() []Col
	WithStyle(style *properties.Cell) Row
	Render(provider Provider, cell entity.Cell)
}
