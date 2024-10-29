package gofpdf_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/jung-kurt/gofpdf"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/internal/math"
	gofpdf2 "github.com/pchchv/bpdf/internal/providers/gofpdf"
	"github.com/pchchv/bpdf/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewImage(t *testing.T) {
	image := gofpdf2.NewImage(mocks.NewFpdf(t), mocks.NewMath(t))

	assert.NotNil(t, image)
	assert.Equal(t, fmt.Sprintf("%T", image), "*gofpdf.image")
}

func TestImage_Add(t *testing.T) {
	t.Run("when RegisterImageOptionsReader return nil, should return error", func(t *testing.T) {
		cell := fixture.CellEntity()
		margins := fixture.MarginsEntity()
		rect := fixture.RectProp()
		img := fixture.ImageEntity()
		options := gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(img.Extension),
		}
		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().RegisterImageOptionsReader(mock.Anything, options, bytes.NewReader(img.Bytes)).Return(nil)
		image := gofpdf2.NewImage(pdf, mocks.NewMath(t))

		err := image.Add(&img, &cell, &margins, &rect, img.Extension, true)

		assert.NotNil(t, err)
	})

	t.Run("when prop is not center, should work properly", func(t *testing.T) {
		cell := fixture.CellEntity()
		margins := fixture.MarginsEntity()
		rect := fixture.RectProp()
		img := fixture.ImageEntity()
		options := gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(img.Extension),
		}
		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().RegisterImageOptionsReader(mock.Anything, options, bytes.NewReader(img.Bytes)).Return(&gofpdf.ImageInfoType{})
		pdf.EXPECT().Image(mock.Anything, 30.0, 35.0, 98.0, mock.Anything, true, "", 0, "")

		m := math.New()

		image := gofpdf2.NewImage(pdf, m)

		err := image.Add(&img, &cell, &margins, &rect, img.Extension, true)

		assert.Nil(t, err)
	})

	t.Run("when prop is center, should work properly", func(t *testing.T) {
		cell := fixture.CellEntity()
		margins := fixture.MarginsEntity()
		rect := fixture.RectProp()
		rect.Center = true
		img := fixture.ImageEntity()
		options := gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(img.Extension),
		}

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().RegisterImageOptionsReader(mock.Anything, options, bytes.NewReader(img.Bytes)).Return(&gofpdf.ImageInfoType{})
		pdf.EXPECT().Image(mock.Anything, 21.0, mock.Anything, 98.0, mock.Anything, true, "", 0, "")

		m := math.New()

		image := gofpdf2.NewImage(pdf, m)

		err := image.Add(&img, &cell, &margins, &rect, img.Extension, true)

		assert.Nil(t, err)
	})
}

func TestImage_GetImageInfo(t *testing.T) {
	t.Run("when RegisterImageOptionsReader return nil, should return nil", func(t *testing.T) {

		img := fixture.ImageEntity()
		options := gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(img.Extension),
		}

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().RegisterImageOptionsReader(mock.Anything, options, bytes.NewReader(img.Bytes)).Return(nil)

		image := gofpdf2.NewImage(pdf, mocks.NewMath(t))

		info, _ := image.GetImageInfo(&img, img.Extension)

		assert.Nil(t, info)
	})

	t.Run("when RegisterImageOptionsReader return info, should return info", func(t *testing.T) {

		img := fixture.ImageEntity()
		options := gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(img.Extension),
		}

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().RegisterImageOptionsReader(mock.Anything, options, bytes.NewReader(img.Bytes)).Return(&gofpdf.ImageInfoType{})

		image := gofpdf2.NewImage(pdf, mocks.NewMath(t))

		info, _ := image.GetImageInfo(&img, img.Extension)

		assert.NotNil(t, info)
	})
}
