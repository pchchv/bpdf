// Package code implements creation of Barcode, MatrixCode and QrCode.
// nolint:dupl
// It's similar to Barcode.go and it's hard to extract common code.
package code

import (
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/row"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/node"
	"github.com/pchchv/bpdf/properties"
)

type Barcode struct {
	code   string
	prop   properties.Barcode
	config *entity.Config
}

// NewBar is responsible to create an instance of a Barcode:
//   - code: value that must be placed in the barcode.
//   - ps: set of settings that must be applied to the barcode.
func NewBar(code string, ps ...properties.Barcode) core.Component {
	prop := properties.Barcode{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid()

	return &Barcode{
		code: code,
		prop: prop,
	}
}

// NewBarCol is responsible to create an instance of a Barcode wrapped in a Col.
//   - size: O tamanho da coluna
//   - code: The value that must be placed in the barcode
//   - ps: A set of settings that must be applied to the barcode
func NewBarCol(size int, code string, ps ...properties.Barcode) core.Col {
	bar := NewBar(code, ps...)
	return col.New(size).Add(bar)
}

// NewBarRow is responsible to create an instance of a Barcode wrapped in a Row.
// using this method the col size will be automatically set to the maximum value
//   - height: The height of the line
//   - code: The value that must be placed in the barcode
//   - ps: A set of settings that must be applied to the barcode
func NewBarRow(height float64, code string, ps ...properties.Barcode) core.Row {
	bar := NewBar(code, ps...)
	c := col.New().Add(bar)
	return row.New(height).Add(c)
}

// NewAutoBarRow is responsible to create an instance of a Barcode wrapped in a Row with automatic height.
// using this method the col size will be automatically set to the maximum value
//   - code: The value that must be placed in the barcode
//   - ps: A set of settings that must be applied to the barcode
func NewAutoBarRow(code string, ps ...properties.Barcode) core.Row {
	bar := NewBar(code, ps...)
	c := col.New().Add(bar)
	return row.New().Add(c)
}

// Render renders a Barcode into a PDF context.
// The bpdf cal this method in process to generate the pdf:
//   - provider: is the creator provider used to generate the pdf.
//   - cell: cell represents the space available to draw the component.
func (b *Barcode) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddBarCode(b.code, cell, &b.prop)
}

// SetConfig sets the configuration of a Barcode.
func (b *Barcode) SetConfig(config *entity.Config) {
	b.config = config
}

// GetStructure returns the structure of a barcode.
// This method is typically used when creating tests.
func (b *Barcode) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "barcode",
		Value:   b.code,
		Details: b.prop.ToMap(),
	}

	return node.New(str)
}

// GetHeight returns the height that the barcode will have in the PDF.
func (b *Barcode) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	proportion := b.prop.Proportion.Height / b.prop.Proportion.Width
	width := (b.prop.Percent / 100) * cell.Width
	return proportion * width
}
