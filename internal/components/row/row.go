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
