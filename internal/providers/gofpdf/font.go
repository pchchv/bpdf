package gofpdf

import (
	"github.com/pchchv/bpdf/consts/fontstyle"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

const (
	gofpdfFontScale1 = 72.0
	gofpdfFontScale2 = 25.4
)

type font struct {
	pdf         fpdfwrapper.Fpdf
	size        float64
	family      string
	style       fontstyle.Fontstyle
	scaleFactor float64
	fontColor   *properties.Color
}

// NewFont create a Font.
func NewFont(pdf fpdfwrapper.Fpdf, size float64, family string, style fontstyle.Fontstyle) *font {
	pdf.SetFont(family, string(style), size)
	return &font{
		pdf:         pdf,
		size:        size,
		family:      family,
		style:       style,
		scaleFactor: gofpdfFontScale1 / gofpdfFontScale2, // Bytes defined inside gofpdf constructor,
		fontColor:   &properties.Color{Red: 0, Green: 0, Blue: 0},
	}
}
