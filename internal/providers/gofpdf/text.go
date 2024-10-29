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
