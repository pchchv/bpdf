// Page package implements creation of pages.
package page

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

type Page struct {
	number int
	total  int
	rows   []core.Row
	config *entity.Config
	prop   properties.PageNumber
}

// Render renders a Page into a PDF context.
func (p *Page) Render(provider core.Provider, cell entity.Cell) {
	prop := &properties.Rect{}
	innerCell := cell.Copy()
	prop.MakeValid()
	if p.config.BackgroundImage != nil {
		provider.AddBackgroundImageFromBytes(p.config.BackgroundImage.Bytes, &innerCell, prop, p.config.BackgroundImage.Extension)
	}

	for _, row := range p.rows {
		row.Render(provider, innerCell)
		innerCell.Y += row.GetHeight(provider, &innerCell)
	}

	if p.prop.Pattern != "" {
		provider.AddText(p.prop.GetPageString(p.number, p.total), &cell, p.prop.GetNumberTextProp(cell.Height))
	}
}

// SetConfig sets the Page configuration.
func (p *Page) SetConfig(config *entity.Config) {
	p.config = config
	for _, row := range p.rows {
		row.SetConfig(config)
	}
}

// SetNumber sets the Page number and total.
func (p *Page) SetNumber(number int, total int) {
	p.number = number
	p.total = total
}

// GetRows returns the rows of the Page.
func (p *Page) GetRows() []core.Row {
	return p.rows
}

// GetNumber returns the Page number.
func (p *Page) GetNumber() int {
	return p.number
}
