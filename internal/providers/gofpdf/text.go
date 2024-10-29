package gofpdf

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
)

type text struct {
	pdf  fpdfwrapper.Fpdf
	math core.Math
	font core.Font
}

// NewText create a Text.
func NewText(pdf fpdfwrapper.Fpdf, math core.Math, font core.Font) *text {
	return &text{
		pdf,
		math,
		font,
	}
}
