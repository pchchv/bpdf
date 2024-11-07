package text_test

import (
	"testing"

	"github.com/pchchv/bpdf/components/text"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/test"
)

func TestNew(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := text.New("code")

		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := text.New("code", fixture.TextProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_custom_prop.json")
	})
}

func TestNewCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := text.NewCol(12, "code")

		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_col_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := text.NewCol(12, "code", fixture.TextProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_col_custom_prop.json")
	})
}

func TestNewRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := text.NewRow(10, "code")

		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_row_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := text.NewRow(10, "code", fixture.TextProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_row_custom_prop.json")
	})
}

func TestNewAutoRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := text.NewAutoRow("code")

		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_auto_row_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := text.NewAutoRow("code", fixture.TextProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/texts/new_text_auto_row_custom_prop.json")
	})
}
