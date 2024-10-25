package core

import (
	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/consts/fontstyle"
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

// Text is the abstraction which deals of how to add text inside PDF.
type Text interface {
	Add(text string, cell *entity.Cell, textProp *properties.Text)
	GetLinesQuantity(text string, textProp *properties.Text, colWidth float64) int
}

// Font is the abstraction which deals of how to set fontstyle configurations.
type Font interface {
	SetFamily(family string)
	SetStyle(style fontstyle.Fontstyle)
	SetSize(size float64)
	SetFont(family string, style fontstyle.Fontstyle, size float64)
	GetFamily() string
	GetStyle() fontstyle.Fontstyle
	GetSize() float64
	GetFont() (string, fontstyle.Fontstyle, float64)
	GetHeight(family string, style fontstyle.Fontstyle, size float64) float64
	SetColor(color *properties.Color)
	GetColor() *properties.Color
}

// Image is the abstraction which deals of how to add images in a PDF.
type Image interface {
	Add(img *entity.Image, cell *entity.Cell, margins *entity.Margins, prop *properties.Rect, extension extension.Extension, flow bool) error
	GetImageInfo(img *entity.Image, extension extension.Extension) (*gofpdf.ImageInfoType, uuid.UUID)
}
