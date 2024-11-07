package text_test

import (
	"testing"

	"github.com/pchchv/bpdf/components/text"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/properties"
	"github.com/pchchv/bpdf/test"
	"github.com/stretchr/testify/assert"
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

func TestText_GetHeight(t *testing.T) {
	t.Run("When top margin is sent, should increment row height with top margin", func(t *testing.T) {
		cell := fixture.CellEntity()
		font := fixture.FontProp()
		textProp := properties.Text{Top: 10}
		textProp.MakeValid(&font)
		sut := text.New("text", textProp)
		provider := mocks.NewProvider(t)
		provider.EXPECT().GetLinesQuantity("text", &textProp, 100.0).Return(5.0)
		provider.EXPECT().GetFontHeight(&font).Return(2.0)

		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, 20.0, height)
	})

	t.Run("When vertical padding is sent, should increment row height with vertical padding", func(t *testing.T) {
		cell := fixture.CellEntity()
		font := fixture.FontProp()
		textProp := properties.Text{VerticalPadding: 5}
		textProp.MakeValid(&font)
		sut := text.New("text", textProp)
		provider := mocks.NewProvider(t)
		provider.EXPECT().GetLinesQuantity("text", &textProp, 100.0).Return(5.0)
		provider.EXPECT().GetFontHeight(&font).Return(2.0)

		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, 30.0, height)
	})

	t.Run("When font has a height of 2, should return 10", func(t *testing.T) {
		cell := fixture.CellEntity()
		font := fixture.FontProp()
		textProp := properties.Text{}
		textProp.MakeValid(&font)
		sut := text.New("text", textProp)
		provider := mocks.NewProvider(t)
		provider.EXPECT().GetLinesQuantity("text", &textProp, 100.0).Return(5.0)
		provider.EXPECT().GetFontHeight(&font).Return(2.0)

		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, 10.0, height)
	})
}
