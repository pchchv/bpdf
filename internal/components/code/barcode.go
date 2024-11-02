// Package code implements creation of Barcode, MatrixCode and QrCode.
// nolint:dupl
// It's similar to Barcode.go and it's hard to extract common code.
package code

import (
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

// Render renders a Barcode into a PDF context.
// The bpdf cal this method in process to generate the pdf:
//   - provider: Is the creator provider used to generate the pdf.
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
