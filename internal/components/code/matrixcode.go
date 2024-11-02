// nolint:dupl
package code

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/node"
	"github.com/pchchv/bpdf/properties"
)

type MatrixCode struct {
	code   string
	prop   properties.Rect
	config *entity.Config
}

// SetConfig sets the configuration of a MatrixCode.
func (m *MatrixCode) SetConfig(config *entity.Config) {
	m.config = config
}

// Render renders a MatrixCode into a PDF context.
func (m *MatrixCode) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddMatrixCode(m.code, cell, &m.prop)
}

// GetStructure returns the Structure of a MatrixCode.
func (m *MatrixCode) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "matrixcode",
		Value:   m.code,
		Details: m.prop.ToMap(),
	}

	return node.New(str)
}

// GetHeight returns the height that the code will have in the PDF.
func (m *MatrixCode) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	dimensions, err := provider.GetDimensionsByMatrixCode(m.code)
	if err != nil {
		return 0
	}

	proportion := dimensions.Height / dimensions.Width
	width := (m.prop.Percent / 100) * cell.Width
	return proportion * width
}
