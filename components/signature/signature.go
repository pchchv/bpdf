// Package signature implements creation of signatures.
package signature

import (
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/row"
	"github.com/pchchv/bpdf/consts/align"
	"github.com/pchchv/bpdf/consts/fontfamily"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/node"
	"github.com/pchchv/bpdf/properties"
)

type Signature struct {
	value  string
	prop   properties.Signature
	config *entity.Config
}

// New is responsible to create an instance of a Signature.
func New(value string, ps ...properties.Signature) core.Component {
	prop := properties.Signature{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid(fontfamily.Arial)

	return &Signature{
		value: value,
		prop:  prop,
	}
}

// NewCol is responsible to create an instance of a Signature wrapped in a Col.
func NewCol(size int, value string, ps ...properties.Signature) core.Col {
	signature := New(value, ps...)
	return col.New(size).Add(signature)
}

// NewRow is responsible to create an instance of a Signature wrapped in a Row.
func NewRow(height float64, value string, ps ...properties.Signature) core.Row {
	signature := New(value, ps...)
	c := col.New().Add(signature)
	return row.New(height).Add(c)
}

// NewAutoRow is responsible to create an instance of a Signature wrapped in a automatic Row.
func NewAutoRow(value string, ps ...properties.Signature) core.Row {
	signature := New(value, ps...)
	c := col.New().Add(signature)
	return row.New().Add(c)
}

// GetStructure returns the Structure of a Signature.
func (s *Signature) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "signature",
		Value:   s.value,
		Details: s.prop.ToMap(),
	}

	return node.New(str)
}

// GetHeight returns the height that the signature will have in the PDF.
func (s *Signature) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	return s.prop.LineThickness + provider.GetFontHeight(s.prop.ToFontProp())*s.prop.SafePadding
}

// Render renders a Signature into a PDF context.
func (s *Signature) Render(provider core.Provider, cell *entity.Cell) {
	fontSize := provider.GetFontHeight(s.prop.ToFontProp()) * s.prop.SafePadding
	textProp := s.prop.ToTextProp(align.Center, cell.Height-fontSize, 0)
	offsetPercent := (cell.Height - fontSize) / cell.Height * 100.0

	provider.AddText(s.value, cell, textProp)
	provider.AddLine(cell, s.prop.ToLineProp(offsetPercent))
}

// SetConfig sets the config.
func (s *Signature) SetConfig(config *entity.Config) {
	s.config = config
}
