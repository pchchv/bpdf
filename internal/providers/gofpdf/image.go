package gofpdf

import (
	"bytes"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
)

type image struct {
	pdf  fpdfwrapper.Fpdf
	math core.Math
}

// NewImage create an Image.
func NewImage(pdf fpdfwrapper.Fpdf, math core.Math) *image {
	return &image{
		pdf,
		math,
	}
}

// GetImageInfo is responsible for loading the image in PDF and returning its information
func (s image) GetImageInfo(img *entity.Image, extension extension.Extension) (*gofpdf.ImageInfoType, uuid.UUID) {
	imageID, _ := uuid.NewRandom()
	info := s.pdf.RegisterImageOptionsReader(
		imageID.String(),
		gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(extension),
		},
		bytes.NewReader(img.Bytes),
	)
	return info, imageID
}
