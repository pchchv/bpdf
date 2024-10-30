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

func TestNewBorderColorStyler(t *testing.T) {
	sut := cellwriter.NewBorderColorStyler(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.borderColorStyler", fmt.Sprintf("%T", sut))
}

func TestBorderColorStyler_Apply(t *testing.T) {
	t.Run("When prop is nil and next is nil, should skip calls", func(t *testing.T) {
		sut := cellwriter.NewBorderColorStyler(nil)

		sut.Apply(100, 100, &entity.Config{}, nil)
	})

	t.Run("When prop is nil and next is filled, should skip current and call next", func(t *testing.T) {
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		var nilCellProp *properties.Cell
		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, nilCellProp)
		sut := cellwriter.NewBorderColorStyler(nil)
		sut.SetNext(inner)

		sut.Apply(width, height, cfg, nilCellProp)

		inner.AssertNumberOfCalls(t, "Apply", 1)
	})

	t.Run("When has prop but border color is nil, should skip current and call next", func(t *testing.T) {
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &properties.Cell{}
		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, prop)
		sut := cellwriter.NewBorderColorStyler(nil)
		sut.SetNext(inner)

		sut.Apply(width, height, cfg, prop)

		inner.AssertNumberOfCalls(t, "Apply", 1)
	})

	t.Run("When has prop and border color is defined, should apply current and call next", func(t *testing.T) {
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &properties.Cell{
			BorderColor: &properties.Color{Red: 140, Green: 100, Blue: 80},
		}
		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, prop)
		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().SetDrawColor(prop.BorderColor.Red, prop.BorderColor.Green, prop.BorderColor.Blue)
		fpdf.EXPECT().SetDrawColor(0, 0, 0)
		sut := cellwriter.NewBorderColorStyler(fpdf)
		sut.SetNext(inner)

		sut.Apply(width, height, cfg, prop)

		inner.AssertNumberOfCalls(t, "Apply", 1)
		fpdf.AssertNumberOfCalls(t, "SetDrawColor", 2)
	})
}
