package entity_test

import (
	"testing"

	"github.com/pchchv/bpdf/core/entity"
	"github.com/stretchr/testify/assert"
)

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
