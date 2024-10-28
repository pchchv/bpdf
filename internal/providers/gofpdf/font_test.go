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
