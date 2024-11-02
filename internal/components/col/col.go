// Package col implements creation of columns.
package col

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/node"
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

// GetSize returns the size of a core.Col.
func (c *Col) GetSize() int {
	if c.isMax {
		return c.config.MaxGridSize
	}

	return c.size
}

// GetHeight returns the height of the column content
func (c *Col) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	innerCell := cell.Copy()
	percent := float64(c.GetSize()) / float64(c.config.MaxGridSize)
	innerCell.Width *= percent
	greaterHeight := 0.0
	for _, component := range c.components {
		height := component.GetHeight(provider, &innerCell)
		if greaterHeight < height {
			greaterHeight = height
		}
	}

	return greaterHeight
}

// GetStructure returns the Structure of a core.Col.
func (c *Col) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "col",
		Value:   c.size,
		Details: c.style.ToMap(),
	}

	if c.isMax {
		if len(str.Details) == 0 {
			str.Details = make(map[string]interface{})
		}
		str.Details["is_max"] = true
	}

	node := node.New(str)
	for _, c := range c.components {
		inner := c.GetStructure()
		node.AddNext(inner)
	}

	return node
}

// SetConfig set the config for the component.
func (c *Col) SetConfig(config *entity.Config) {
	c.config = config
	for _, component := range c.components {
		component.SetConfig(config)
	}
}

// WithStyle sets the style for the column.
func (c *Col) WithStyle(style *properties.Cell) core.Col {
	c.style = style
	return c
}
