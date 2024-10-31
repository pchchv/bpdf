package code

import (
	"bytes"
	"image"
	"image/color/palette"
	"image/draw"
	"image/png"

	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core/entity"
)

// codeInstance is a singleton of code.
// It is used to ensure that it is not instantiated more than once,
// as it is not necessary since the code is stateless.
var codeInstance *code = nil

type code struct{}

// New create a Code (Singleton).
func New() *code {
	if codeInstance == nil {
		codeInstance = &code{}
	}
	return codeInstance
}

func (c *code) getImage(img image.Image) (*entity.Image, error) {
	var buf bytes.Buffer
	dst := image.NewPaletted(img.Bounds(), palette.Plan9)
	drawer := draw.Drawer(draw.Src)
	drawer.Draw(dst, dst.Bounds(), img, img.Bounds().Min)
	if err := png.Encode(&buf, dst); err != nil {
		return nil, err
	}

	imgEntity := &entity.Image{
		Bytes:     buf.Bytes(),
		Extension: extension.Png,
		Dimensions: &entity.Dimensions{
			Width:  float64(dst.Bounds().Dx()),
			Height: float64(dst.Bounds().Dy()),
		},
	}

	return imgEntity, nil
}
