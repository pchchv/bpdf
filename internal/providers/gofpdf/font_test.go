package gofpdf_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/consts/fontfamily"
	"github.com/pchchv/bpdf/consts/fontstyle"
	"github.com/pchchv/bpdf/internal/providers/gofpdf"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/properties"
	"github.com/stretchr/testify/assert"
)

func TestNewFont(t *testing.T) {
	size := 10.0
	family := fontfamily.Arial
	style := fontstyle.Bold
	fpdf := mocks.NewFpdf(t)
	fpdf.EXPECT().SetFont(family, string(style), size)

	font := gofpdf.NewFont(fpdf, size, family, style)

	assert.NotNil(t, font)
	assert.Equal(t, fmt.Sprintf("%T", font), "*gofpdf.font")
	assert.Equal(t, family, font.GetFamily())
	assert.Equal(t, style, font.GetStyle())
	assert.Equal(t, size, font.GetSize())
	assert.Equal(t, &properties.Color{Red: 0, Green: 0, Blue: 0}, font.GetColor())
}

func TestFont_GetHeight(t *testing.T) {
	size := 10.0
	family := fontfamily.Arial
	style := fontstyle.Bold
	fpdf := mocks.NewFpdf(t)
	fpdf.EXPECT().SetFont(family, string(style), size)
	font := gofpdf.NewFont(fpdf, size, family, style)

	height := font.GetHeight(family, style, size)

	assert.Equal(t, 3.527777777777778, height)
}

func TestFont_SetFamily(t *testing.T) {
	size := 10.0
	family := fontfamily.Arial
	style := fontstyle.Bold
	fpdf := mocks.NewFpdf(t)
	fpdf.EXPECT().SetFont(family, string(style), size)
	fpdf.EXPECT().SetFont(fontfamily.Helvetica, string(style), size)
	font := gofpdf.NewFont(fpdf, size, family, style)

	font.SetFamily(fontfamily.Helvetica)

	assert.Equal(t, fontfamily.Helvetica, font.GetFamily())
}

func TestFont_SetStyle(t *testing.T) {
	size := 10.0
	family := fontfamily.Arial
	style := fontstyle.Bold

	fpdf := mocks.NewFpdf(t)
	fpdf.EXPECT().SetFont(family, string(style), size)
	fpdf.EXPECT().SetFontStyle(string(fontstyle.BoldItalic))
	font := gofpdf.NewFont(fpdf, size, family, style)

	font.SetStyle(fontstyle.BoldItalic)

	assert.Equal(t, fontstyle.BoldItalic, font.GetStyle())
}

func TestFont_SetSize(t *testing.T) {
	size := 10.0
	family := fontfamily.Arial
	style := fontstyle.Bold
	fpdf := mocks.NewFpdf(t)
	fpdf.EXPECT().SetFont(family, string(style), size)
	fpdf.EXPECT().SetFontSize(14.0)
	font := gofpdf.NewFont(fpdf, size, family, style)

	font.SetSize(14.0)

	assert.Equal(t, 14.0, font.GetSize())
}
