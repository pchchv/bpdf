// Package image implements creation of images from file and bytes.
package image

import (
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

type FileImage struct {
	path   string
	prop   properties.Rect
	config *entity.Config
}
