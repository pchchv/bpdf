package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/cellwriter"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/properties"
	"github.com/stretchr/testify/assert"
)

func TestNewFillColorStyler(t *testing.T) {
	sut := cellwriter.NewFillColorStyler(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.fillColorStyler", fmt.Sprintf("%T", sut))
}

func TestFillColorStyle_Apply(t *testing.T) {
	t.Run("When prop is nil and next is nil, should skip calls", func(t *testing.T) {
		sut := cellwriter.NewFillColorStyler(nil)

		sut.Apply(100, 100, &entity.Config{}, nil)
	})

	t.Run("When prop is nil and next is filled, should skip current and call next", func(t *testing.T) {
		var nilCellProp *properties.Cell
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, nilCellProp)
		sut := cellwriter.NewFillColorStyler(nil)
		sut.SetNext(inner)

		sut.Apply(width, height, cfg, nilCellProp)

		inner.AssertNumberOfCalls(t, "Apply", 1)
	})

	t.Run("When has prop but background color is nil, should skip current and call next", func(t *testing.T) {
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &properties.Cell{}
		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, prop)
		sut := cellwriter.NewFillColorStyler(nil)
		sut.SetNext(inner)

		sut.Apply(width, height, cfg, prop)

		inner.AssertNumberOfCalls(t, "Apply", 1)
	})

	t.Run("When has prop and color is filled, should apply current and call next", func(t *testing.T) {
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &properties.Cell{
			BackgroundColor: &properties.Color{Red: 100, Green: 150, Blue: 170},
		}
		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, prop)
		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().SetFillColor(prop.BackgroundColor.Red, prop.BackgroundColor.Green, prop.BackgroundColor.Blue)
		fpdf.EXPECT().SetFillColor(255, 255, 255)
		sut := cellwriter.NewFillColorStyler(fpdf)
		sut.SetNext(inner)

		sut.Apply(width, height, cfg, prop)

		inner.AssertNumberOfCalls(t, "Apply", 1)
		fpdf.AssertNumberOfCalls(t, "SetFillColor", 2)
	})
}
