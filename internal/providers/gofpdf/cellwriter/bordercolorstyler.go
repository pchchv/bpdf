package cellwriter

import (
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
