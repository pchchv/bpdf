package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/consts/linestyle"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/cellwriter"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/properties"
	"github.com/stretchr/testify/assert"
)

func TestNewBorderThicknessStyler(t *testing.T) {
	sut := cellwriter.NewBorderThicknessStyler(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.borderThicknessStyler", fmt.Sprintf("%T", sut))
}

func TestBorderThicknessStyler_Apply(t *testing.T) {
	t.Run("When prop is nil and next is nil, should skip calls", func(t *testing.T) {
		sut := cellwriter.NewBorderThicknessStyler(nil)

		sut.Apply(100, 100, &entity.Config{}, nil)
	})

	t.Run("When prop is nil and next is filled, should skip current and call next", func(t *testing.T) {
		var nilCellProp *properties.Cell
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, nilCellProp)
		sut := cellwriter.NewBorderThicknessStyler(nil)
		sut.SetNext(inner)

		sut.Apply(width, height, cfg, nilCellProp)

		inner.AssertNumberOfCalls(t, "Apply", 1)
	})

	t.Run("When has prop but thickness is 0.0, should skip current and call next", func(t *testing.T) {
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &properties.Cell{}
		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, prop)
		sut := cellwriter.NewBorderThicknessStyler(nil)
		sut.SetNext(inner)

		sut.Apply(width, height, cfg, prop)

		inner.AssertNumberOfCalls(t, "Apply", 1)
	})

	t.Run("When has prop and line style is dashed, should apply current and call next", func(t *testing.T) {
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &properties.Cell{
			BorderThickness: 1.0,
		}
		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, prop)
		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().SetLineWidth(prop.BorderThickness)
		fpdf.EXPECT().SetLineWidth(linestyle.DefaultLineThickness)
		sut := cellwriter.NewBorderThicknessStyler(fpdf)
		sut.SetNext(inner)

		sut.Apply(width, height, cfg, prop)

		inner.AssertNumberOfCalls(t, "Apply", 1)
		fpdf.AssertNumberOfCalls(t, "SetLineWidth", 2)
	})
}
