package cellwriter

import (
	"github.com/pchchv/bpdf/consts/linestyle"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

type borderLineStyler struct {
	stylerTemplate
}

func NewBorderLineStyler(fpdf fpdfwrapper.Fpdf) *borderLineStyler {
	return &borderLineStyler{
		stylerTemplate: stylerTemplate{
			fpdf: fpdf,
			name: "borderLineStyler",
		},
	}
}

func (b *borderLineStyler) Apply(width, height float64, config *entity.Config, prop *properties.Cell) {
	if prop == nil {
		b.GoToNext(width, height, config, prop)
		return
	}

	if prop.LineStyle == linestyle.Solid || prop.LineStyle == "" {
		b.GoToNext(width, height, config, prop)
		return
	}

	b.fpdf.SetDashPattern([]float64{1, 1}, 0)
	b.GoToNext(width, height, config, prop)
	b.fpdf.SetDashPattern([]float64{1, 0}, 0)
}
