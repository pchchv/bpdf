package code

import (
	"bytes"
	"image"
	"image/color/palette"
	"image/draw"
	"image/png"

	bc "github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/boombuler/barcode/ean"
	"github.com/boombuler/barcode/qr"
	"github.com/pchchv/bpdf/consts/barcode"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
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

// GenDataMatrix is responsible to generate a data matrix byte array.
func (c *code) GenDataMatrix(code string) (*entity.Image, error) {
	dataMatrix, err := datamatrix.Encode(code)
	if err != nil {
		return nil, err
	}

	return c.getImage(dataMatrix)
}

// GenQr is responsible to generate a qr code byte array.
func (c *code) GenQr(code string) (*entity.Image, error) {
	qrCode, err := qr.Encode(code, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	return c.getImage(qrCode)
}

// GenBar is responsible to generate a barcode byte array.
func (c *code) GenBar(code string, _ *entity.Cell, prop *properties.Barcode) (*entity.Image, error) {
	barcodeGen := getBarcodeClosure(prop.Type)
	barCode, err := barcodeGen(code)
	if err != nil {
		return nil, err
	}

	width := float64(barCode.Bounds().Dx())
	heightPercentFromWidth := prop.Proportion.Height / prop.Proportion.Width
	height := int(width * heightPercentFromWidth)
	scaledBarCode, err := bc.Scale(barCode, int(width), height)
	if err != nil {
		return nil, err
	}

	return c.getImage(scaledBarCode)
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

func getBarcodeClosure(barcodeType barcode.Code) func(code string) (bc.BarcodeIntCS, error) {
	if barcodeType == barcode.EAN {
		return ean.Encode
	}
	return code128.Encode
}
