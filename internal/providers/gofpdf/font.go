package gofpdf

import (
	"github.com/pchchv/bpdf/consts/fontstyle"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

type font struct {
	pdf         fpdfwrapper.Fpdf
	size        float64
	family      string
	style       fontstyle.Fontstyle
	scaleFactor float64
	fontColor   *properties.Color
}
