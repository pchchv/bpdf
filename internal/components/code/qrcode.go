// nolint:dupl
package code

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/node"
	"github.com/pchchv/bpdf/properties"
)

type QrCode struct {
	code   string
	prop   properties.Rect
	config *entity.Config
}

// Render renders a QrCode into a PDF context.
func (q *QrCode) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddQrCode(q.code, cell, &q.prop)
}

// SetConfig set the config for the component.
func (q *QrCode) SetConfig(config *entity.Config) {
	q.config = config
}

// GetStructure returns the Structure of a QrCode.
func (q *QrCode) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "qrcode",
		Value:   q.code,
		Details: q.prop.ToMap(),
	}

	return node.New(str)
}

// GetHeight returns the height that the QrCode will have in the PDF.
func (q *QrCode) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	dimensions, err := provider.GetDimensionsByQrCode(q.code)
	if err != nil {
		return 0
	}

	proportion := dimensions.Height / dimensions.Width
	width := (q.prop.Percent / 100) * cell.Width
	return proportion * width
}
