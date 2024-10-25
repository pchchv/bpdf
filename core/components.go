package core

import "github.com/pchchv/bpdf/core/entity"

// Math is the abstraction which deals with useful calc.
type Math interface {
	GetInnerCenterCell(inner *entity.Dimensions, outer *entity.Dimensions) *entity.Cell
	Resize(inner *entity.Dimensions, outer *entity.Dimensions, percent float64, justReferenceWidth bool) *entity.Dimensions
}
