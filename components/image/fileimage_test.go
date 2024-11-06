package image_test

import (
	"testing"

	"github.com/pchchv/bpdf/components/image"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/test"
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
