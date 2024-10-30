package cellwriter

import (
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

type fillColorStyler struct {
	stylerTemplate
	defaultFillColor *properties.Color
}

func NewFillColorStyler(fpdf fpdfwrapper.Fpdf) *fillColorStyler {
	return &fillColorStyler{
		stylerTemplate: stylerTemplate{
			fpdf: fpdf,
			name: "fillColorStyler",
		},
		defaultFillColor: &properties.WhiteColor,
	}
}

func (f *fillColorStyler) Apply(width, height float64, config *entity.Config, prop *properties.Cell) {
	if prop == nil {
		f.GoToNext(width, height, config, prop)
		return
	}

	if prop.BackgroundColor == nil {
		f.GoToNext(width, height, config, prop)
		return
	}

	f.fpdf.SetFillColor(prop.BackgroundColor.Red, prop.BackgroundColor.Green, prop.BackgroundColor.Blue)
	f.GoToNext(width, height, config, prop)
	f.fpdf.SetFillColor(f.defaultFillColor.Red, f.defaultFillColor.Green, f.defaultFillColor.Blue)
}
