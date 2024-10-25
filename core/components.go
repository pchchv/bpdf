package core

import (
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

// Math is the abstraction which deals with useful calc.
type Math interface {
	GetInnerCenterCell(inner *entity.Dimensions, outer *entity.Dimensions) *entity.Cell
	Resize(inner *entity.Dimensions, outer *entity.Dimensions, percent float64, justReferenceWidth bool) *entity.Dimensions
}

// Code is the abstraction which deals of how to add QrCodes or Barcode in a PDF.
type Code interface {
	GenQr(code string) (*entity.Image, error)
	GenDataMatrix(code string) (*entity.Image, error)
	GenBar(code string, cell *entity.Cell, prop *properties.Barcode) (*entity.Image, error)
}

type Line interface {
	Add(cell *entity.Cell, prop *properties.Line)
}
