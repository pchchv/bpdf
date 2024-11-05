package page_test

import (
	"testing"

	"github.com/pchchv/bpdf/components/image"
	"github.com/pchchv/bpdf/components/page"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/test"
)

func TestNew(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := page.New()

		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_page_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := page.New(fixture.PageProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_page_custom_prop.json")
	})

	t.Run("when prop is sent and there is rows, should use the provided", func(t *testing.T) {
		sut := page.New(fixture.PageProp())
		row := image.NewFromFileRow(10, "path")
		sut.Add(row)

		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_page_custom_prop_and_with_rows.json")
	})
}
