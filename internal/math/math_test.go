package math_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/math"
	"github.com/stretchr/testify/assert"
)

func TestNewMath(t *testing.T) {
	sut := math.New()

	assert.NotNil(t, sut)
	assert.Equal(t, "*math.math", fmt.Sprintf("%T", sut))
}

func TestMath_GetCenterCorrection(t *testing.T) {
	t.Run("should get center correction correctly", func(t *testing.T) {
		sut := math.New()
		outerSize := 100.0
		innerSize := 50.0

		correction := sut.GetCenterCorrection(outerSize, innerSize)

		assert.Equal(t, 25.0, correction)
	})
}

func TestMath_GetInnerCenterCell(t *testing.T) {
	t.Run("there is not side-effect", func(t *testing.T) {
		sut := math.New()

		inner := &entity.Dimensions{Width: 100, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		_ = sut.GetInnerCenterCell(inner, outer)

		assert.Equal(t, 100.0, inner.Width)
		assert.Equal(t, 100.0, inner.Height)
		assert.Equal(t, 100.0, outer.Width)
		assert.Equal(t, 100.0, outer.Height)
	})

	t.Run("when inner and outer have the same size, should return the center", func(t *testing.T) {
		sut := math.New()

		inner := &entity.Dimensions{Width: 100, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		cell := sut.GetInnerCenterCell(inner, outer)

		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
	})

	t.Run("when inner is smaller than outer and has equal proportion, the center of the cell must be returned", func(t *testing.T) {
		sut := math.New()

		inner := &entity.Dimensions{Width: 80, Height: 80}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		cell := sut.GetInnerCenterCell(inner, outer)

		assert.Equal(t, 10.0, cell.X)
		assert.Equal(t, 10.0, cell.Y)
	})

	t.Run("when the internal one has a smaller height and smaller proportion than the external one, the center of the cell must be returned",
		func(t *testing.T) {
			sut := math.New()

			outer := &entity.Dimensions{Width: 100, Height: 100}
			inner := &entity.Dimensions{Width: 75.0, Height: 60.0}

			cell := sut.GetInnerCenterCell(inner, outer)

			assert.Equal(t, 12.5, cell.X)
			assert.Equal(t, 20.0, cell.Y)
		})

	t.Run("when internal has a smaller width and greater proportion than external, the center of the cell must be returned",
		func(t *testing.T) {
			sut := math.New()

			inner := &entity.Dimensions{Width: 80.0, Height: 100}
			outer := &entity.Dimensions{Width: 100, Height: 100}

			cell := sut.GetInnerCenterCell(inner, outer)

			assert.Equal(t, 10.0, cell.X)
			assert.Equal(t, 0.0, cell.Y)
		})

	t.Run("when internal has greater height and proportion than external, the center of the cell must be returned", func(t *testing.T) {
		sut := math.New()

		inner := &entity.Dimensions{Width: 60.0, Height: 75.0}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		cell := sut.GetInnerCenterCell(inner, outer)

		assert.Equal(t, 20.0, cell.X)
		assert.Equal(t, 12.5, cell.Y)
	})

	t.Run("quando interno tiver largura maior e proporção menor que externa, the center of the cell must be returned", func(t *testing.T) {
		sut := math.New()

		inner := &entity.Dimensions{Width: 100, Height: 80}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		cell := sut.GetInnerCenterCell(inner, outer)

		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 10.0, cell.Y)
	})
}
