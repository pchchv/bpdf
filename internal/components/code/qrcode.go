// nolint:dupl
package code

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/components/col"
	"github.com/pchchv/bpdf/internal/components/row"
	"github.com/pchchv/bpdf/node"
	"github.com/pchchv/bpdf/properties"
)

type QrCode struct {
	code   string
	prop   properties.Rect
	config *entity.Config
}

// NewQr is responsible to create an instance of a QrCode.
func NewQr(code string, barcodeProps ...properties.Rect) core.Component {
	prop := properties.Rect{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &QrCode{
		code: code,
		prop: prop,
	}
}

// NewQrCol is responsible to create an instance of a QrCode wrapped in a Col.
func NewQrCol(size int, code string, ps ...properties.Rect) core.Col {
	qrCode := NewQr(code, ps...)
	return col.New(size).Add(qrCode)
}

// NewQrRow is responsible to create an instance of a QrCode wrapped in a Row.
func NewQrRow(height float64, code string, ps ...properties.Rect) core.Row {
	qrCode := NewQr(code, ps...)
	c := col.New().Add(qrCode)
	return row.New(height).Add(c)
}

// NewAutoMatrixRow is responsible to create an instance of a qrcode wrapped in a Row with automatic height.
//   - code: The value that must be placed in the qrcode
//   - ps: A set of settings that must be applied to the qrcode
func NewAutoQrRow(code string, ps ...properties.Rect) core.Row {
	qrCode := NewQr(code, ps...)
	c := col.New().Add(qrCode)
	return row.New().Add(c)
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
