// Package col implements creation of columns.
package col

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

type Col struct {
	size       int
	isMax      bool
	components []core.Component
	config     *entity.Config
	style      *properties.Cell
}

// Add is responsible to add a component to a core.Col.
func (c *Col) Add(components ...core.Component) core.Col {
	c.components = append(c.components, components...)
	return c
}

// Render renders a core.Col into a PDF context.
func (c *Col) Render(provider core.Provider, cell entity.Cell, createCell bool) {
	if createCell {
		provider.CreateCol(cell.Width, cell.Height, c.config, c.style)
	}

	for _, component := range c.components {
		component.Render(provider, &cell)
	}
}
