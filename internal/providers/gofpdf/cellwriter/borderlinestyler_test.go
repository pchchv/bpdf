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

func TestNewBorderLineStyler(t *testing.T) {
	sut := cellwriter.NewBorderLineStyler(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.borderLineStyler", fmt.Sprintf("%T", sut))
}

func TestBorderLineStyler_Apply(t *testing.T) {
	t.Run("When prop is nil and next is nil, should skip calls", func(t *testing.T) {
		sut := cellwriter.NewBorderLineStyler(nil)

		sut.Apply(100, 100, &entity.Config{}, nil)
	})
	t.Run("When prop is nil and next is filled, should skip current and call next", func(t *testing.T) {
		var nilCellProp *properties.Cell
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, nilCellProp)
		sut := cellwriter.NewBorderLineStyler(nil)
		sut.SetNext(inner)

		sut.Apply(width, height, cfg, nilCellProp)

		inner.AssertNumberOfCalls(t, "Apply", 1)
	})

	t.Run("When has prop but line style is solid, should skip current and call next", func(t *testing.T) {
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &properties.Cell{
			LineStyle: linestyle.Solid,
		}
		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, prop)
		sut := cellwriter.NewBorderLineStyler(nil)
		sut.SetNext(inner)

		sut.Apply(width, height, cfg, prop)

		inner.AssertNumberOfCalls(t, "Apply", 1)
	})

	t.Run("When has prop but line style is empty, should skip current and call next", func(t *testing.T) {
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &properties.Cell{}
		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, prop)
		sut := cellwriter.NewBorderLineStyler(nil)
		sut.SetNext(inner)

		sut.Apply(width, height, cfg, prop)

		inner.AssertNumberOfCalls(t, "Apply", 1)
	})

	t.Run("When has prop and line style is dashed, should apply current and call next", func(t *testing.T) {
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &properties.Cell{
			LineStyle: linestyle.Dashed,
		}
		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, prop)
		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().SetDashPattern([]float64{1, 1}, 0.0)
		fpdf.EXPECT().SetDashPattern([]float64{1, 0}, 0.0)
		sut := cellwriter.NewBorderLineStyler(fpdf)
		sut.SetNext(inner)

		sut.Apply(width, height, cfg, prop)

		inner.AssertNumberOfCalls(t, "Apply", 1)
		fpdf.AssertNumberOfCalls(t, "SetDashPattern", 2)
	})
}
