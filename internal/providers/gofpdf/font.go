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

// GetFamily return the currently Font family configured.
func (s *font) GetFamily() string {
	return s.family
}

// GetStyle return the currently Font style configured.
func (s *font) GetStyle() fontstyle.Fontstyle {
	return s.style
}

// GetSize return the currently Font size configured.
func (s *font) GetSize() float64 {
	return s.size
}
