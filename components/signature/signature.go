// Package signature implements creation of signatures.
package signature

import (
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
