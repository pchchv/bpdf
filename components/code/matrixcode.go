// nolint:dupl
package code

import (
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/row"
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

// NewMatrix is responsible to create an instance of a MatrixCode.
func NewMatrix(code string, barcodeProps ...properties.Rect) core.Component {
	prop := properties.Rect{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &MatrixCode{
		code: code,
		prop: prop,
	}
}

// NewMatrixCol is responsible to create an instance of a MatrixCode wrapped in a Col.
func NewMatrixCol(size int, code string, ps ...properties.Rect) core.Col {
	matrixCode := NewMatrix(code, ps...)
	return col.New(size).Add(matrixCode)
}

// NewAutoMatrixRow is responsible to create an instance of a Matrix code wrapped in a Row with automatic height.
//   - code: The value that must be placed in the matrixcode
//   - ps: A set of settings that must be applied to the matrixcode
func NewAutoMatrixRow(code string, ps ...properties.Rect) core.Row {
	matrixCode := NewMatrix(code, ps...)
	c := col.New().Add(matrixCode)
	return row.New().Add(c)
}

// NewMatrixRow is responsible to create an instance of a MatrixCode wrapped in a Row.
func NewMatrixRow(height float64, code string, ps ...properties.Rect) core.Row {
	matrixCode := NewMatrix(code, ps...)
	c := col.New().Add(matrixCode)
	return row.New(height).Add(c)
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
