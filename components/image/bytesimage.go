// Package image implements creation of images from file and bytes.
package image

import (
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

type BytesImage struct {
	bytes     []byte
	extension extension.Extension
	prop      properties.Rect
	config    *entity.Config
}
