package image_test

import (
	"errors"
	"testing"

	"github.com/pchchv/bpdf/components/image"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/test"
	"github.com/stretchr/testify/assert"
)

func TestNewFromFile(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := image.NewFromFile("path")

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := image.NewFromFile("path", fixture.RectProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_custom_prop.json")
	})
}

func TestNewFromFileCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := image.NewFromFileCol(12, "path")

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_col_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := image.NewFromFileCol(12, "path", fixture.RectProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_col_custom_prop.json")
	})
}

func TestNewFromFileRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := image.NewFromFileRow(10, "path")

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_row_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := image.NewFromFileRow(12, "path", fixture.RectProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_row_custom_prop.json")
	})
}

func TestNewAutoFromFileRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := image.NewAutoFromFileRow("path")

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_auto_row_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := image.NewAutoFromFileRow("path", fixture.RectProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/images/new_image_from_file_auto_row_custom_prop.json")
	})
}

func TestFileImage_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		path := "path"
		cell := fixture.CellEntity()
		prop := fixture.RectProp()
		sut := image.NewFromFile(path, prop)
		provider := mocks.NewProvider(t)
		provider.EXPECT().AddImageFromFile(path, &cell, &prop)

		sut.Render(provider, &cell)

		provider.AssertNumberOfCalls(t, "AddImageFromFile", 1)
	})
}

func TestFileImageSetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		path := "path"
		prop := fixture.RectProp()
		sut := image.NewFromFile(path, prop)

		sut.SetConfig(nil)
	})
}

func TestFileImage_GetHeight(t *testing.T) {
	t.Run("When it is not possible to know the dimensions of the file image, should return height 0", func(t *testing.T) {
		cell := fixture.CellEntity()
		provider := mocks.NewProvider(t)
		provider.EXPECT().GetDimensionsByImage("path").Return(nil, errors.New("anyError2"))
		sut := image.NewFromFile("path")

		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, height, 0.0)
	})

	t.Run("When the height of the file image is half the width, should return half the width of the cell", func(t *testing.T) {
		cell := fixture.CellEntity()
		provider := mocks.NewProvider(t)
		provider.EXPECT().GetDimensionsByImage("path").Return(&entity.Dimensions{Width: 10, Height: 5}, nil)
		sut := image.NewFromFile("path")

		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, height, cell.Width/2)
	})
}
