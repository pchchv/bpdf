// Package col implements creation of columns.
package col

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

type Col struct {
	size       int
	isMax      bool
	components []core.Component
	config     *entity.Config
	style      *properties.Cell
}
