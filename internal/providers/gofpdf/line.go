package gofpdf

import (
	"github.com/pchchv/bpdf/consts/linestyle"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

type line struct {
	pdf              fpdfwrapper.Fpdf
	defaultColor     *properties.Color
	defaultThickness float64
}

func NewLine(pdf fpdfwrapper.Fpdf) *line {
	return &line{
		pdf:              pdf,
		defaultColor:     &properties.BlackColor,
		defaultThickness: linestyle.DefaultLineThickness,
	}
}
