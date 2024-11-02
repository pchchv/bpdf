// Package row implements creation of rows.
package row

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

type Row struct {
	height     float64
	autoHeight bool
	cols       []core.Col
	style      *properties.Cell
	config     *entity.Config
}

// Returns the height of the row content.
func (r *Row) getBiggestCol(provider core.Provider, cell *entity.Cell) float64 {
	greaterHeight := 0.0
	for _, col := range r.cols {
		height := col.GetHeight(provider, cell)
		if greaterHeight < height {
			greaterHeight = height
		}
	}
	return greaterHeight
}

// resetHeight resets the line height to 0.
func (r *Row) resetHeight() {
	r.height = 0
}
