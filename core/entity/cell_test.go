package entity_test

import (
	"testing"

	"github.com/pchchv/bpdf/core/entity"
	"github.com/stretchr/testify/assert"
)

func TestCell_GetDimensions(t *testing.T) {
	cell := entity.Cell{
		X:      10,
		Y:      10,
		Width:  100,
		Height: 100,
	}
	dimensions := cell.GetDimensions()

	assert.Equal(t, 100.0, dimensions.Width)
	assert.Equal(t, 100.0, dimensions.Height)
}

func TestCell_Copy(t *testing.T) {
	t.Run("copy should return same values", func(t *testing.T) {
		cell := entity.Cell{
			X:      10,
			Y:      10,
			Width:  100,
			Height: 100,
		}
		copyCell := cell.Copy()

		assert.Equal(t, cell.X, copyCell.X)
		assert.Equal(t, cell.Y, copyCell.Y)
		assert.Equal(t, cell.Width, copyCell.Width)
		assert.Equal(t, cell.Height, copyCell.Height)
	})

	t.Run("copy should not allow side-effects", func(t *testing.T) {
		cell := entity.Cell{
			X:      10,
			Y:      10,
			Width:  100,
			Height: 100,
		}
		copyCell := cell.Copy()
		copyCell.X = 15
		copyCell.Y = 15
		copyCell.Width = 90
		copyCell.Height = 90

		assert.Equal(t, 10.0, cell.X)
		assert.Equal(t, 10.0, cell.Y)
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})
}

func TestNewRootContext(t *testing.T) {
	width := 100.0
	height := 300.0
	margins := entity.Margins{
		Left:   10,
		Right:  10,
		Top:    10,
		Bottom: 20,
	}
	cell := entity.NewRootCell(width, height, margins)

	assert.Equal(t, 0.0, cell.X)
	assert.Equal(t, 0.0, cell.Y)
	assert.Equal(t, 80.0, cell.Width)
	assert.Equal(t, 270.0, cell.Height)
}
