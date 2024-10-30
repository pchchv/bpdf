package cellwriter

import (
	"github.com/pchchv/bpdf/consts/linestyle"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

type borderThicknessStyler struct {
	stylerTemplate
	defaultLineThickness float64
}

func NewBorderThicknessStyler(fpdf fpdfwrapper.Fpdf) *borderThicknessStyler {
	return &borderThicknessStyler{
		stylerTemplate: stylerTemplate{
			fpdf: fpdf,
			name: "borderThicknessStyler",
		},
		defaultLineThickness: linestyle.DefaultLineThickness,
	}
}

func (b *borderThicknessStyler) Apply(width, height float64, config *entity.Config, prop *properties.Cell) {
	if prop == nil {
		b.GoToNext(width, height, config, prop)
		return
	}

	if prop.BorderThickness == 0 {
		b.GoToNext(width, height, config, prop)
		return
	}

	b.fpdf.SetLineWidth(prop.BorderThickness)
	b.GoToNext(width, height, config, prop)
	b.fpdf.SetLineWidth(b.defaultLineThickness)
}
