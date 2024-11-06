// Package image implements creation of images from file and bytes.
package image

import (
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/row"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/node"
	"github.com/pchchv/bpdf/properties"
)

type BytesImage struct {
	bytes     []byte
	extension extension.Extension
	prop      properties.Rect
	config    *entity.Config
}

// NewFromBytes is responsible to create an instance of an Image.
func NewFromBytes(bytes []byte, extension extension.Extension, ps ...properties.Rect) core.Component {
	prop := properties.Rect{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid()

	return &BytesImage{
		bytes:     bytes,
		prop:      prop,
		extension: extension,
	}
}

// NewFromBytesCol is responsible to create an instance of an Image wrapped in a Col.
func NewFromBytesCol(size int, bytes []byte, extension extension.Extension, ps ...properties.Rect) core.Col {
	image := NewFromBytes(bytes, extension, ps...)
	return col.New(size).Add(image)
}

// NewFromBytesRow is responsible to create an instance of an Image wrapped in a Row.
func NewFromBytesRow(height float64, bytes []byte, extension extension.Extension, ps ...properties.Rect) core.Row {
	image := NewFromBytes(bytes, extension, ps...)
	c := col.New().Add(image)
	return row.New(height).Add(c)
}

// GetStructure returns the Structure of an Image.
func (b *BytesImage) GetStructure() *node.Node[core.Structure] {
	trimLength := 10
	if len(b.bytes) < trimLength {
		trimLength = len(b.bytes)
	}

	str := core.Structure{
		Type:    "bytesImage",
		Value:   b.bytes[:trimLength],
		Details: b.prop.ToMap(),
	}

	str.Details["extension"] = b.extension
	str.Details["bytes_size"] = len(b.bytes)

	return node.New(str)
}

// GetHeight returns the height that the image will have in the PDF
func (b *BytesImage) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	dimensions, err := provider.GetDimensionsByImageByte(b.bytes, b.extension)
	if err != nil {
		return 0
	}

	proportion := dimensions.Height / dimensions.Width
	width := (b.prop.Percent / 100) * cell.Width
	return proportion * width
}

// SetConfig sets the pdf config.
func (b *BytesImage) SetConfig(config *entity.Config) {
	b.config = config
}

// Render renders an Image into a PDF context.
func (b *BytesImage) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddImageFromBytes(b.bytes, cell, &b.prop, b.extension)
}
