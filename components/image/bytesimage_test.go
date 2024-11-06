package image_test

import (
	"testing"

	"github.com/pchchv/bpdf/components/image"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/test"
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
