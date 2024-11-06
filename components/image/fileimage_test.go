package image_test

import (
	"testing"

	"github.com/pchchv/bpdf/components/image"
	"github.com/pchchv/bpdf/internal/fixture"
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
