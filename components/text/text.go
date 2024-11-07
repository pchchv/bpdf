// Package text implements creation of texts.
package text

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/node"
	"github.com/pchchv/bpdf/properties"
)

type Text struct {
	value  string
	prop   properties.Text
	config *entity.Config
}

// GetHeight returns the height that the text will have in the PDF.
func (t *Text) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	amountLines := provider.GetLinesQuantity(t.value, &t.prop, cell.Width-t.prop.Left-t.prop.Right)
	fontHeight := provider.GetFontHeight(&properties.Font{Family: t.prop.Family, Style: t.prop.Style, Size: t.prop.Size, Color: t.prop.Color})
	textHeight := float64(amountLines)*fontHeight + float64(amountLines-1)*t.prop.VerticalPadding
	return textHeight + t.prop.Top + t.prop.Bottom
}

// GetStructure returns the Structure of a Text.
func (t *Text) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "text",
		Value:   t.value,
		Details: t.prop.ToMap(),
	}

	return node.New(str)
}

// SetConfig sets the config.
func (t *Text) SetConfig(config *entity.Config) {
	t.config = config
	t.prop.MakeValid(t.config.DefaultFont)
}

// Render renders a Text into a PDF context.
func (t *Text) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddText(t.value, cell, &t.prop)
}
