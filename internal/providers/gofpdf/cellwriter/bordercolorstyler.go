package cellwriter

import (
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

type borderColorStyler struct {
	stylerTemplate
	defaultColor *properties.Color
}

func NewBorderColorStyler(fpdf fpdfwrapper.Fpdf) *borderColorStyler {
	return &borderColorStyler{
		stylerTemplate: stylerTemplate{
			fpdf: fpdf,
			name: "borderColorStyler",
		},
		defaultColor: &properties.BlackColor,
	}
}

func (b *borderColorStyler) Apply(width, height float64, config *entity.Config, prop *properties.Cell) {
	if prop == nil {
		b.GoToNext(width, height, config, prop)
		return
	}

	if prop.BorderColor == nil {
		b.GoToNext(width, height, config, prop)
		return
	}

	b.fpdf.SetDrawColor(prop.BorderColor.Red, prop.BorderColor.Green, prop.BorderColor.Blue)
	b.GoToNext(width, height, config, prop)
	b.fpdf.SetDrawColor(b.defaultColor.Red, b.defaultColor.Green, b.defaultColor.Blue)
}
