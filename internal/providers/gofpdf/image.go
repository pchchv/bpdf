package gofpdf

import (
	"bytes"
	"errors"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
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

// Add use a byte array to add image to PDF.
func (s *image) Add(img *entity.Image, cell *entity.Cell, margins *entity.Margins, prop *properties.Rect, extension extension.Extension, flow bool) error {
	imageID, _ := uuid.NewRandom()
	info := s.pdf.RegisterImageOptionsReader(
		imageID.String(),
		gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(extension),
		},
		bytes.NewReader(img.Bytes),
	)
	if info == nil {
		return errors.New("could not register image options, maybe path/name is wrong")
	}

	s.addImageToPdf(imageID.String(), info, cell, margins, prop, flow)
	return nil
}

func (s *image) addImageToPdf(imageLabel string, info *gofpdf.ImageInfoType, cell *entity.Cell, margins *entity.Margins, prop *properties.Rect, flow bool) {
	dimensions := s.math.Resize(&entity.Dimensions{
		Width:  info.Width(),
		Height: info.Height(),
	}, cell.GetDimensions(), prop.Percent, prop.JustReferenceWidth)
	rectCell := &entity.Cell{X: prop.Left, Y: prop.Top, Width: dimensions.Width, Height: dimensions.Height}
	if prop.Center {
		rectCell = s.math.GetInnerCenterCell(dimensions, cell.GetDimensions())
	}

	s.pdf.Image(imageLabel, cell.X+rectCell.X+margins.Left, cell.Y+rectCell.Y+margins.Top,
		rectCell.Width, rectCell.Height, flow, "", 0, "")
}
