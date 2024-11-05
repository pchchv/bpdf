// Package image implements creation of images from file and bytes.
package image

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

type FileImage struct {
	path   string
	prop   properties.Rect
	config *entity.Config
}

// Render renders an Image into a PDF context.
func (f *FileImage) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddImageFromFile(f.path, cell, &f.prop)
}

// SetConfig sets the pdf config.
func (f *FileImage) SetConfig(config *entity.Config) {
	f.config = config
}
