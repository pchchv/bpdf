package cellwriter

import (
	"github.com/pchchv/bpdf/consts/linestyle"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
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
