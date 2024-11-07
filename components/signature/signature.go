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
