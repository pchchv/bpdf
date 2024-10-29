package gofpdf_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/consts/breakline"
	"github.com/pchchv/bpdf/consts/fontfamily"
	"github.com/pchchv/bpdf/consts/fontstyle"
	"github.com/pchchv/bpdf/internal/providers/gofpdf"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/properties"
	"github.com/stretchr/testify/assert"
)

func TestNewText(t *testing.T) {
	text := gofpdf.NewText(mocks.NewFpdf(t), mocks.NewMath(t), mocks.NewFont(t))

	assert.NotNil(t, text)
	assert.Equal(t, fmt.Sprintf("%T", text), "*gofpdf.text")
}

func TestGetLinesheight(t *testing.T) {
	t.Run("when a text that occupies two lines is sent with EmptySpaceStrategy, should two is returned", func(t *testing.T) {
		textProp := &properties.Text{}
		textProp.MakeValid(&properties.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().GetStringWidth("text ").Return(5)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		height := text.GetLinesQuantity("text text text text", textProp, 11)

		assert.Equal(t, 2, height)
	})

	t.Run("When a text that occupies two lines is sent with EmptySpaceStrategy, should two is returned", func(t *testing.T) {
		textProp := &properties.Text{BreakLineStrategy: breakline.DashStrategy}
		textProp.MakeValid(&properties.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().GetStringWidth("t").Return(1)
		pdf.EXPECT().GetStringWidth(" ").Return(1)
		pdf.EXPECT().GetStringWidth(" - ").Return(1)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		height := text.GetLinesQuantity("tttt tttt tttt tttt", textProp, 11)

		assert.Equal(t, 2, height)
	})
}