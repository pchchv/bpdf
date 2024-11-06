package image_test

import (
	"errors"
	"testing"

	"github.com/pchchv/bpdf/components/image"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/test"
	"github.com/stretchr/testify/assert"
)

func TestNewFromBytes(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := image.NewFromBytes([]byte{1, 2, 3}, extension.Jpg)

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_bytes_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := image.NewFromBytes([]byte{1, 2, 3}, extension.Jpg, fixture.RectProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_bytes_custom_prop.json")
	})
}

func TestNewFromBytesCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := image.NewFromBytesCol(12, []byte{1, 2, 3}, extension.Jpg)

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_bytes_col_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := image.NewFromBytesCol(12, []byte{1, 2, 3}, extension.Jpg, fixture.RectProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_bytes_col_custom_prop.json")
	})
}

func TestNewFromBytesRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := image.NewFromBytesRow(10, []byte{1, 2, 3}, extension.Jpg)

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_bytes_row_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := image.NewFromBytesRow(10, []byte{1, 2, 3}, extension.Jpg, fixture.RectProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_bytes_row_custom_prop.json")
	})
}

func TestNewAutoFromBytesRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := image.NewAutoFromBytesRow([]byte{1, 2, 3}, extension.Jpg)

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_bytes_auto_row_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := image.NewAutoFromBytesRow([]byte{1, 2, 3}, extension.Jpg, fixture.RectProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_bytes_auto_row_custom_prop.json")
	})
}

func TestBytesImage_GetHeight(t *testing.T) {
	t.Run("When it is not possible to know the dimensions of the bytes image, should return height 0", func(t *testing.T) {
		cell := fixture.CellEntity()
		img := fixture.ImageEntity()
		provider := mocks.NewProvider(t)
		provider.EXPECT().GetDimensionsByImageByte(img.Bytes, img.Extension).Return(nil, errors.New("anyError2"))
		sut := image.NewFromBytes(img.Bytes, img.Extension)

		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, height, 0.0)
	})

	t.Run("When the height of the bytes image is half the width, should return half the width of the cell", func(t *testing.T) {
		cell := fixture.CellEntity()
		img := fixture.ImageEntity()
		provider := mocks.NewProvider(t)
		provider.EXPECT().GetDimensionsByImageByte(img.Bytes, img.Extension).Return(&entity.Dimensions{Width: 10, Height: 5}, nil)
		sut := image.NewFromBytes(img.Bytes, img.Extension)

		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, height, cell.Width/2)
	})
}
