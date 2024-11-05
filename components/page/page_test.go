package page_test

import (
	"testing"

	"github.com/pchchv/bpdf/components/image"
	"github.com/pchchv/bpdf/components/page"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/test"
	"github.com/stretchr/testify/assert"
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

func TestPage_GetRows(t *testing.T) {
	t.Run("when called get rows, should return rows correctly", func(t *testing.T) {
		row := mocks.NewRow(t)
		sut := page.New()
		sut.Add(row)

		rows := sut.GetRows()

		assert.Equal(t, []core.Row{row}, rows)
	})
}

func TestPage_SetNumber(t *testing.T) {
	t.Run("when called set number, should set correctly", func(t *testing.T) {
		sut := page.New()

		sut.SetNumber(1, 2)

		assert.Equal(t, 1, sut.GetNumber())
	})
}
