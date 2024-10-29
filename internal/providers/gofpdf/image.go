package gofpdf

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
)

type image struct {
	pdf  fpdfwrapper.Fpdf
	math core.Math
}

// NewImage create an Image.
func NewImage(pdf fpdfwrapper.Fpdf, math core.Math) *image {
	return &image{
		pdf,
		math,
	}
}
