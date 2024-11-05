// Package image implements creation of images from file and bytes.
package image

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/node"
	"github.com/pchchv/bpdf/properties"
)

type FileImage struct {
	path   string
	prop   properties.Rect
	config *entity.Config
}

// NewFromFile is responsible to create an instance of an Image.
func NewFromFile(path string, ps ...properties.Rect) core.Component {
	prop := properties.Rect{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid()

	return &FileImage{
		path: path,
		prop: prop,
	}
}

// Render renders an Image into a PDF context.
func (f *FileImage) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddImageFromFile(f.path, cell, &f.prop)
}

// SetConfig sets the pdf config.
func (f *FileImage) SetConfig(config *entity.Config) {
	f.config = config
}

// GetStructure returns the Structure of an Image.
func (f *FileImage) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "fileImage",
		Value:   f.path,
		Details: f.prop.ToMap(),
	}

	return node.New(str)
}

// GetHeight returns the height that the image will have in the PDF
func (f *FileImage) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	dimensions, err := provider.GetDimensionsByImage(f.path)
	if err != nil {
		return 0.0
	}

	proportion := dimensions.Height / dimensions.Width
	width := (f.prop.Percent / 100) * cell.Width
	return (proportion * width) + f.prop.Top
}
