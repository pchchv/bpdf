package core

import (
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

// Provider is the abstraction of a document creator provider.
type Provider interface {
	// Grid
	CreateRow(height float64)
	CreateCol(width, height float64, config *entity.Config, prop *properties.Cell)
	// Features
	AddLine(cell *entity.Cell, prop *properties.Line)
	AddText(text string, cell *entity.Cell, prop *properties.Text)
	GetFontHeight(prop *properties.Font) float64
	GetLinesQuantity(text string, textProp *properties.Text, colWidth float64) int
	AddMatrixCode(code string, cell *entity.Cell, prop *properties.Rect)
	AddQrCode(code string, cell *entity.Cell, rect *properties.Rect)
	AddBarCode(code string, cell *entity.Cell, prop *properties.Barcode)
	GetDimensionsByMatrixCode(code string) (*entity.Dimensions, error)
	GetDimensionsByImageByte(bytes []byte, extension extension.Extension) (*entity.Dimensions, error)
	GetDimensionsByImage(file string) (*entity.Dimensions, error)
	GetDimensionsByQrCode(code string) (*entity.Dimensions, error)
	AddImageFromFile(value string, cell *entity.Cell, prop *properties.Rect)
	AddImageFromBytes(bytes []byte, cell *entity.Cell, prop *properties.Rect, extension extension.Extension)
	AddBackgroundImageFromBytes(bytes []byte, cell *entity.Cell, prop *properties.Rect, extension extension.Extension)
	// General
	GenerateBytes() ([]byte, error)
	SetProtection(protection *entity.Protection)
	SetCompression(compression bool)
	SetMetadata(metadata *entity.Metadata)
}
