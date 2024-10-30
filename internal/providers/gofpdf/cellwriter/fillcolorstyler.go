package cellwriter

import (
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
