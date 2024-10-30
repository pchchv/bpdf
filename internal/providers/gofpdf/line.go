package gofpdf

import (
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

type line struct {
	pdf              fpdfwrapper.Fpdf
	defaultColor     *properties.Color
	defaultThickness float64
}
